package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type GHFile struct {
	Name string
	Path string
	URL  string
	Repo string
}

func sendGithubRequest(url string) (string, http.Header) {
	tokenLen := len(config.Github.Token)
	if tokenLen >= 4 && config.Github.Username != "" {
		debugLog(
			"Sending request to %s with username %s and token (last 4 digits) %s",
			bolden(url), bolden(config.Github.Username),
			bolden(
				config.Github.Token[len(config.Github.Token)-4:], // Get last 4 digits
			),
		)
	} else {
		debugLog("Invalid/default credentials.")
	}

	errMsgAdded := "An error occurred while getting information from the github API. URL: " + bolden(url)
	resp, err := makeGithubReq(url)

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

		return "", nil
	}

	debugLog("Github API requests limit: %s.", resp.Header.Get("x-ratelimit-limit"))
	debugLog("Github API requests remaining: %s.", resp.Header.Get("x-ratelimit-remaining"))

	return string(final), resp.Header
}

func getAllPkgsFromGh() ([]GHFile, http.Header) {
	urls := parseSources()
	final := make([]GHFile, 0)

	var headers http.Header

	convertURL := func(url string) string {
		apiLink := strings.ReplaceAll(url, "raw.githubusercontent.com", "api.github.com/repos")
		split := strings.Split(apiLink, "/")
		index := 6
		inserted := split[:index]
		inserted = append(inserted, split[index:]...)
		branch := split[index]
		inserted[index] = "contents"
		trimmed := strings.TrimSuffix(strings.Join(inserted, "/"), "/")

		return trimmed + "?ref=" + branch
	}

	for _, url := range urls {
		url = parseURL(url, true)
		if !strings.HasPrefix(url, "http://raw.githubusercontent.com") {
			log(3, "Non-github repositories can't be queried. Repo: %s", url)

			continue
		}

		convURL := convertURL(url)
		debugLog("URL: %s", convURL)
		response, h := sendGithubRequest(convURL)
		headers = h

		var files []GHFile
		err := json.Unmarshal([]byte(response), &files)
		errorLog(err, "An error occurred while parsing package list")

		for _, file := range files {
			if !strings.HasSuffix(file.Name, ".json") {
				continue
			}

			file.Name = strings.TrimSuffix(file.Name, ".json")
			file.Repo = url
			final = append(final, file)
		}
	}

	return final, headers
}

func getPkgFromGh(query string) ([]GHFile, http.Header) {
	files, headers := getAllPkgsFromGh()
	matches := make([]GHFile, 0)

	for _, file := range files {
		if strings.Contains(file.Name, query) {
			file.Repo = repoLabel(file.Repo, true)
			matches = append(matches, file)
		}
	}

	if len(matches) == 0 {
		log(4, "No matches found.")
		os.Exit(1)
	}

	return matches, headers
}

func getRepoInfo(author string, repo string) {
	type GithubRepoInfo struct {
		Name        string
		Description string
		Owner       struct {
			Login string
		}
		CloneURL string `json:"clone_url"`
		Language string
		License  struct {
			SpdxID string `json:"spdx_id"`
		}
	}

	log(1, "Getting repository info from the Github API...")

	url := "https://api.github.com/repos/" + author + "/" + repo
	response, _ := sendGithubRequest(url)

	if response == "" {
		errorLogRaw("The Github API returned an empty response. This may be because you are getting rate limited. URL: %s", url)
		os.Exit(1)
	}

	debugLog("Response:\n%s", response)

	log(1, "Parsing response...")

	var repoInfo GithubRepoInfo
	err := json.Unmarshal([]byte(response), &repoInfo)
	errorLog(err, "An error occurred while parsing the response.")

	path := "samples/templates/" + strings.ToLower(repoInfo.Language) + ".json"
	if pathExists(path, "An error occurred while checking for language template.") {
		log(1, "Using language template: %s", path)
	} else {
		path = "samples/basic.json"
		log(1, "Using default language template")
	}

	log(1, "Parsing description...")

	repoInfo.Description = strings.ToUpper(string(repoInfo.Description[0])) + repoInfo.Description[1:]

	if !strings.HasSuffix(repoInfo.Description, ".") {
		repoInfo.Description += "."
	}

	file := readFile(path, "An error occurred while reading the sample file for %s.", repoInfo.Language)

	finalFile := file
	finalFile = strings.ReplaceAll(finalFile, "nameofpkg", repoInfo.Name)
	finalFile = strings.ReplaceAll(finalFile, "mypkglicense", repoInfo.License.SpdxID)
	finalFile = strings.ReplaceAll(finalFile, "mypkgdescription", repoInfo.Description)
	finalFile = strings.ReplaceAll(finalFile, "mypkgauthor", repoInfo.Owner.Login)
	finalFile = strings.ReplaceAll(finalFile, "git url", repoInfo.CloneURL)
	finalFile = strings.ReplaceAll(finalFile, "nameoflang", repoInfo.Language)

	log(1, "Writing generated package info...")
	newFile(repoInfo.Name+".json", finalFile, "An error occurred while writing generated package info file")

	log(0, "Successfully generated package info for %s.", repoInfo.Name)
}
