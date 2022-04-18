package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func cloneRepo(pkg Package, cloneDir string) {
	log(1, "Cloning source code for %s...", bolden(pkg.Name))
	if pkg.Branch == "" {
		debugLog("Cloning to %s", bolden(cloneDir+pkg.Name))

		_, err := git.PlainClone(cloneDir+pkg.Name, false, &git.CloneOptions{
			URL:      pkg.Url,
			Progress: os.Stdout,
			Depth:    1,
			Tags:     git.NoTags,
		})

		errorLog(err, 4, "An error occurred while cloning repository for %s", bolden(pkg.Name))
	} else {
		log(1, "Getting branch %s%s%s...", textFx["BOLD"], pkg.Branch, RESETCOL)
		debugLog("Cloning to %s on branch %s", cloneDir+pkg.Name, pkg.Branch)
		_, err := git.PlainClone(cloneDir+pkg.Name, false, &git.CloneOptions{
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
	if err != nil {
		return err
	}

	log(1, "Getting git worktree...")
	w, err := r.Worktree()
	errorLog(err, 4, "An error occurred while getting worktree for %s", bolden(pkgName))

	log(1, "Getting head branch...")
	b, err := r.Head()
	ref := b.Name().String()
	errorLog(err, 4, "An error occurred while getting head for %s", bolden(pkgName))

	log(1, "Pulling %s with ref %s", bolden(srcPath+pkgName), bolden(b.Name().String()))
	err = w.Pull(&git.PullOptions{
		RemoteName:    "origin",
		Progress:      os.Stdout,
		ReferenceName: plumbing.ReferenceName(ref),
	})

	return err
}
