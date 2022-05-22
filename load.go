package main

import (
	"encoding/json"
	"strings"
)

func loadPkg(packageFile string, pkgName string) Package {
	log(1, "Loading package info for %s...", bolden(pkgName))

	var pkg Package

	environmentVariables := map[string]string{
		"PREFIX":        config.Paths.Prefix,
		"BIN":           config.Paths.Prefix + "bin/",
		"HOME":          home,
		"INDIEPKG_PATH": mainPath,
		"BOLD":          textFx.Bold,
		"RESET":         RESETCOL,
	}

	debugLog("Finding environment variables...")

	keySlice := make([]string, 0)

	for key := range environmentVariables { // Iterate through environment variables & add them to keySlice
		keySlice = append(keySlice, key)
	}

	debugLog("Replacing environment variables...")

	for _, key := range keySlice { // Iterate through keySlice & replace them in the package file with their proper values
		environmentVariables["PREFIX"] = config.Paths.Prefix
		debugLog("Replacing %s with %s...", key, environmentVariables[key])
		packageFile = strings.ReplaceAll(packageFile, ":("+key+"):", environmentVariables[key])
	}

	err := json.Unmarshal([]byte(packageFile), &pkg)
	errorLog(err, "An error occurred while loading package info for %s", pkgName)

	return pkg
}

func readLoad(pkgName string) Package {
	pkgDispName := bolden(pkgName)

	log(1, "Reading package info for %s...", pkgDispName)
	pkgFile := readFile(infoPath+pkgName+".json", "An error occurred while reading package %s", pkgDispName)
	pkg := loadPkg(pkgFile, pkgName)

	return pkg
}
