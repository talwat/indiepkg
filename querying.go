package main

import (
	"encoding/json"
	"os"
	"strings"
)

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
		branch := split[index]
		inserted[index] = "contents"
		trimmed := strings.TrimSuffix(strings.Join(inserted, "/"), "/")
		return trimmed + "?ref=" + branch
	}

	for _, url := range urls {
		replaceRepo := func(repo string) string {
			var m map[string]string = map[string]string{
				"https://raw.githubusercontent.com/talwat/indiepkg/main/packages/linux-only/": textCol["BLUE"] + "(Linux only)" + RESETCOL,
				"https://raw.githubusercontent.com/talwat/indiepkg/":                          textCol["CYAN"] + "(Official repo)" + RESETCOL,
			}

			for k, v := range m {
				if strings.HasPrefix(repo, k) {
					return v
				}
			}
			return textCol["YELLOW"] + "(Third party repo: " + repo + ")" + RESETCOL
		}

		if !strings.HasPrefix(url, "https://raw.githubusercontent.com") {
			log(3, "Non-github repositories can't be queried. Repo: %s", url)
			continue
		}

		convUrl := convertUrl(url)
		debugLog("URL: %s", convUrl)
		r, _ := viewFile(convUrl, "An error occurred while getting package list")

		if r == "" {
			log(4, "The Github API returned an empty response. This may be because you are getting rate limited. URL: %s", convUrl)
			os.Exit(1)
		}

		var files []GH_File
		err := json.Unmarshal([]byte(r), &files)
		errorLog(err, 4, "An error occurred while parsing package list")

		for _, file := range files {
			file.Name = strings.TrimSuffix(file.Name, ".json")
			if strings.Contains(file.Name, query) {
				file.Repo = replaceRepo(url)
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
