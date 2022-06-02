package main

import (
	"strings"
)

// Pulls source for indiepkg
func pullSrcRepo(silent bool) bool {
	output, err := runCommand(indiePkgSrcPath, "git", "pull", "--no-tags")
	errorLog(err, "An error occurred while pulling source code for IndiePKG. Git output:\n%s", output)

	if strings.Contains(output, "Already up to date.") {
		if !silent {
			if force {
				log(3, "IndiePKG already up to date, but force is on, so continuing.")

				return false
			}

			log(0, "IndiePKG already up to date")
		}

		return true
	}

	rawLog(output + "\n")

	return false
}

// Clones indiepkg source
func cloneSrcRepo() {
	log(1, "Cloning IndiePKG source with branch %s...", config.Updating.Branch)
	runCommandRealTime(
		mainPath,
		"git",
		"clone",
		"--branch",
		config.Updating.Branch,
		"--progress",
		"--no-tags",
		"--depth",
		"1",
		"https://github.com/talwat/indiepkg.git",
		indiePkgSrcPath,
	)
}

// Clones source code for a package
func clonePkgRepo(pkg Package, cloneDir string) {
	log(1, "Cloning source code for %s...", bolden(pkg.Name))

	if pkg.Branch == "" {
		runCommandRealTime(
			cloneDir,
			"git",
			"clone",
			"--no-tags",
			"--progress",
			"--depth",
			"1",
			pkg.URL,
			pkg.Name,
		)
	} else {
		log(1, "Getting branch %s...", pkg.Branch)
		debugLog("Cloning to %s on branch %s.", cloneDir+pkg.Name, pkg.Branch)
		runCommandRealTime(
			cloneDir,
			"git",
			"clone",
			"--branch",
			pkg.Branch,
			"--no-tags",
			"--progress",
			"--depth",
			"1",
			pkg.URL,
			pkg.Name,
		)
	}
}

// Pulls source code for a package
func pullPkgRepo(pkgName string) (bool, bool) {
	output, err := runCommand(srcPath+pkgName, "git", "pull", "--no-tags")

	debugLog("Git output from pull:\n%s", output)

	// Check git output
	if strings.Contains(output, "Already up to date.") {
		return true, false
	} else if strings.Contains(output, "not a git repository") {
		return false, true
	}

	errorLog(err, "An error occurred while pulling source code for %s", bolden(pkgName))

	return false, false
}
