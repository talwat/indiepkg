package main

type Bin struct {
	Installed []string
	InSource  []string `json:"in_source"`
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
	Name        string
	Author      string
	Description string
	URL         string
	Download    map[string]interface{}
	Version     string
	License     string
	Language    string
	Branch      string
	Bin         *Bin
	Manpages    []string
	Deps        *Deps
	FileDeps    *Deps `json:"file_deps"`
	Commands    *OSCommands
	ConfigPaths []string `json:"config_paths"`
	Notes       []string
}
