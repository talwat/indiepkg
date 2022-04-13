package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func writeLoadPkg(pkgName string, pkgFile string, load bool) Package {
	newFile(installedPath+pkgName+".json", pkgFile, "An error occurred while writing package information for %s", pkgName)

	if load {
		var pkg Package

		if load {
			pkg = readLoad(pkgName)
		}

		return pkg
	}

	return Package{}
}

func findPkg(pkgName string) string {
	urls := parseSources()
	var validInfos, validUrls []string
	urlsLen := len(urls)

	if urlsLen <= 0 {
		log(4, "You don't have any sources defined in %s.", bolden(configPath+"sources.txt"))
		os.Exit(1)
	} else if urlsLen == 1 {
		debugLog("Only one source defined in %s. Using that source.", bolden(configPath+"sources.txt"))
		log(1, "Getting package info for %s...", bolden(pkgName))
		pkgFile, err := viewFile(urls[0]+pkgName+".json", "An error occurred while getting package information for %s", pkgName)

		if errIs404(err) {
			log(4, "Package %s not found.", bolden(pkgName))
			os.Exit(1)
		}

		errorLog(err, 4, "An error occurred while getting package information for %s", bolden(pkgName))

		return pkgFile
	}

	for _, url := range urls {
		pkgUrl := url + pkgName + ".json"
		log(1, "Checking %s for package info...", bolden(url))
		infoFile, err := viewFile(pkgUrl, "An error occurred while getting package information for %s", pkgName)

		if errIs404(err) {
			continue
		}

		errorLog(err, 4, "An error occurred while getting package information for %s", bolden(pkgName))

		log(1, "Saving valid info & url...")
		validInfos = append(validInfos, infoFile)
		validUrls = append(validUrls, url)
	}

	lenValidInfos := len(validInfos)

	if lenValidInfos < 1 {
		log(4, "Package %s not found in any repo.", bolden(pkgName))
		os.Exit(1)
	} else if lenValidInfos == 1 {
		return validInfos[0]
	} else {
		fmt.Printf("\n")
		log(1, "Multiple packages found. Please choose one:")
		for i, url := range validUrls {
			log(1, "%d) %s", i, bolden(url+pkgName))
		}

		choice := input("0", "Number between 0 and %d or q to quit", lenValidInfos-1)

		if strings.Contains(choice, "q") {
			os.Exit(1)
		} else {
			convChoice, err := strconv.Atoi(choice)
			errorLog(err, 4, "An error occurred while converting choice to int")
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

func downloadPkg(pkgName string, load bool) Package {
	log(1, "Downloading package info for %s...", bolden(pkgName))

	return writeLoadPkg(pkgName, findPkg(pkgName), load)
}

type GH_File struct {
	Name         string
	Path         string
	Url          string
	Html_url     string //blob
	Download_url string //raw
	Repo         string
}

func getPkgFromGh(query string) []GH_File {
	urls := parseSources()
	var matches []GH_File

	convertUrl := func(url string) string {
		apiLink := strings.ReplaceAll(url, "raw.githubusercontent.com", "api.github.com/repos")
		split := strings.Split(apiLink, "/")
		index := 6
		inserted := append(split[:index], split[index:]...)
		inserted[index] = "contents"
		return strings.Join(inserted, "/")
	}

	for _, url := range urls {
		if !strings.HasPrefix(url, "https://raw.githubusercontent.com") {
			log(3, "Non-github repositories can't be queried. Repo: %s", url)
			continue
		}

		convUrl := convertUrl(url)
		debugLog("URL: %s", convUrl)
		r, _ := viewFile(convUrl, "An error occurred while getting package list")
		var files []GH_File
		err := json.Unmarshal([]byte(r), &files)
		errorLog(err, 4, "An error occurred while parsing package list")

		for _, file := range files {
			file.Name = strings.TrimSuffix(file.Name, ".json")
			if strings.Contains(file.Name, query) {
				if strings.HasPrefix(url, "https://raw.githubusercontent.com/talwat/indiepkg/") {
					file.Repo = textCol["CYAN"] + "(official repo)" + RESETCOL
				} else {
					file.Repo = textCol["YELLOW"] + "(3rd party repo: " + url + ")" + RESETCOL
				}

				matches = append(matches, file)
			}
		}

	}
	if len(matches) == 0 {
		log(4, "No matches found.")
		os.Exit(1)
	}

	return matches
}
