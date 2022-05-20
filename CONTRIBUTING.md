<!-- markdownlint-disable MD013 -->

# Contributing

## Table of contents

- [Contributing](#contributing)
  - [Table of contents](#table-of-contents)
  - [Package guidelines](#package-guidelines)
  - [Standard PR guidelines](#standard-pr-guidelines)

## Package guidelines

In order to submit your package, it **must** follow the following guidelines:

- It must be licensed under a free and open source license.
- It must either have an accessible tarball/binary or a git repository. *Git repository highly recommended*.
- It must **not** require `sudo` privileges and it shouldn't touch any directories outside the users home directory. *With the exception of `/tmp`*.
- It must be a CLI or TUI application and **not** a GUI application.
- It should **not** call another package manager. *With the exception of package managers such as `cargo` serving as build systems*.
- It also should obviously not be malicious, contain inappropriate content, etc...
- Your package should also have been **tested** beforehand.

## Standard PR guidelines

If you would like to contribute code to IndiePKG, please make sure that your code has been **tested**.

Additionally, it should pass all [golangci-lint](https://golangci-lint.run/) tests.

When making your PR, request to merge to the **main** *(development)* branch.
