/*
Copyright The Kmodules Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"gomodules.xyz/jsonpatch/v2"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/yaml"
)

const fmtJSON = "json"

var (
	pt     string
	output string
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "jsonpatch-gen from.yaml to.yaml",
		Short: "Generate json patch",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("usage: jsonpatch-gen from.json to.json")
			}

			var patch string
			var err error
			if pt == fmtJSON {
				patch, err = generateJsonPatch(args[0], args[1])
			} else if pt == "strategic" {
				patch, err = generateStrategicMergePatch(args[0], args[1])
			} else {
				return fmt.Errorf("unknown patch type %s", pt)
			}
			if err != nil {
				return err
			}

			fmt.Println(patch)
			return nil
		},
	}
	rootCmd.Flags().StringVarP(&pt, "type", "t", fmtJSON, "The type of patch being provided; one of [json strategic]")
	rootCmd.Flags().StringVarP(&output, "output", "o", "yaml", "Output format; one of [json yaml]")
	rootCmd.Flags().AddGoFlagSet(flag.CommandLine)
	utilruntime.Must(flag.CommandLine.Parse([]string{}))

	utilruntime.Must(rootCmd.Execute())
}

func generateJsonPatch(fromFile, toFile string) (string, error) {
	fromData, err := ioutil.ReadFile(fromFile)
	if err != nil {
		return "", errors.Errorf("failed to read file %s. reason", fromFile, err)
	}
	fromJson, err := yaml.YAMLToJSON(fromData)
	if err != nil {
		return "", err
	}

	toData, err := ioutil.ReadFile(toFile)
	if err != nil {
		return "", errors.Errorf("failed to read file %s. reason", fromFile, err)
	}
	toJson, err := yaml.YAMLToJSON(toData)
	if err != nil {
		return "", err
	}

	ops, err := jsonpatch.CreatePatch(fromJson, toJson)
	if err != nil {
		return "", err
	}

	var patch []byte
	if output == "json" {
		patch, err = json.MarshalIndent(ops, "", "  ")
	} else {
		patch, err = yaml.Marshal(ops)
	}
	return string(patch), err
}

func generateStrategicMergePatch(fromFile, toFile string) (string, error) {
	fromData, err := ioutil.ReadFile(fromFile)
	if err != nil {
		return "", errors.Errorf("failed to read file %s. reason", fromFile, err)
	}
	fromJson, err := yaml.YAMLToJSON(fromData)
	if err != nil {
		return "", err
	}

	toData, err := ioutil.ReadFile(toFile)
	if err != nil {
		return "", errors.Errorf("failed to read file %s. reason", fromFile, err)
	}
	toJson, err := yaml.YAMLToJSON(toData)
	if err != nil {
		return "", err
	}

	var u unstructured.Unstructured
	_, gvk, err := unstructured.UnstructuredJSONScheme.Decode(fromJson, nil, &u)
	if err != nil {
		return "", err
	}

	obj, err := scheme.Scheme.New(*gvk)
	if err != nil {
		return "", err
	}

	jp, err := strategicpatch.CreateTwoWayMergePatch(fromJson, toJson, obj)
	if err != nil {
		return "", err
	}

	var overlay map[string]interface{}
	err = json.Unmarshal(jp, &overlay)
	if err != nil {
		return "", err
	}
	overlay["apiVersion"] = u.GetAPIVersion()
	overlay["kind"] = u.GetKind()
	err = unstructured.SetNestedField(overlay, u.GetName(), "metadata", "name")
	if err != nil {
		return "", err
	}

	var patch []byte
	if output == fmtJSON {
		patch, err = json.MarshalIndent(overlay, "", "  ")
	} else {
		patch, err = yaml.Marshal(overlay)
	}
	return string(patch), err
}
