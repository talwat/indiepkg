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

	if pkg.Download[runtime.GOOS] == nil {
		url = checkForAll(pkg)
	} else if pkg.Download[runtime.GOOS].(map[string]interface{})[runtime.GOARCH] == nil {
		url = checkForAll(pkg)
	} else {
		url = pkg.Download[runtime.GOOS].(map[string]interface{})[runtime.GOARCH].(string)
	}

	if url == "nil" {
		if pkg.Download["all"] != nil {
			url = pkg.Download["all"].(string)
		} else {
			errorLogRaw("Unsupported OS or architecture")
			os.Exit(1)
		}
	}

	return url
}
