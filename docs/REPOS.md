<!-- markdownlint-disable MD013 -->

# 3rd party repositories

## Table of contents

- [3rd party repositories](#3rd-party-repositories)
  - [Table of contents](#table-of-contents)
  - [Purpose](#purpose)
  - [Structure](#structure)
  - [Rules](#rules)

## Purpose

Third party repositories add more package locations than IndiePKG provides by default, and are super useful if you have a program whose package.json file changes frequently.

## Structure

3rd party repositories are added & handled in the form of simple URL's. This URL should lead to a sort of file tree.

IndiePKG will take your URL, append the name of the package file it's looking for, and download that file.

So if you added the URL: `https://myrepo.com/` and tried installing `my-cool-pkg` IndiePKG would end up downloading a file from a URL that looks something like this: `https://myrepo.com/my-cool-pkg.json`.

This means that you can also specify subdirectories as your repository root.

For example, if you added: `https://myrepo.com/indiepkg-stuff/` then IndiePKG would now look there, and the URL for a package called `my-cool-pkg` would now look like this: `https://myrepo.com/indiepkg-stuff/my-cool-pkg.json`.

This additionally removes the need of having a git repository, since you can use anything. **NOTE**: Querying is only supported on github, so be warned of that.

## Rules

All your packages that you want to be able to install from a 3rd party repository must end in `.json`.

The URL's should be **raw**, meaning there isn't any fancy HTML or CSS like github's view for example.

If your using github, you should put `https://raw.githubusercontent.com/<repo>/<username>/main/` instead of `https://github.com/<repo>/<username>/blob/main/`
