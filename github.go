package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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

func sendGithubRequest(url string) (string, http.Header) {
	debugLog(
		"Sending request to %s with username %s and token (last 4 digits) %s",
		bolden(url), bolden(config.Github.Username),
		bolden(
			config.Github.Token[len(config.Github.Token)-4:], // Get last 4 digits
		),
	)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	errorLog(err, "An error occurred while creating the GET request. URL: %s", url)

	req.SetBasicAuth(config.Github.Username, config.Github.Token)
	resp, err := client.Do(req)
	errMsgAdded := "An error occurred while getting information from the github API. URL: " + bolden(url)
	errorLog(err, errMsgAdded)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errorLogRaw("The Github API returned an error. Status code: %s", bolden(resp.StatusCode))
		return "", nil
	}

	final, err := ioutil.ReadAll(resp.Body)
	errorLog(err, errMsgAdded)

	if string(final) == "" {
		errorLogRaw("The Github API returned an empty response. This may be because you are getting rate limited. If you would like to improve the amount of requests you have, specify the [github] fields. URL: %s", bolden(url))
		os.Exit(1)
	}

	debugLog("Github API requests limit: %s.", resp.Header.Get("x-ratelimit-limit"))
	debugLog("Github API requests remaining: %s.", resp.Header.Get("x-ratelimit-remaining"))
	return string(final), resp.Header
}

func getPkgFromGh(query string) ([]GH_File, http.Header) {
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

	var h http.Header

	for _, url := range urls {
		if !strings.HasPrefix(url, "https://raw.githubusercontent.com") {
			log(3, "Non-github repositories can't be queried. Repo: %s", url)
			continue
		}

		convUrl := convertUrl(url)
		debugLog("URL: %s", convUrl)
		r, headers := sendGithubRequest(convUrl)
		h = headers

		var files []GH_File
		err := json.Unmarshal([]byte(r), &files)
		errorLog(err, "An error occurred while parsing package list")

		for _, file := range files {
			if !strings.HasSuffix(file.Name, ".json") {
				continue
			}
			file.Name = strings.TrimSuffix(file.Name, ".json")
			if strings.Contains(file.Name, query) {
				file.Repo = repoLabel(url)
				matches = append(matches, file)
			}
		}
	}

	if len(matches) == 0 {
		log(4, "No matches found.")
		os.Exit(1)
	}

	return matches, h
}

func getRepoInfo(author string, repo string) {
	type Github_Repo_Info struct {
		Name        string
		Description string
		Owner       struct {
			Login string
		}
		Clone_url string
		Language  string
		License   struct {
			Spdx_id string
		}
	}

	log(1, "Getting repository info from the Github API...")
	url := "https://api.github.com/repos/" + author + "/" + repo
	r, _ := viewFile(url, "An error occurred while getting info from the Github API.")

	if r == "" {
		errorLogRaw("The Github API returned an empty response. This may be because you are getting rate limited. URL: %s", url)
		os.Exit(1)
	}

	debugLog("Response:\n%s", r)

	log(1, "Parsing response...")
	var repo_info Github_Repo_Info
	err := json.Unmarshal([]byte(r), &repo_info)
	errorLog(err, "An error occurred while parsing the response.")

	path := "samples/templates/" + strings.ToLower(repo_info.Language) + ".json"
	if pathExists(path, "An error occurred while checking for language template.") {
		log(1, "Using language template: %s", path)
	} else {
		path = "samples/basic.json"
		log(1, "Using default language template")
	}
	file := readFile(path, "An error occurred while reading the sample file for %s.", repo_info.Language)

	finalFile := file
	finalFile = strings.ReplaceAll(finalFile, "nameofpkg", repo_info.Name)
	finalFile = strings.ReplaceAll(finalFile, "mypkglicense", repo_info.License.Spdx_id)
	finalFile = strings.ReplaceAll(finalFile, "mypkgdescription", repo_info.Description)
	finalFile = strings.ReplaceAll(finalFile, "mypkgauthor", repo_info.Owner.Login)
	finalFile = strings.ReplaceAll(finalFile, "git url", repo_info.Clone_url)

	log(1, "Writing generated package info...")
	newFile(repo_info.Name+".json", finalFile, "An error occurred while writing generated package info file")

	log(0, "Successfully generated package info for %s.", repo_info.Name)
}
