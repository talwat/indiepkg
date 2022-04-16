package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

func readSources() string {
	log(1, "Reading sources file...")
	raw := readFile(configPath+"sources.txt", "An error occurred while reading sources file")
	log(1, "Parsing sources file...")
	trimmed := strings.TrimSpace(raw)
	split := strings.Split(trimmed, "\n")
	final := ""
	for _, line := range split {
		if !strings.HasPrefix(line, "#") && strings.TrimSpace(line) != "" {
			final += line + "\n"
		}
	}

	return strings.TrimSpace(final)
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

	sourcesFile := readSources()

	if strings.Contains(sourcesFile, "\n"+repoLink) {
		if force {
			log(3, "Repo %s already exists in sources file, but continuing because force is set to true.", bolden(repoLink))
		} else {
			log(4, "Repo %s already exists in sources file.", bolden(repoLink))
			os.Exit(1)
		}
	}

	log(1, "Appending %s to sources file...", bolden(repoLink))
	sourcesFile = sourcesFile + "\n" + repoLink
	saveChanges(sourcesFile)
}

func rmRepo(repoLink string) {
	sourcesFile := readSources()
	log(1, "Removing %s from sources file...", bolden(repoLink))
	repos := strings.Split(sourcesFile, "\n")
	for i, repo := range repos {
		if repo == repoLink {
			repos[i] = ""
			debugLog("Match found at index %d.", i)
		}
	}

	sourcesFile = strings.Join(repos, "\n")
	saveChanges(sourcesFile)
}

func listRepos() {
	sourcesFile := readSources()
	repos := strings.Split(sourcesFile, "\n")
	log(1, "Repos:")
	for _, repo := range repos {
		fmt.Printf("        %s - %s\n", bolden(repo), replaceRepo(repo))
	}
}

func replaceRepo(repo string) string {
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
