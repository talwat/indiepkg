<!-- markdownlint-disable MD013 -->
# How to make a package

This will guide you how to make, test, and submit an IndiePKG package.

## Table of contents

- [How to make a package](#how-to-make-a-package)
  - [Table of contents](#table-of-contents)
  - [Package requirements](#package-requirements)
  - [Constructing your package.json file](#constructing-your-packagejson-file)
  - [Testing your package](#testing-your-package)
    - [Installing from a file (recommended)](#installing-from-a-file-recommended)
    - [Installing from a 3rd party repository](#installing-from-a-3rd-party-repository)
  - [Making a pull request](#making-a-pull-request)
  - [Badges](#badges)
    - [Flat](#flat)
      - [Markdown (flat)](#markdown-flat)
    - [Flat square](#flat-square)
      - [Markdown (flat square)](#markdown-flat-square)
    - [Plastic](#plastic)
      - [Markdown (plastic)](#markdown-plastic)
    - [For the badge](#for-the-badge)
      - [Markdown (for the badge)](#markdown-for-the-badge)

## Package requirements

Before you start, you should know there are a few rules to making a package.

- Your package should **not** have a GUI, only command line apps are functional.
- Your package should have a git repository.

## Constructing your package.json file

First you need to make a package.json file.

If your package is on Github, use `indiepkg github-gen <username> <repo>` to generate a package.json file. (This still should be manually checked after it's generated)

If your package isn't on Github, copy a fitting template from `samples`.

Go to [PACKAGES.md](PACKAGES.md) and look at what each field does to construct your package.json files. Package templates for specific languages are available in [samples/templates](samples/templates).

## Testing your package

### Installing from a file (recommended)

Simply run `indiepkg install <the name of your package file>` to install your package.

If it was successful and your package installed correctly, you can move on to the final step.

### Installing from a 3rd party repository

This is a little more complicated, but not too complicated.

First, [fork](https://github.com/talwat/IndiePKG/fork) IndiePKG on github.

After your done forking, add your package.json file to the `packages` directory by either cloning the fork locally and committing, or using the Github Web UI.

If your package is only functional on Linux, add it to `packages/linux-only`.

Once that's done, add your repository as a [3rd party repo](docs/REPOS.md).

```bash
indiepkg repo add https://github.com/<YOUR USERNAME HERE>/indiepkg/main/packages/
```

Or if your package is only available on Linux, run this command:

```bash
indiepkg repo add https://github.com/<YOUR USERNAME HERE>/indiepkg/main/packages/linux-only/
```

Finally, install your package by running:

```bash
indiepkg install <YOUR PACKAGE NAME>
```

If everything worked fine, your ready to move on to the final step.

## Making a pull request

Finally, go to your Github repo and make a pull request.

I will then test the package myself, and if everything works, I will merge it into the main repository.

Make sure when you make your PR you merge to the **main** branch.

## Badges

If you have successfully submitted your program to IndiePKG, if you wish, you can add a badge to your repository.

### Flat

[![indiepkg-badge](https://img.shields.io/badge/get_on-indiepkg-blue?style=flat&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADwAAAA8CAYAAAA6/NlyAAAAxklEQVRoge3Zuw3CMBRGYYzoGYEhIrEGSLTswCywQ0pGyRKMQEthRvCPFF7H56uvnBy5uUrKIlRrrenst5RSSmtm+YkX+SUG0xlMZzCdwXSrdxyabDyvmHPL6+6GDaYzmM5gOoPpyj98q5pTdzdsMJ3BdAbTGUwXf9M6nm7xoeNlE82lS951GKK5wzQ1Z7q7YYPpDKYzmM5gunjTWu938aH1nG1Q6U/Gx30bP7uluxs2mM5gOoPpDKbz7yGdwXQG0xlMZzDdE3J4HtdWbCB+AAAAAElFTkSuQmCC)](https://github.com/talwat/indiepkg)

#### Markdown (flat)

```markdown
[![indiepkg-badge](https://img.shields.io/badge/get_on-indiepkg-blue?style=flat&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADwAAAA8CAYAAAA6/NlyAAAAxklEQVRoge3Zuw3CMBRGYYzoGYEhIrEGSLTswCywQ0pGyRKMQEthRvCPFF7H56uvnBy5uUrKIlRrrenst5RSSmtm+YkX+SUG0xlMZzCdwXSrdxyabDyvmHPL6+6GDaYzmM5gOoPpyj98q5pTdzdsMJ3BdAbTGUwXf9M6nm7xoeNlE82lS951GKK5wzQ1Z7q7YYPpDKYzmM5gunjTWu938aH1nG1Q6U/Gx30bP7uluxs2mM5gOoPpDKbz7yGdwXQG0xlMZzDdE3J4HtdWbCB+AAAAAElFTkSuQmCC)](https://github.com/talwat/indiepkg)
```

### Flat square

[![indiepkg-badge](https://img.shields.io/badge/get_on-indiepkg-blue?style=flat-square&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADwAAAA8CAYAAAA6/NlyAAAAxklEQVRoge3Zuw3CMBRGYYzoGYEhIrEGSLTswCywQ0pGyRKMQEthRvCPFF7H56uvnBy5uUrKIlRrrenst5RSSmtm+YkX+SUG0xlMZzCdwXSrdxyabDyvmHPL6+6GDaYzmM5gOoPpyj98q5pTdzdsMJ3BdAbTGUwXf9M6nm7xoeNlE82lS951GKK5wzQ1Z7q7YYPpDKYzmM5gunjTWu938aH1nG1Q6U/Gx30bP7uluxs2mM5gOoPpDKbz7yGdwXQG0xlMZzDdE3J4HtdWbCB+AAAAAElFTkSuQmCC)](https://github.com/talwat/indiepkg)

#### Markdown (flat square)

```markdown
[![indiepkg-badge](https://img.shields.io/badge/get_on-indiepkg-blue?style=flat-square&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADwAAAA8CAYAAAA6/NlyAAAAxklEQVRoge3Zuw3CMBRGYYzoGYEhIrEGSLTswCywQ0pGyRKMQEthRvCPFF7H56uvnBy5uUrKIlRrrenst5RSSmtm+YkX+SUG0xlMZzCdwXSrdxyabDyvmHPL6+6GDaYzmM5gOoPpyj98q5pTdzdsMJ3BdAbTGUwXf9M6nm7xoeNlE82lS951GKK5wzQ1Z7q7YYPpDKYzmM5gunjTWu938aH1nG1Q6U/Gx30bP7uluxs2mM5gOoPpDKbz7yGdwXQG0xlMZzDdE3J4HtdWbCB+AAAAAElFTkSuQmCC)](https://github.com/talwat/indiepkg)
```

### Plastic

[![indiepkg-badge](https://img.shields.io/badge/get_on-indiepkg-blue?style=plastic&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADwAAAA8CAYAAAA6/NlyAAAAxklEQVRoge3Zuw3CMBRGYYzoGYEhIrEGSLTswCywQ0pGyRKMQEthRvCPFF7H56uvnBy5uUrKIlRrrenst5RSSmtm+YkX+SUG0xlMZzCdwXSrdxyabDyvmHPL6+6GDaYzmM5gOoPpyj98q5pTdzdsMJ3BdAbTGUwXf9M6nm7xoeNlE82lS951GKK5wzQ1Z7q7YYPpDKYzmM5gunjTWu938aH1nG1Q6U/Gx30bP7uluxs2mM5gOoPpDKbz7yGdwXQG0xlMZzDdE3J4HtdWbCB+AAAAAElFTkSuQmCC)](https://github.com/talwat/indiepkg)

#### Markdown (plastic)

```markdown
[![indiepkg-badge](https://img.shields.io/badge/get_on-indiepkg-blue?style=plastic&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADwAAAA8CAYAAAA6/NlyAAAAxklEQVRoge3Zuw3CMBRGYYzoGYEhIrEGSLTswCywQ0pGyRKMQEthRvCPFF7H56uvnBy5uUrKIlRrrenst5RSSmtm+YkX+SUG0xlMZzCdwXSrdxyabDyvmHPL6+6GDaYzmM5gOoPpyj98q5pTdzdsMJ3BdAbTGUwXf9M6nm7xoeNlE82lS951GKK5wzQ1Z7q7YYPpDKYzmM5gunjTWu938aH1nG1Q6U/Gx30bP7uluxs2mM5gOoPpDKbz7yGdwXQG0xlMZzDdE3J4HtdWbCB+AAAAAElFTkSuQmCC)](https://github.com/talwat/indiepkg)
```

### For the badge

[![indiepkg-badge](https://img.shields.io/badge/get_on-indiepkg-blue?style=for-the-badge&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADwAAAA8CAYAAAA6/NlyAAAAxklEQVRoge3Zuw3CMBRGYYzoGYEhIrEGSLTswCywQ0pGyRKMQEthRvCPFF7H56uvnBy5uUrKIlRrrenst5RSSmtm+YkX+SUG0xlMZzCdwXSrdxyabDyvmHPL6+6GDaYzmM5gOoPpyj98q5pTdzdsMJ3BdAbTGUwXf9M6nm7xoeNlE82lS951GKK5wzQ1Z7q7YYPpDKYzmM5gunjTWu938aH1nG1Q6U/Gx30bP7uluxs2mM5gOoPpDKbz7yGdwXQG0xlMZzDdE3J4HtdWbCB+AAAAAElFTkSuQmCC)](https://github.com/talwat/indiepkg)

#### Markdown (for the badge)

```markdown
[![indiepkg-badge](https://img.shields.io/badge/get_on-indiepkg-blue?style=for-the-badge&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADwAAAA8CAYAAAA6/NlyAAAAxklEQVRoge3Zuw3CMBRGYYzoGYEhIrEGSLTswCywQ0pGyRKMQEthRvCPFF7H56uvnBy5uUrKIlRrrenst5RSSmtm+YkX+SUG0xlMZzCdwXSrdxyabDyvmHPL6+6GDaYzmM5gOoPpyj98q5pTdzdsMJ3BdAbTGUwXf9M6nm7xoeNlE82lS951GKK5wzQ1Z7q7YYPpDKYzmM5gunjTWu938aH1nG1Q6U/Gx30bP7uluxs2mM5gOoPpDKbz7yGdwXQG0xlMZzDdE3J4HtdWbCB+AAAAAElFTkSuQmCC)](https://github.com/talwat/indiepkg)
```
