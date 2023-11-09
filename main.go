/*
Copyright © 2023 alvtsky github.com/Ra-sky
*/
package main

import (
	"sshabu/cmd"
	"time"

	"github.com/carlmjohnson/versioninfo"
)

func main() {
	cmd.SetVersionInfo(versioninfo.Version, versioninfo.Revision, versioninfo.LastCommit.Format(time.RFC3339))
	cmd.Execute()
}
