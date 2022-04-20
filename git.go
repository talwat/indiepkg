package main

import (
	"strings"
)

func pullSrcRepo(silent bool) bool {
	output, err := runCommand(indiePkgSrcDir, "git", "pull", "--no-tags", "--depth", "1")
	debugLog("Git output from pull: %s", output)
	errorLog(err, 4, "An error occurred while pulling source code for IndiePKG")

	if strings.Contains(output, "Already up to date.") {
		if !silent {
			log(0, "IndiePKG already up to date")
		}
		return true
	}

	return false
}

func cloneSrcRepo() {
	log(1, "Cloning IndiePKG source...")
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
		"src",
	)
}

func clonePkgRepo(pkg Package, cloneDir string) {
	log(1, "Cloning source code for %s...", bolden(pkg.Name))

	if pkg.Branch == "" {
		log(1, "Getting branch %s...", bolden(pkg.Name))
		runCommandRealTime(
			cloneDir,
			"git",
			"clone",
			"--no-tags",
			"--progress",
			"--depth",
			"1",
			pkg.Url,
			pkg.Name,
		)
	} else {
		log(1, "Getting branch %s...", bolden(pkg.Name))
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
			pkg.Url,
			pkg.Name,
		)
	}
}

func pullPkgRepo(pkgName string) (bool, bool) {
	output, err := runCommand(srcPath+pkgName, "git", "pull", "--no-tags", "--depth", "1")

	debugLog("Git output from pull:\n%s", output)

	if strings.Contains(output, "Already up to date.") {
		return true, false
	} else if strings.Contains(output, "not a git repository") {
		return false, true
	}

	errorLog(err, 4, "An error occurred while pulling source code for %s", bolden(pkgName))

	return false, false
}
