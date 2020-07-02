/*
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */

// Package cmd is the one containing all the cli commands
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/torresashjian/labs-kubectl/version"
)

var (
	// versionCmd prints out the current cli version
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the app version",
		Run:   printVersion,
	}
)

// getVersion return version of CLI/node and commit hash
func getVersion() string {
	v := version.Version
	if version.GitCommit != "" {
		v = v + "-" + version.GitCommit
	}
	return v
}

// printVersion prints the version
func printVersion(cmd *cobra.Command, args []string) {
	v := getVersion()
	fmt.Println(v)
}
