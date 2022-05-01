package main

import (
	"os"
	"runtime"
)

func getDownloadURL(pkg Package) string {
	debugLog("GOOS: %s. GOARCH: %s", runtime.GOOS, runtime.GOARCH)

	var url string

	checkForAll := func(pkg Package) string {
		if pkg.Download["all"] != nil {
			return pkg.Download["all"].(string)
		}

		return "nil"
	}

	checkForAllArch := func(pkg Package) string {
		if pkg.Download[runtime.GOOS].(map[string]interface{})["all"] != nil {
			return pkg.Download[runtime.GOOS].(map[string]interface{})["all"].(string)
		}

		return "nil"
	}

	if pkg.Download[runtime.GOOS] == nil {
		url = checkForAll(pkg)
	} else if pkg.Download[runtime.GOOS].(map[string]interface{})[runtime.GOARCH] == nil {
		url = checkForAllArch(pkg)
	} else {
		url = pkg.Download[runtime.GOOS].(map[string]interface{})[runtime.GOARCH].(string)
	}

	if url == "nil" {
		errorLogRaw("Unsupported OS or architecture")
		os.Exit(1)
	}

	return url
}
