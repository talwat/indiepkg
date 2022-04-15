package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func loadPkg(packageFile string, pkgName string) Package {
	var pkg Package

	debugLog("Finding environment variables...")
	keySlice := make([]string, 0)
	for key := range environmentVariables {
		keySlice = append(keySlice, key)
	}

	debugLog("Replacing environment variables...")
	for _, key := range keySlice {
		packageFile = strings.Replace(packageFile, ":("+key+"):", environmentVariables[key], -1)
	}

	err := json.Unmarshal([]byte(packageFile), &pkg)
	errorLog(err, 4, "An error occurred while loading package info for %s", pkgName)
	return pkg
}

func readLoad(pkgName string) Package {
	packageDisplayName := bolden(pkgName)

	log(1, "Reading package info for %s...", packageDisplayName)
	pkgFile := readFile(infoPath+pkgName+".json", "An error occurred while reading package %s", packageDisplayName)

	log(1, "Loading package info for %s...", packageDisplayName)
	pkg := loadPkg(pkgFile, fmt.Sprintf("An error occurred while loading package information for %s", packageDisplayName))

	return pkg
}
