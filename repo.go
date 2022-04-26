package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

func readSources() ([]string, string) {
	log(1, "Reading sources file...")
	raw := readFile(configPath+"sources.txt", "An error occurred while reading sources file")
	log(1, "Parsing sources file...")
	trimmed := strings.TrimSpace(raw)
	return strings.Split(trimmed, "\n"), trimmed
}

func saveChanges(sourcesFile string) {
	log(1, "Saving changes...")
	newFile(configPath+"sources.txt", sourcesFile, "An error occurred while saving changes to sources file")
}

func addRepo(repoLink string) {
	_, err := url.ParseRequestURI(repoLink)
	if err != nil {
		if force {
			log(3, "Invalid url, but continuing because force is set to true.")
		} else {
			log(4, "Invalid url: %s.", bolden(repoLink))
			os.Exit(1)
		}
	}

	_, sourcesFile := readSources()

	if strings.Contains(sourcesFile, "\n"+repoLink) {
		if force {
			log(3, "Repo %s already exists in sources file, but continuing because force is set to true.", bolden(repoLink))
		} else {
			errorLogRaw("Repo %s already exists in sources file", bolden(repoLink))
			os.Exit(1)
		}
	}

	log(1, "Appending %s to sources file...", bolden(repoLink))
	sourcesFile = sourcesFile + "\n" + repoLink
	saveChanges(sourcesFile)
}

func rmRepo(repoLink string) {
	repos, _ := readSources()
	log(1, "Removing %s from sources file...", bolden(repoLink))

	for i, repo := range repos {
		if repo == repoLink {
			repos[i] = ""
			debugLog("Match found at index %d.", i)
		}
	}

	sourcesFile := strings.Join(repos, "\n")
	saveChanges(sourcesFile)
}

func listRepos() {
	repos, _ := readSources()
	log(1, "Repos:")
	for _, repo := range repos {
		fmt.Printf("        %s - %s\n", bolden(repo), repoLabel(repo))
	}
}

func repoLabel(repo string) string {
	var m [][]string = [][]string{
		{"https://raw.githubusercontent.com/talwat/indiepkg/main/packages/linux-only/", textCol["BLUE"] + "(Linux only)" + RESETCOL},
		{"https://raw.githubusercontent.com/talwat/indiepkg/main/packages/bin/", textCol["VIOLET"] + "(Binary package)" + RESETCOL},
		{"https://raw.githubusercontent.com/talwat/indiepkg/main/", textCol["CYAN"] + "(Official repo)" + RESETCOL},
		{"https://raw.githubusercontent.com/talwat/indiepkg/", textCol["BLUE"] + "(Other branch)" + RESETCOL},
	}

	for k := range m {
		if strings.HasPrefix(repo, m[k][0]) {
			return m[k][1]
		}
	}
	return textCol["YELLOW"] + "(Third party repo: " + repo + ")" + RESETCOL
}
