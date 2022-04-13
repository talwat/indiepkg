package main

import (
	"net/url"
	"os"
	"strings"
)

func readSources() string {
	log(1, "Reading sources file...")
	return strings.TrimSpace(readFile(configPath+"sources.txt", "An error occurred while reading sources file"))
}

func saveChanges(sourcesFile string) {
	log(1, "Saving changes...")
	newFile(configPath+"sources.txt", sourcesFile, "An error occurred while saving changes to sources file")
}

func addRepo(repoLink string) {
	_, err := url.ParseRequestURI(repoLink)
	if err != nil && !force {
		log(4, "Invalid url: %s.", bolden(repoLink))
		os.Exit(1)
	} else if force {
		log(3, "Invalid url, but continuing because force is set to true.")
	}

	sourcesFile := readSources()

	if strings.Contains(sourcesFile, "\n"+repoLink) && !force {
		log(4, "Repo %s already exists in sources file.", bolden(repoLink))
		os.Exit(1)
	} else if force {
		log(3, "Repo %s already exists in sources file, but continuing because force is set to true.", bolden(repoLink))
	}

	log(1, "Appending %s to sources file...", bolden(repoLink))
	sourcesFile = sourcesFile + "\n" + repoLink
	saveChanges(sourcesFile)
}

func rmRepo(repoLink string) {
	sourcesFile := readSources()
	log(1, "Removing %s from sources file...", bolden(repoLink))
	sourcesFile = strings.ReplaceAll(sourcesFile, "\n"+repoLink, "")
	saveChanges(sourcesFile)
}
