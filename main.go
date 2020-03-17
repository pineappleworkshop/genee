package main

import (
	"genee/cmd"
)

func main() {
	cmd.Execute()
}

// todo: genee project from git repo
// todo: implement verbosity flag
// todo: implement a debug flag
// todo: stop execution if root folder exists
// todo: implement a force flag
// todo: parse permissions with folder/file generation
// todo: must run in the goprojs with the *.yaml file located there (follow up)
// todo: clean up folder/file generation on failure
// todo: change genee gen -> genee project
// todo: impelement genee file
// todo: deploy on brew
// todo: deploy on apk-get
// todo: deploy on apt-get
// todo: deploy on yum
// todo: genee configuration file, shell config, template wrappers
