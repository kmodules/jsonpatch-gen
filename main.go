package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"gomodules.xyz/jsonpatch/v2"
	"sigs.k8s.io/yaml"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: jsonpatch-gen from.json to.json")
		os.Exit(1)
	}
	if patch, err := main2(os.Args[1], os.Args[2]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Println(patch)
	}
}

func main2(fromFile, toFile string) (string, error) {
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

	patch, err := json.MarshalIndent(ops, "", "  ")
	return string(patch), err
}
