package main

// This file is mainly for long strings of text, like the default config file & help message.

const helpMsg = `Usage: indiepkg [options...] <command> [args...]

Commands:
  help                            Show this help message.
  install <packages...>           Installs packages.
  uninstall <packages...>         Removes packages.
  update [packages...]            Re-downloads the a package's info & install instructions. If no packages are specified, all packages are updated.
  upgrade [packages...]           Pulls git repository & recompile's a package. If no package is specified, all packages are upgraded.
  info <package>                  Displays information about a specific package.
  remove-data <packages...>       Removes package data from .indiepkg. Use this only if a package installation has failed and the uninstall command won't work.
  re-clone                        Removes and re-clones IndiePKG source code in the src directory. Useful if you just changed the branch in the config file, or git is throwing errors when updating.
  list                            Lists all installed packages.
  list-all                        Lists all packages in all github repositories.
  version                         Shows version.
  init                            Re-generates all the default config files needed for indiepkg to function properly. This is ran automatically.
  search <query>                  Searches for packages. NOTE: Only searches repo's that are on Github.
  indiepkg-update                 Updated source code & re-compiles IndiePKG.
  setup                           Sets up everything for a functional installation of IndiePKG.
  env-add                         Adds IndiePKG to several environment variables. This is useful if packages installed with IndiePKG are not being found.
  repo                            Manages the package sources file.
    repo add <url>                Adds a repository to the package sources file.
    repo remove <url>             Removes a repository to the package sources file.
    repo list                     Lists all repositories in the package sources file.

Developer & Debugging Commands:
  fetch                           Fetches OS & IndiePKG info.
  raw-version                     Shows version without any formatting (For use in scripts).
  github-gen <username> <repo>    Generates a package info file for a github repo.
  help2man                        Generates a manpage in the current working directory. NOTE: This only works if your current working directory is in the IndiePKG source directory.

Options:
  -p, --purge                     Removes a package's configuration files as well as the package itself.
  -d, --debug                     Displays variable & debugging information.
  -y, --assumeyes                 Assumes yes to all prompts. (Use with caution!)
  -f, --force                     Bypasses all checks before preforming an action. Use will almost certainly lead to an error.
  -r, --ignoreroot                Continues even if user is root.
  -n, --nodeps                    Continues even if there are unmet dependencies.

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

# Configures updating and auto-updating
[updating]

# Default: testing
# The branch to pull from. https://github.com/talwat/indiepkg#branches for more information about the branches.
branch = "testing"

# Default: true
auto_update = true

# Configures progressbar, purely visual.
[progressbar]

saucer = "[cyan]=[reset]"
saucer_head = "[cyan]>[reset]"
alt_saucer_head = "[cyan]>[reset]"
saucer_padding = " "
bar_start = "("
bar_end = ")"

# Configures github authentication, this is used for querying packages.
[github]

# Github username
username = ""

# Github access token, this should not expire and have no permissions granted for security purposes.
# Information about getting a token is available here. https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token
token = ""
`

const defaultSources = `# Please only add sources you trust.
# This file contains the links to the pkg.json files. If you mess up, you can simply run 'indiepkg init' to reset it.
# You can also edit this file with 'indiepkg repo'.

https://raw.githubusercontent.com/talwat/indiepkg/main/packages/
# https://raw.githubusercontent.com/talwat/indiepkg/main/packages/bin/
# https://raw.githubusercontent.com/talwat/indiepkg/main/packages/linux-only/`
