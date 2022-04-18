package main

// This file is mainly for long strings of text, like the default config file & help message.

const helpMsg = `Usage: indiepkg [<options>...] <command>

Commands:
  help                         Show this help message.
  install <packages...>        Installs packages.
  uninstall <packages...>      Removes packages.
  update [packages...]         Re-downloads the a package's info & install instructions. If no packages are specified, all packages are updated.
  upgrade [packages...]        Pulls git repository & recompile's a package. If no package is specified, all packages are upgraded.
  info <package>               Displays information about a specific package.
  remove-data <packages...>    Removes package data from .indiepkg. Use this only if a package installation has failed and the uninstall command won't work.
  sync                         Sync package info & package source.
  list                         Lists all installed packages.
  version                      Shows version.
  init                         Re-generates all the default config files needed for indiepkg to function properly. This is ran automatically.
  search <query>               Searches for packages. NOTE: Only searches repo's that are on Github.
  indiepkg-update              Updated source code & re-compiles IndiePKG.
  repo                         Manages the package sources file.
    repo add <url>             Adds a repository to the package sources file.
    repo remove <url>          Removes a repository to the package sources file.

Options:
  -p, --purge                  Removes a package's configuration files as well as the package itself.
  -d, --debug                  Displays variable & debugging information.
  -y, --assumeyes              Assumes yes to all prompts. (Use with caution!)
  -f, --force                  Bypasses all checks before preforming an action. Use will almost certainly lead to an error.

Examples:
  indiepkg install my-pkg
  indiepkg uninstall other-pkg
  indiepkg upgrade third-pkg
`

const defaultConf = `# IndiePKG config file

# Locations for files
# The home directory is prepended automatically for each value, do not include it.
[paths]

# Default: .local
prefix = ".local"
`

const defaultSources = `# Please only add sources you trust.
# This file contains the links to the pkg.json files. If you mess up, you can simply run 'indiepkg init' to reset it.
# You can also edit this file with 'indiepkg repo'.

https://raw.githubusercontent.com/talwat/indiepkg/main/packages/
# https://raw.githubusercontent.com/talwat/indiepkg/main/packages/bin/
# https://raw.githubusercontent.com/talwat/indiepkg/main/packages/linux-only/`
