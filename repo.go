package main

import (
	"os"
	"strings"
)

func parseURL(url string, silent bool) string {
	if !silent {
		log(1, "Parsing URL %s...", bolden(url))
	}

	repoURL := url
	repoURL = strings.TrimSpace(repoURL)
	repoURL = strings.TrimSuffix(repoURL, "/")

	if split := strings.Split(repoURL, "//"); strings.Contains(repoURL, "://") {
		repoURL = "http://" + strings.Join(split[1:], "://")
	} else {
		repoURL = "http://" + repoURL
	}

	debugLog("Parsed URL: %s", repoURL)

	return repoURL
}

func readSources() ([]string, string) {
	log(1, "Reading sources file...")

	raw := readFile(configPath+"sources.txt", "An error occurred while reading sources file")

	log(1, "Parsing sources file...")

	trimmed := strings.TrimSpace(raw)

	return strings.Split(trimmed, "\n"), trimmed
}

func stripSources(sourcesFile string, noParse bool) ([]string, string) {
	log(1, "Stripping sources file of comments...")

	final := []string{}
	finalStr := ""

	for _, line := range strings.Split(strings.TrimSpace(sourcesFile), "\n") {
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		var parsedLine string
		if noParse {
			parsedLine = line
		} else {
			parsedLine = parseURL(line, true)
		}

		finalStr += parsedLine + "\n"
		final = append(final, parsedLine)
	}

	return final, finalStr
}

func saveChanges(sourcesFile string) {
	log(1, "Saving changes...")
	newFile(configPath+"sources.txt", sourcesFile, "An error occurred while saving changes to sources file")
}

func addRepo(repoLink string) {
	if !isURL(repoLink) {
		if force {
			log(3, "Invalid url, but continuing because force is set to true.")
		} else {
			log(4, "Invalid url: %s.", bolden(repoLink))
			os.Exit(1)
		}
	}

	_, sourcesFile := readSources()
	_, stripped := stripSources(sourcesFile, true)

	if strings.Contains(stripped, repoLink) {
		if force {
			log(3, "Repo %s already exists in sources file, but continuing because force is set to true.", bolden(repoLink))
		} else {
			errorLogRaw("Repo %s already exists in sources file", bolden(repoLink))
			os.Exit(1)
		}
	}

	log(1, "Appending %s to sources file...", bolden(repoLink))
	sourcesFile = strings.TrimSpace(sourcesFile + "\n" + repoLink)
	saveChanges(sourcesFile)
}

func rmRepo(repoLink string) {
	repos, _ := readSources()

	log(1, "Removing %s from sources file...", bolden(repoLink))

	final := ""

	for i, repo := range repos {
		if repo == repoLink {
			debugLog("Match found at index %d.", i)

			continue
		}

		final = final + repo + "\n"
	}

	sourcesFile := strings.TrimSpace(final)

	saveChanges(sourcesFile)
}

func listRepos() {
	rawRepos, _ := readSources()
	repos, _ := stripSources(strings.Join(rawRepos, "\n"), true)

	log(1, "Repos:")

	for _, repo := range repos {
		rawLog("        %s - %s\n", bolden(repo), repoLabel(repo, false))
	}
}

func repoLabel(repo string, includeLink bool) string {
	prefixes := [][]string{
		{"http://raw.githubusercontent.com/talwat/indiepkg/main/packages/linux-only/", textCol["BLUE"] + "(Linux only)" + RESETCOL},
		{"http://raw.githubusercontent.com/talwat/indiepkg/main/packages/bin/", textCol["VIOLET"] + "(Binary package)" + RESETCOL},
		{"http://raw.githubusercontent.com/talwat/indiepkg/main/", textCol["CYAN"] + "(Official repo)" + RESETCOL},
		{"http://raw.githubusercontent.com/talwat/indiepkg/", textCol["BLUE"] + "(Other branch)" + RESETCOL},
	}

	for k := range prefixes {
		if strings.HasPrefix(parseURL(repo, true), prefixes[k][0]) {
			return prefixes[k][1]
		}
	}

	if includeLink {
		return textCol["YELLOW"] + "(Third party repo: " + repo + ")" + RESETCOL
	}

	return textCol["YELLOW"] + "(Third party repo)" + RESETCOL
}
