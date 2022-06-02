package main

import "strings"

// Copies binaries to their proper location for a specific package
func copyBins(pkg Package, srcPath string) {
	if pkg.Bin == nil { // Safeguard to prevent nil pointer dereference
		return
	} else if len(pkg.Bin.InSource) == 0 { // If package doesn't have binaries, return
		return
	}

	pkgDispName := bolden(pkg.Name)
	binPath := config.Paths.Prefix + "bin/"

	log(1, "Copying files for %s...", pkgDispName)

	for i := range pkg.Bin.InSource {
		srcDir := srcPath + pkg.Name + "/" + pkg.Bin.InSource[i] // Combine package's source path and binary name
		destDir := binPath + pkg.Bin.Installed[i]

		log(1, "Copying %s to %s...", bolden(srcDir), bolden(destDir))
		copyFile(srcDir, destDir)

		log(1, "Making %s executable...", bolden(destDir))
		changePerms(destDir, 0o770)
	}
}

// Copies manpages to their proper location for a specific package
func copyManpages(pkg Package, srcPath string) {
	if len(pkg.Manpages) == 0 { // If package doesn't have manpages, return
		return
	}

	pkgDispName := bolden(pkg.Name)
	manPath := config.Paths.Prefix + "share/man/"

	log(1, "Copying manpages for %s...", pkgDispName)

	for _, manPage := range pkg.Manpages {
		srcFile := srcPath + pkg.Name + "/" + manPage

		// Splitting to get file name
		split := strings.Split(manPage, "/")

		// Splitting and getting extension to put in proper man directory, eg. man1, man3, etc...
		destDir := manPath + "man" + strings.Split(manPage, ".")[1] + "/"
		destFile := destDir + split[len(split)-1]

		log(1, "Making manpage directory %s...", bolden(destDir))
		newDir(destDir, "An error occurred while making making manpage directory for %s", pkgDispName)

		log(1, "Copying manpage %s to %s...", bolden(srcFile), bolden(destDir))
		copyFile(srcFile, destFile)
	}
}
