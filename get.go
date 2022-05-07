package main

import (
	"os"
	"strconv"
	"strings"
)

func writePkg(pkgName string, pkgFile string) {
	newFile(infoPath+pkgName+".json", pkgFile, "An error occurred while writing package information for %s", pkgName)
}

func findPkg(pkgName string) string {
	log(1, "Finding package %s...", bolden(pkgName))

	urls := parseSources()
	validInfos, validUrls := make([]string, 0), make([]string, 0)

	log(1, "Checking urls length...")

	if urlsLen := len(urls); urlsLen <= 0 {
		errorLogRaw("You don't have any sources defined in %s", bolden(configPath+"sources.txt"))
		os.Exit(1)
	} else if urlsLen == 1 {
		debugLog("Only one source defined in %s. Using that source.", bolden(configPath+"sources.txt"))
		log(1, "Getting package info for %s...", bolden(pkgName))

		pkgURL := parseURL(urls[0], false) + "/" + pkgName + ".json"

		log(1, "Getting info from %s...", bolden(pkgURL))

		pkgFile, statusCode, err := viewFile(pkgURL)

		errorLog(err, "An error occurred while getting package information for %s", bolden(pkgName))
		debugLog("Status code: %d", statusCode)

		if checkFor404(statusCode, pkgName) {
			errorLogRaw("Package %s not found", bolden(pkgName))
			os.Exit(1)
		}

		errorLog(err, "An error occurred while getting package information for %s", bolden(pkgName))

		return pkgFile
	}

	rawLog("\n")

	for _, url := range urls {
		pkgURL := parseURL(url, false) + "/" + pkgName + ".json"

		log(1, "Checking %s for package info...", bolden(pkgURL))

		infoFile, statusCode, err := viewFile(pkgURL)

		errorLog(err, "An error occurred while getting package information for %s", bolden(pkgName))
		debugLog("Status code: %d", statusCode)

		if checkFor404(statusCode, pkgName) {
			log(3, "Not found in %s", bolden(pkgURL))
			rawLog("\n")

			continue
		}

		errorLog(err, "An error occurred while getting package information for %s", bolden(pkgName))

		log(0, "Found %s in %s!", bolden(pkgName), bolden(pkgURL))
		log(1, "Saving valid info & url...")
		rawLog("\n")

		validInfos = append(validInfos, infoFile)
		validUrls = append(validUrls, pkgURL)
	}

	log(1, "Checking valid info...")

	lenValidInfos := len(validInfos)

	switch {
	case lenValidInfos < 1:
		errorLogRaw("Package %s not found in any repo", bolden(pkgName))
		os.Exit(1)
	case lenValidInfos == 1:
		log(1, "Only 1 valid repo found. Using that repo...")

		return validInfos[0]
	default:
		log(1, "Multiple packages found. Please choose one:")

		for i, url := range validUrls {
			log(1, "%d) %s - %s", i, bolden(url), repoLabel(url, false))
		}

		choice := input("0", "Number between 0 and %d or q to quit", lenValidInfos-1)

		if strings.Contains(choice, "q") {
			os.Exit(1)
		} else {
			convChoice, err := strconv.Atoi(choice)
			errorLog(err, "An error occurred while converting choice to int")

			return validInfos[convChoice]
		}
	}

	return ""
}

func getPkgFromNet(pkgName string) (Package, string) {
	packageFile := findPkg(pkgName)

	pkg := loadPkg(packageFile, pkgName)

	return pkg, packageFile
}

func getPkgFromURL(pkgName string, url string) string {
	packageFile, statusCode, err := viewFile(url)
	errorLog(err, "An error occurred while getting package information for %s", bolden(pkgName))

	if checkFor404(statusCode, pkgName) {
		errorLogRaw("The info URL provided from %s does not exist", bolden(pkgName))
	}

	return packageFile
}

func doDirectDownload(pkg Package, pkgName string, srcPath string) {
	pkgDispName := bolden(pkgName)

	log(1, "Making sure %s is not already downloaded...", pkgDispName)
	delPath(false, srcPath+pkg.Name, "An error occurred while deleting temporary downloaded files for %s", pkgName)

	log(1, "Getting download URL for %s", pkgDispName)

	url := getDownloadURL(pkg)

	log(1, "Making directory for %s...", pkgDispName)
	newDir(srcPath+pkg.Name, "An error occurred while creating temporary directory for %s", pkgName)

	log(1, "Downloading file for %s from %s...", pkgDispName, bolden(url))

	nameOfFile := srcPath + pkg.Name + "/" + pkg.Name

	debugLog("Downloading and saving to %s", bolden(nameOfFile))
	downloadFileWithProg(nameOfFile, url, "An error occurred while downloading file for %s", pkgName)
}

func getPkgInfo(pkgName string, isURL bool) string {
	var pkgFile string

	switch {
	case isURL: // Run this if a URL is selected
		log(1, "Reading info from direct URL...")

		parsedURL := parseURL(pkgName, false)
		raw, statusCode, err := viewFile(parsedURL)
		pkgFile = raw

		errorLog(err, "An error occurred while getting info from %s", bolden(pkgName))

		if checkFor404(statusCode, pkgName) {
			errorLogRaw("Package %s not found", bolden(pkgName))
			os.Exit(1)
		}
	case strings.HasSuffix(pkgName, ".json"): // Run this if a file is selected
		log(1, "Reading info from file...")

		pkgFile = readFile(pkgName, "An error occurred while reading %s", bolden(pkgName))
	default: // Run this to read from repos
		log(1, "Reading info from official repositories...")

		pkgFile = findPkg(pkgName)
	}

	return pkgFile
}
