# Making & submitting packages

## Table of contents

- [Making & submitting packages](#making--submitting-packages)
  - [Table of contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Constructing the package file](#constructing-the-package-file)
    - [Basic fields](#basic-fields)
      - [Name](#name)
      - [Author](#author)
      - [Description](#description)
      - [URL](#url)
      - [License](#license)
      - [Branch (optional)](#branch-optional)
    - [Paths](#paths)
      - [Bin](#bin)
        - [In Source (optional)](#in-source-optional)
        - [Installed](#installed)
    - [Commands (optional in some cases)](#commands-optional-in-some-cases)
      - [Install](#install)
      - [Update](#update)
      - [Uninstall](#uninstall)

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

### Basic fields

These are super simple pieces of information to fill out.

#### Name

This is the name of the package. You are **not** allowed to use spaces in the name.

#### Author

This is your name, it can be a username or a real name.

#### Description

This is a super short summary of the purpose or function of the package.

#### URL

This is the **git** url. It must end with `.git`.

Your repository **must** be public. This is the URL you would use to `git clone`.

#### License

This is the license. It can be any free license.

#### Branch (optional)

You can specify which branch to clone. If this isn't specified it will simply clone the default branch *Usually master or main*.

### Paths

There are a few options which specify certain directories/file paths for various things.

#### Bin

This describes the location of **binary** files.

##### In Source (optional)

This is where the binary is directly **after** the package is compiled. It is a relative path.

For example, if the install instructions say to run `make` and then a binary is produced in `bin` directory, you would add `bin/<file>`. *Replace \<file> with the name of the generated file*

##### Installed

This is the **name** of the final installed binary file, **not** the path.

For example, if your generated binary file was called `my_file` then you would add `my_file`.

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
