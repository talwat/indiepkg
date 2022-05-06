<!-- markdownlint-disable MD033 MD041 MD013 -->

<div align="center">
    <img src="./assets/logoDarkRound.svg">
    <h1>IndiePKG</h1>
</div>

[![Go Report Card](https://goreportcard.com/badge/github.com/talwat/indiepkg)](https://goreportcard.com/report/github.com/talwat/indiepkg)
[![Total alerts](https://img.shields.io/lgtm/alerts/g/talwat/indiepkg.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/talwat/indiepkg/alerts/)
![GitHub Action](https://img.shields.io/github/workflow/status/talwat/indiepkg/test)
[![GitHub stars](https://img.shields.io/github/stars/talwat/indiepkg)](https://github.com/talwat/indiepkg/stargazers)
![GitHub watchers](https://img.shields.io/github/watchers/talwat/indiepkg)
[![GitHub forks](https://img.shields.io/github/forks/talwat/indiepkg)](https://github.com/talwat/indiepkg/network)
[![GitHub license](https://img.shields.io/github/license/talwat/indiepkg)](https://github.com/talwat/indiepkg)
![GitHub last commit](https://img.shields.io/github/last-commit/talwat/indiepkg)
![GitHub contributors](https://img.shields.io/github/contributors/talwat/indiepkg)
[![GitHub issues](https://img.shields.io/github/issues/talwat/indiepkg)](https://github.com/talwat/indiepkg/issues)
![GitHub closed pull requests](https://img.shields.io/github/issues-pr-closed/talwat/indiepkg)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Ftalwat%2Findiepkg.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Ftalwat%2Findiepkg?ref=badge_shield)

A package manager written in [Go](https://go.dev) for small CLI programs. It is available on [GNU](https://gnu.org)/[Linux](https://kernel.org) and [macOS](https://www.apple.com/macos/).

## Notice

IndiePKG is **NOT** ready for use yet. It's still early software.

However, if you would like to [submit issues](https://github.com/talwat/indiepkg/issues) or [PR's](https://github.com/talwat/indiepkg/pulls), you are **more** than welcome to.

## Table of contents

- [Notice](#notice)
- [Table of contents](#table-of-contents)
- [What is IndiePKG?](#what-is-indiepkg)
  - [Pros](#pros)
  - [Cons](#cons)
- [Installation](#installation)
  - [Dependencies](#dependencies)
  - [Install script](#install-script)
    - [Dependencies for install script](#dependencies-for-install-script)
  - [Manual install](#manual-install)
- [Uninstallation](#uninstallation)
- [Making & submitting packages](#making--submitting-packages)
- [Basic usage](#basic-usage)
- [Supported operating systems](#supported-operating-systems)
  - [List of operating systems](#list-of-operating-systems)
- [Branches](#branches)
- [Libraries used](#libraries-used)
- [FAQ](#faq)
  - [So it's like the AUR?](#so-its-like-the-aur)

## What is IndiePKG?

IndiePKG is mainly for small simple [CLI](https://en.wikipedia.org/wiki/Command-line_interface) and [TUI](https://en.wikipedia.org/wiki/Text-based_user_interface) programs. Most of them are just for fun, such as **cmatrix**, while others have a bit more utility such as **btop**.

IndiePKG uses [git](https://git-scm.com/) to install packages, and everything is compiled **from source**. This means that while there aren't any versions, it does mean that you get the absolute latest software.

It's also much simpler than your standard package manager, and if a package installation goes wrong you don't have to worry about all your packages failing, because you can super easily remove it.

### Pros

- Easy to use with simple commands.
- Explains what it's doing at every step.
- Simple compared to other package managers, and easy to repair if something goes wrong.
- Clones directly from the package's [git](https://git-scm.com/) repository, giving you the latest version with no manual intervention by a maintainer.
- Easy to submit to, with no notability requirements.
- Installs **everything** locally, not interfering with any system components and not requiring root privileges.

### Cons

- Doesn't install & manage dependencies.
- Slower than other package managers due to needing to compile source code.

## Installation

### Dependencies

- [Git](https://git-scm.com/)
- [Go](https://go.dev/) 1.17+
- [Make](https://en.wikipedia.org/wiki/Make_(software)#Derivatives)

### Install script

#### Dependencies for install script

- [cURL](https://curl.se/)
- [bash](https://gnu.org/software/bash)/[zsh](https://www.zsh.org/)

To install using the install script, run:

```bash
bash <(curl -s https://raw.githubusercontent.com/talwat/indiepkg/testing/scripts/install.sh) # To use bash
zsh <(curl -s https://raw.githubusercontent.com/talwat/indiepkg/testing/scripts/install.sh) # Or if you only have zsh installed
```

### Manual install

If the install script doesn't work for you, you can manually install the following:

```bash
git clone -b testing https://github.com/talwat/indiepkg.git
cd indiepkg
make
make install
```

## Uninstallation

Currently, there is no uninstallation script, but you can super easily uninstall IndiePKG by running these commands.

```bash
rm -rf ~/.indiepkg
rm ~/.local/bin/indiepkg
```

## Making & submitting packages

If you want to make a package, you should look at [PACKAGES.md](MAKE_A_PKG.md)

## Basic usage

You can run `indiepkg install <packages>` to install a package.

If you want to uninstall a package, you can run `indiepkg uninstall <packages>`.

`indiepkg upgrade [packages]` will pull the latest changes and recompile packages. If you don't specify any packages, it will upgrade all packages.

`indiepkg update [packages]` will update the information for your installed packages. This command doesn't need to be ran frequently at all, but it's best to run it every now and then.

## Supported operating systems

IndiePKG officially works on **GNU/Linux** and **macOS**.

It will work on any other unix-like operating system but many packages won't compile properly. If you are interested in supporting a new operating system, please [open an issue](https://github.com/talwat/indiepkg/issues/new/choose) and we can discuss making a new third party repository.

IndiePKG is **completely** broken on Windows. IndiePKG is overall not meant to run on Windows.

### List of operating systems

| Darwin                                                                                  | Linux                                                                                 | Other                                                                                   |
| --------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- |
| ![Darwin arm64](https://img.shields.io/badge/darwin%20arm64-full%20support-brightgreen) | ![Linux amd64](https://img.shields.io/badge/linux%20amd64-full%20support-brightgreen) | ![BSDs](https://img.shields.io/badge/BSDs-partial%20support-yellow)            |
| ![Darwin amd64](https://img.shields.io/badge/darwin%20amd64-great%20support-green)      | ![Linux arm64](https://img.shields.io/badge/linux%20arm64-great%20support-green)      | ![Other unix-like](https://img.shields.io/badge/other%20unix--like-partial%20support-yellow) |
|                                                                                         | ![Linux other](https://img.shields.io/badge/linux%20other-good%20support-yellowgreen) | ![Windows](https://img.shields.io/badge/windows-borked-red) |

- **Full Support** - All packages & binary packages in the official repo are guaranteed to work. *This excludes linux-only*
- **Great Support** - All source packages and most binary packages are supposed to work.
- **Partial Support** - Some packages will work, but not all.
- **Borked** - IndiePKG itself will not function properly.

## Branches

IndiePKG has 3 branches to install from, **testing** *(recommended)*, **stable**, and **main**.

The **testing** branch is a good balance of new and stable, and is updated whenever all Github actions pass.

The **stable** branch is only updated whenever a release is made, and is not recommended for use *yet* because it will sometimes not work due to changes in the package files.

The **main** branch is the default git branch, and contains the **newest** code. This is not recommended for use because it has a higher chance of being broken.

## Libraries used

A huge thank you to the following libraries:

- [go-toml](https://github.com/pelletier/go-toml.git) - For parsing the configuration file.
- [progressbar](https://github.com/schollz/progressbar.git) - For making progressbar's during downloads.

## FAQ

### So it's like the AUR?

Kind of. For one packages are checked, and there is also an official package manager unlike the AUR which is commonly paired with an AUR helper.

It also obviously works on other operating systems besides just Arch Linux.


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Ftalwat%2Findiepkg.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Ftalwat%2Findiepkg?ref=badge_large)