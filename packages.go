package main

import (
	"encoding/json"
	"strings"
)

type Package struct {
	Name         string
	Description  string
	Url          string
	Install      []string
	Uninstall    []string
	Update       []string
	Config_paths []string
}

var environmentVariables = map[string]string{
	"PATH": "$HOME/.local/",
}

func load_package() {
	var pkg Package
	packageFile, _ := readFile("samples/pkg.json")

	keySlice := make([]string, 0)
	for key := range environmentVariables {
		keySlice = append(keySlice, key)
	}

	for _, key := range keySlice {
		packageFile = strings.Replace(packageFile, ":("+key+"):", environmentVariables[key], -1)
	}
	json.Unmarshal([]byte(packageFile), &pkg)
}
