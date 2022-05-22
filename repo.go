package main

import (
	"strings"
)

func parseURL(url string, silent bool) string {
	if !silent {
		log(1, "Parsing URL %s...", bolden(url))
	}

	repoURL := url

	// Trim
	repoURL = strings.TrimSpace(repoURL)
	repoURL = strings.TrimSuffix(repoURL, "/")

	// Change protocol to http://
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

	// Iterate through each line
	for _, line := range strings.Split(strings.TrimSpace(sourcesFile), "\n") {
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" { // Skip comments and empty lines
			continue
		}

		parsedLine := line
		if !noParse { // Parse URL if noParse is set to false
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
	if !isValidURL(repoLink) {
		if force {
			log(3, "Invalid url, but continuing because force is set to true.")
		} else {
			errorLogRaw("Invalid url: %s", bolden(repoLink))
		}
	}

	_, sourcesFile := readSources()
	_, stripped := stripSources(sourcesFile, true)

	if strings.Contains(stripped, repoLink) {
		if force {
			log(3, "Repo %s already exists in sources file, but continuing because force is set to true.", bolden(repoLink))
		} else {
			errorLogRaw("Repo %s already exists in sources file", bolden(repoLink))
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
		rawLogf("        %s - %s\n", bolden(repo), repoLabel(repo, false))
	}
}

func repoLabel(repo string, includeLink bool) string {
	prefixes := [][]string{
		{"http://raw.githubusercontent.com/talwat/indiepkg/main/packages/linux-only", textCol.Blue + "(Linux only)" + RESETCOL},
		{"http://raw.githubusercontent.com/talwat/indiepkg/main/packages/bin", textCol.Violet + "(Binary package)" + RESETCOL},
		{"http://raw.githubusercontent.com/talwat/indiepkg/main", textCol.Cyan + "(Official repo)" + RESETCOL},
		{"http://raw.githubusercontent.com/talwat/indiepkg", textCol.Blue + "(Other branch)" + RESETCOL},
	}

	// Iterate through 2D array
	for prefix := range prefixes {
		if strings.HasPrefix(parseURL(repo, true), prefixes[prefix][0]) { // Check prefix matches
			return prefixes[prefix][1]
		}
	}

	if includeLink {
		return textCol.Yellow + "(Third party repo: " + repo + ")" + RESETCOL
	}

	return textCol.Yellow + "(Third party repo)" + RESETCOL
}
