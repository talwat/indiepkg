package main

// This file is mainly for long strings of text, like the default config file & help message.

const helpMsg = `Usage: indiepkg [<options>...] <command>

Commands:
  help                       Show this help message.
  install <packages...>      Installs packages.
  uninstall <packages...>    Removes packages.
  update [packages...]       Re-downloads the a package's info & install instructions. If no packages are specified, all packages are updated.
  upgrade [packages...]      Pulls git repository & recompile's a package. If no package is specified, all packages are upgraded.
  info <package>             Displays information about a specific package.
  remove-data <packages...>  Removes package data from .indiepkg. Use this only if a package installation has failed and the uninstall command won't work.
  sync                       Sync package info & package source.
  version                    Show version.

Options:
  -p, --purge                Removes a package's configuration files as well as the package itself.
  -d, --debug                Displays variable & debugging information.
  -y, --assumeyes            Assumes yes to all prompts. (Use with caution!)

Examples:
  indiepkg install my-pkg
  indiepkg uninstall other-pkg
  indiepkg upgrade third-pkg
`

const defaultConf = `{}`

const defaultSources = `# This file contains the links to the pkg.json files. If you mess up, you can simply run 'indiepkg init' to reset it.'

https://raw.githubusercontent.com/talwat/indiepkg/main/packages/`
