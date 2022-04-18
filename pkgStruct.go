package main

type Bin struct {
	Installed []string
	In_source []string
}

type Commands struct {
	Install   []string
	Uninstall []string
	Update    []string
}

type OSCommands struct {
	All    *Commands
	Linux  *Commands
	Darwin *Commands
}

type Deps struct {
	All    []string
	Linux  []string
	Darwin []string
}

type Package struct {
	Name         string
	Author       string
	Description  string
	Url          string
	Download     map[string]interface{}
	Version      string
	License      string
	Branch       string
	Bin          *Bin
	Manpages     []string
	Deps         *Deps
	Commands     *OSCommands
	Config_paths []string
	Notes        []string
}
