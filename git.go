package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func cloneRepo(pkg Package) {
	log(1, "Cloning source code for %s...", bolden(pkg.Name))
	if pkg.Branch == "" {
		debugLog("Cloning to %s", bolden(srcPath+pkg.Name))

		_, err := git.PlainClone(srcPath+pkg.Name, false, &git.CloneOptions{
			URL:          pkg.Url,
			Progress:     os.Stdout,
			SingleBranch: true,
			Depth:        1,
			Tags:         git.NoTags,
		})

		errorLog(err, 4, "An error occurred while cloning repository for %s", bolden(pkg.Name))
	} else {
		log(1, "Getting branch %s%s%s...", textFx["BOLD"], pkg.Branch, RESETCOL)
		debugLog("Cloning to %s on branch %s", srcPath+pkg.Name, pkg.Branch)
		_, err := git.PlainClone(srcPath+pkg.Name, false, &git.CloneOptions{
			URL:           pkg.Url,
			Progress:      os.Stdout,
			ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", pkg.Branch)),
			SingleBranch:  true,
			Depth:         1,
			Tags:          git.NoTags,
		})

		errorLog(err, 4, "An error occurred while cloning repository for %s", bolden(pkg.Name))
	}
}

func pullRepo(pkgName string) error {
	var err error
	r, err := git.PlainOpen(srcPath + pkgName)
	errorLog(err, 4, "An error occurred while opening repository for %s", bolden(pkgName))

	w, err := r.Worktree()
	errorLog(err, 4, "An error occurred while getting worktree for %s", bolden(pkgName))

	debugLog("Pulling %s", bolden(srcPath+pkgName))
	err = w.Pull(&git.PullOptions{
		RemoteName: "origin",
		Progress:   os.Stdout,
	})

	if err.Error() == "already up-to-date" {
		if force {
			log(3, "%s is already up to date, but force is on, so continuing.", bolden(pkgName))
			return errors.New("continuing because force is on")
		}
		log(0, "%s already up to date.", bolden(pkgName))
	} else {
		errorLog(err, 4, "An error occurred while pulling repository for %s", bolden(pkgName))
	}
	return err
}
