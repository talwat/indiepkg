// TODO: Clean this code up

package main

import "runtime"

func getInstCmd(pkg Package) []string {
	log(1, "Getting install commands...")

	var cmds []string

	if pkg.Commands == nil {
		return []string{}
	}

	if pkg.Commands.All != nil && pkg.Commands.All.Install != nil {
		cmds = pkg.Commands.All.Install
	}

	switch runtime.GOOS {
	case "darwin":
		if pkg.Commands.Darwin != nil && pkg.Commands.Darwin.Install != nil {
			debugLog("Getting install commands for Darwin. Package: %s", bolden(pkg.Name))
			cmds = pkg.Commands.Darwin.Install
		}
	case "linux":
		if pkg.Commands.Linux != nil && pkg.Commands.Linux.Install != nil {
			debugLog("Getting install commands for Linux. Package: %s", bolden(pkg.Name))
			cmds = pkg.Commands.Linux.Install
		}
	default:
		log(3, "Unknown OS: %s", runtime.GOOS)
	}
	return cmds
}

func getUninstCmd(pkg Package) []string {
	log(1, "Getting uninstall commands...")

	var cmds []string

	if pkg.Commands == nil {
		return []string{}
	}

	if pkg.Commands.All != nil && pkg.Commands.All.Uninstall != nil {
		cmds = pkg.Commands.All.Uninstall
	}

	switch runtime.GOOS {
	case "darwin":
		if pkg.Commands.Darwin != nil && pkg.Commands.Darwin.Uninstall != nil {
			debugLog("Getting uninstall commands for Darwin. Package: %s", bolden(pkg.Name))
			cmds = pkg.Commands.Darwin.Uninstall
		}
	case "linux":
		if pkg.Commands.Linux != nil && pkg.Commands.Linux.Uninstall != nil {
			debugLog("Getting uninstall commands for Linux. Package: %s", bolden(pkg.Name))
			cmds = pkg.Commands.Linux.Uninstall
		}
	default:
		log(3, "Unknown OS: %s", runtime.GOOS)
	}
	return cmds
}

func getUpdCmd(pkg Package) []string {
	log(1, "Getting upgrade commands...")

	var cmds []string

	if pkg.Commands == nil {
		return []string{}
	}

	if pkg.Commands.All != nil && pkg.Commands.All.Update != nil {
		cmds = pkg.Commands.All.Update
	}

	switch runtime.GOOS {
	case "darwin":
		if pkg.Commands.Darwin != nil && pkg.Commands.Darwin.Update != nil {
			debugLog("Getting update commands for Darwin. Package: %s", bolden(pkg.Name))
			cmds = pkg.Commands.Darwin.Update
		}
	case "linux":
		if pkg.Commands.Linux != nil && pkg.Commands.Linux.Update != nil {
			debugLog("Getting update commands for Linux. Package: %s", bolden(pkg.Name))
			cmds = pkg.Commands.Linux.Update
		}
	default:
		log(3, "Unknown OS: %s", runtime.GOOS)
	}
	return cmds
}
