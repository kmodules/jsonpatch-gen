/*
Copyright The Gomodules Authors.

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
	"fmt"
	"io/ioutil"
	"os"

	"gomodules.xyz/jsonpatch/v2"

	"github.com/pkg/errors"
	flag "github.com/spf13/pflag"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/yaml"
)

var (
	pt     string
	output string
)

func main() {
	flag.StringVar(&pt, "type", "json", "The type of patch being provided; one of [json strategic]")
	flag.StringVarP(&output, "output", "o", "yaml", "Output format json|yaml")
	flag.Parse()

	if len(os.Args) != 3 {
		fmt.Println("Usage: jsonpatch-gen from.json to.json")
		os.Exit(1)
	}

	var patch string
	var err error
	if pt == "json" {
		patch, err = generateJsonPatch(os.Args[1], os.Args[2])
	} else if pt == "strategic" {
		patch, err = generateStrategicMergePatch(os.Args[1], os.Args[2])
	} else {
		_, _ = fmt.Fprintln(os.Stderr, "unknown patch type", pt)
		os.Exit(1)
	}
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Println(patch)
	}
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

	patch, err := strategicpatch.CreateTwoWayMergePatch(fromJson, toJson, obj)
	if err != nil {
		return "", err
	}
	return string(patch), err
}
