<!-- markdownlint-disable MD013 -->

# The package format

## Table of contents

- [The package format](#the-package-format)
  - [Table of contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Constructing the package file](#constructing-the-package-file)
    - [Environment variables](#environment-variables)
    - [Basic fields](#basic-fields)
      - [Name](#name)
      - [Author](#author)
      - [Description](#description)
      - [URL (don't use with Download URL)](#url-dont-use-with-download-url)
      - [Download URL (don't use with URL)](#download-url-dont-use-with-url)
      - [Info URL (optional)](#info-url-optional)
      - [License (optional, but heavily encouraged)](#license-optional-but-heavily-encouraged)
      - [Language (optional, but heavily encouraged)](#language-optional-but-heavily-encouraged)
      - [Branch (optional)](#branch-optional)
    - [Paths](#paths)
      - [Bin](#bin)
        - [In Source (optional)](#in-source-optional)
        - [Installed (optional)](#installed-optional)
      - [Manpages](#manpages)
    - [Commands (optional in some cases)](#commands-optional-in-some-cases)
      - [Install](#install)
      - [Update](#update)
      - [Uninstall](#uninstall)
    - [Others](#others)
    - [Deps](#deps)
      - [Config paths](#config-paths)
      - [Notes](#notes)

## Introduction

IndiePKG's packages are contained in a single `.json` file called the package file.

This file contains various information including metadata, git link, install commands, and a few other things.

You don't need to update the file frequently because IndiePKG will simply get the **git** url from the file and not a hard-coded link to an archive.

## Constructing the package file

First, you can copy the minimal example from `samples/pkg_minimal.json`. Or copy paste it from here:

```json
{
    "name": "my-package-name",
    "author": "my username or name",
    "description": "this is my awesome package description",
    "url": "https://github.com/username/repo.git",
    "license": "license name",
    "language": "the language my package is written in",
    "bin": {
        "in_source": [
            "bin/my_file"
        ],
        "installed": [
            "my_file"
        ]
    },
    "deps": {
        "all": [
            "my_dep"
        ]
    },
    "config_paths": [
        ".config/my_file"
    ]
}
```

Also, please refrain from using any unicode/special symbols in any field as it can confuse IndiePKG.

### Environment variables

Environment variables are structured like this `:(VARIABLE_NAME):`.

Right now, IndiePKG supports the following environment variables:

- `:(PREFIX):` - By default this is set to `~/.local`. You can use this with Makefile's. For example in the command section you could put `make install PREFIX=:(PREFIX):`.
- `:(BIN)` - By default this is set to `~/.local/bin`. This is usually not needed.
- `:(HOME):` - This is the users home directory.

These will work in any field, although they don't have much use outside `commands`.

### Basic fields

These are super simple pieces of information to fill out.

#### Name

This is the name of the package. You are **not** allowed to use spaces in the name.

Your package.json file should also be named the same as the `name` argument in the package file.

#### Author

This is your name, it can be a username or a real name.

#### Description

This is a super short summary of the purpose or function of the package.

#### URL (don't use with [Download URL](#download-url-dont-use-with-url))

This is the **git** url. It must end with `.git`.

Your repository **must** be public. This is the URL you would use to `git clone`.

#### Download URL (don't use with [URL](#url-dont-use-with-download-url))

This is the **download** url to the file/archive. Do not use this with the `URL` property.

The structure goes something like this:

```json
...
"download": {
    "darwin": {
        "arm64": "link",
        "amd64": "second link"
    },
    "linux": {
        "amd64": "other link",
        "arm64": "third link"
    }
}
```

As you can specify which architecture should download which file, however if your archive works on all operating systems and architectures you can also do:

```json
...
"download": {
    "all": "link"
}
```

Or if it works on a specific operating system but all architectures:

```json
...
"download": {
    "linux": {
        "all": "link"
    },
    "darwin": {
        "all": "other link"
    }
}
```

**Note**: This will save the downloaded file as `<package name>.indiepkg.downloaded`.

#### Info URL (optional)

This is the URL to the package's info file. Please only put this if you are **not** distributing the package in a repository and instead using a direct URL/file.

#### License (optional, but heavily encouraged)

This is the license. It can be any license. Make sure to use the SPDX ID.

#### Language (optional, but heavily encouraged)

This is the primary programming language used to make the program.

#### Branch (optional)

You can specify which branch to clone. If this isn't specified it will simply clone the default branch *Usually `master` or `main`*.

### Paths

There are a few options which specify certain directories/file paths for various things.

#### Bin

This describes the location of **binary** files.

##### In Source (optional)

This is where the binary is directly **after** the package is compiled. It is a relative path.

For example, if the install instructions say to run `make` and then a binary is produced in `bin` directory, you would add `bin/<file>`. *Replace \<file> with the name of the generated file*

This binary will be **moved** to the proper `bin` directory.

##### Installed (optional)

This is the **name** of the final installed binary file, **not** the path.

For example, if your generated binary file was called `my_file` then you would add `my_file`.

This is used to uninstall the package easily.

#### Manpages

The location of manpages inside the source directory. You don't need to specify **where** to install them because IndiePKG can figure that out automatically.

Make sure the manpage file extension ends with `.x`. `x` can be any number from 1-9*

### Commands (optional in some cases)

These are the commands which are ran when the package is installed in the root of the source code directory.

If your package is a super simple bash script, you don't need to specify any commands at all because you can just specify the [binary files](#bin).

You can also specify commands to run on specific operating systems.

- `all` This refers to all *unix-like* operating systems.
- `linux` These commands will be run **instead** of the `all` commands if the user is running a **Linux** system.
- `darwin` These commands will be run **instead** of the `all` commands if the user is running a **Darwin** *(macOS)* system.

For example:

```json
{
    ...
    "commands": {
        "all": {
            "install": [
                "make",
            ],
            "update": [
                "make"
            ]
        },
        "darwin": {
            "install": [
                "gmake",
            ],
            "update": [
                "gmake"
            ]
        }
    }
}
```

If you want the program to continue even if a command throws an error, prefix the command with `!(FORCE)!`.

Eg.

```json
[
    "make",
    "!(FORCE)! mkdir -p :(PREFIX):share/foo",
]
```

]

#### Install

These are the commands that are run to **compile** the package. For example, it would be:

```json
[
    "./configure",
    "make"
]
```

if you had a typical C program.

If you have specific install instructions that aren't **just** moving a binary, you should also add them here.

#### Update

These are the commands to **re-compile** *(and sometimes also to **re-install** the package)*.

In some cases you can even just copy paste the install commands.

Going back to our C example:

```json
[
    "make"
]
```

#### Uninstall

Usually, this isn't required, but if your program needs a specific set of instructions to be **fully** uninstalled then you would specify them here.

IndiePKG will automatically delete the binary files specified [here](#installed).

### Others

### Deps

This is where you put the dependencies of your program. As of now you can't put paths or version numbers, only commands.

You should also put your **build** dependencies here.

You can also as usual specify the operating system. Keep in mind that OS specific dependencies will be **appended** to the list of all dependencies.

Again, to our C example:

```json
{
    ...
    "deps": {
        "all": [
            "gcc",
            "g++",
            "autoconf"
        ],
        "linux": [
            "make"
        ],
        "darwin": [
            "gmake"
        ]
    }
}
```

#### Config paths

This tells IndiePKG where your configuration files are located so it can remove them if needed.

**Note**: The home directory is *automatically* pre-pended, so you shouldn't put an absolute path.xs

#### Notes

This is where you can put any additional information about using the program with IndiePKG.

For example, if there's a command you need to run or a symlink you need to create, you can put that information here.
