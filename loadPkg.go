package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func loadPkg(packageFile string, pkgName string) Package {
	var pkg Package

	environmentVariables := map[string]string{
		"PREFIX": config.Paths.Prefix,
		"BIN":    config.Paths.Prefix + "bin/",
		"HOME":   strings.TrimSuffix(home, "/"),
		"BOLD":   textFx["BOLD"],
		"RESET":  RESETCOL,
	}

	debugLog("Finding environment variables...")
	keySlice := make([]string, 0)
	for key := range environmentVariables {
		keySlice = append(keySlice, key)
	}

	debugLog("Replacing environment variables...")
	for _, key := range keySlice {
		environmentVariables["PREFIX"] = config.Paths.Prefix
		debugLog("Replacing %s with %s...", key, environmentVariables[key])
		packageFile = strings.ReplaceAll(packageFile, ":("+key+"):", environmentVariables[key])
	}

	err := json.Unmarshal([]byte(packageFile), &pkg)
	errorLog(err, "An error occurred while loading package info for %s", pkgName)

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
