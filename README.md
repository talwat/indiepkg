# IndiePKG

A package manager written in go for small CLI programs. It is available on GNU/Linux and macOS.

## Notice

IndiePKG is **NOT** ready for use yet. It's still extremely early software.

However, if you would like to submit issues or PR's, you are **more** than welcome to.

## What is IndiePKG?

IndiePKG is mainly for small simple CLI and TUI programs. Most of them are just for fun, such as **cmatrix**, while others have a bit more utility such as **btop**.

IndiePKG uses **git** to install packages, and everything is compiled **from source**. This means that while there aren't any versions, it does mean that you get the absolute latest software.

It's also much simpler than your standard package manager, and if a package installation goes wrong you don't have to worry about all your packages failing, because you can super easily remove it.

## Installation

While IndiePKG doesn't have an install script yet, you can still try it out by cloning the repository and building it from source.
You will need:

- Git
- Go 1.18 *Older versions might work too*
- Make

```bash
git clone https://github.com/talwat/indiepkg.git
cd indiepkg
make
make install
```

## Basic usage

You can run `indiepkg install <packages>` to install a package.

If you want to uninstall a package, you can run `indiepkg uninstall <packages>`.

`indiepkg upgrade [packages]` will pull the latest changes and recompile packages. If you don't specify any packages, it will upgrade all packages.

`indiepkg update [packages]` will update the information for your installed packages. This command doesn't need to be ran frequently at all, but it's best to run it every now and then.
