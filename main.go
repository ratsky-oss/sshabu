/*
Copyright © 2023 alvtsky github.com/Ra-sky
*/
package main

import (
	"sshabu/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cmd.SetVersionInfo(version, commit, date)
	cmd.Execute()
}