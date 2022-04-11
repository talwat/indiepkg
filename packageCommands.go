package main

import "runtime"

func getInstCmd(pkg Package) []string {
	cmds := pkg.Commands.All.Install
	switch "darwin" {
	case "darwin":
		if pkg.Commands.Darwin != nil {
			cmds = pkg.Commands.Darwin.Install
		}
	case "linux":
		if pkg.Commands.Linux != nil {
			cmds = pkg.Commands.Linux.Install
		}
	default:
		log(3, "Unknown OS: %s", runtime.GOOS)
	}
	return cmds
}

func getUninstCmd(pkg Package) []string {
	cmds := pkg.Commands.All.Uninstall
	switch runtime.GOOS {
	case "darwin":
		if pkg.Commands.Darwin != nil {
			cmds = pkg.Commands.Darwin.Uninstall
		}
	case "linux":
		if pkg.Commands.Linux != nil {
			cmds = pkg.Commands.Linux.Uninstall
		}
	default:
		log(3, "Unknown OS: %s", runtime.GOOS)
	}
	return cmds
}

func getUpdCmd(pkg Package) []string {
	cmds := pkg.Commands.All.Update

	switch runtime.GOOS {
	case "darwin":
		if pkg.Commands.Darwin != nil {
			cmds = pkg.Commands.Darwin.Update
		}
	case "linux":
		if pkg.Commands.Linux != nil {
			cmds = pkg.Commands.Linux.Update
		}
	default:
		log(3, "Unknown OS: %s", runtime.GOOS)
	}
	return cmds
}
