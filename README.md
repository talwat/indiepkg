# IndiePKG

[![Go Report Card](https://goreportcard.com/badge/github.com/talwat/indiepkg)](https://goreportcard.com/report/github.com/talwat/indiepkg)
[![GitHub stars](https://img.shields.io/github/stars/talwat/indiepkg)](https://github.com/talwat/indiepkg/stargazers)
![GitHub watchers](https://img.shields.io/github/watchers/talwat/indiepkg)
[![GitHub forks](https://img.shields.io/github/forks/talwat/indiepkg)](https://github.com/talwat/indiepkg/network)
![GitHub Action](https://img.shields.io/github/workflow/status/talwat/indiepkg/golangci-lint)
[![GitHub license](https://img.shields.io/github/license/talwat/indiepkg)](https://github.com/talwat/indiepkg)
![GitHub last commit](https://img.shields.io/github/last-commit/talwat/indiepkg)
![GitHub contributors](https://img.shields.io/github/contributors/talwat/indiepkg)
[![GitHub issues](https://img.shields.io/github/issues/talwat/indiepkg)](https://github.com/talwat/indiepkg/issues)
![GitHub closed pull requests](https://img.shields.io/github/issues-pr-closed/talwat/indiepkg)

A package manager written in go for small CLI programs. It is available on GNU/Linux and macOS.

## Notice

IndiePKG is **NOT** ready for use yet. It's still early software.

However, if you would like to submit issues or PR's, you are **more** than welcome to.

## What is IndiePKG?

IndiePKG is mainly for small simple CLI and TUI programs. Most of them are just for fun, such as **cmatrix**, while others have a bit more utility such as **btop**.

IndiePKG uses **git** to install packages, and everything is compiled **from source**. This means that while there aren't any versions, it does mean that you get the absolute latest software.

It's also much simpler than your standard package manager, and if a package installation goes wrong you don't have to worry about all your packages failing, because you can super easily remove it.

### Pros

- Easy to use with simple commands.
- Explains what it's doing at every step.
- Simple compared to other package managers, and easy to repair if something goes wrong.
- Clones directly from the package's git repository, giving you the latest version with no manual intervention by a maintainer.
- Easy to submit to, with no notability requirements.
- Installs **everything** locally, not interfering with any system components and not requiring root privileges.

### Cons

- Doesn't install & manage dependencies.
- Slower than other package managers due to needing to compile source code.

## Installation

While IndiePKG doesn't have an install script yet, you can still try it out by cloning the repository and building it from source.
You will need:

- Git
- Go 1.17+
- Make

```bash
git clone -b testing https://github.com/talwat/indiepkg.git
cd indiepkg
make
make install
```

## Making & submitting packages

If you want to make a package, you should look at [PACKAGES.md](PACKAGES.md)

## Basic usage

You can run `indiepkg install <packages>` to install a package.

If you want to uninstall a package, you can run `indiepkg uninstall <packages>`.

`indiepkg upgrade [packages]` will pull the latest changes and recompile packages. If you don't specify any packages, it will upgrade all packages.

`indiepkg update [packages]` will update the information for your installed packages. This command doesn't need to be ran frequently at all, but it's best to run it every now and then.

## Supported operating systems

IndiePKG officially works on **GNU/Linux** and **macOS**.

It will work on any other unix-like operating system but many packages won't compile properly. If you are interested in supporting a new operating system, please [open an issue](https://github.com/talwat/indiepkg/issues/new/choose) and we can discuss making a new third party repository.

IndiePKG is **completely** broken on Windows. IndiePKG is overall not meant to run on Windows.

## Branches

IndiePKG has 3 branches to install from, **testing** *(recommended)*, **stable**, and **main**.

The **testing** branch is a good balance of new and stable, and is updated whenever all Github actions pass.

The **stable** branch is only updated whenever a release is made, and is not recommended for use *yet* because it will sometimes not work due to changes in the package files.

The **main** branch is the default git branch, and contains the **newest** code. This is not recommended for use because it has a higher chance of being broken.
