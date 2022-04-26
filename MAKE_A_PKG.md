<!-- markdownlint-disable MD013 -->

# How to make a package

This will guide you how to make, test, and submit an IndiePKG package.

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

Now, we will need to make a third party repository to test the package.

This is super simple. First, download and install IndiePKG.

Next, [fork](https://github.com/talwat/IndiePKG/fork) IndiePKG on github.

After your done forking, add your package.json file to the `packages` directory by either cloning the fork locally and committing, or using the Github Web UI.

If your package is only functional on Linux, add it to `packages/linux-only`.

Once that's done, add your repository as a [3rd party repo](docs/REPOS.md).

```bash
indiepkg repo add https://github.com/<YOUR USERNAME HERE>/indiepkg/main/packages/
```

Or if your package is only available on Linux, run this command:

```bash
indiepkg repo add indiepkg repo add https://github.com/<YOUR USERNAME HERE>/indiepkg/main/packages/linux-only/
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
