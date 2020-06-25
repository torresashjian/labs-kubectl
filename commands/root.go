/*
* Copyright © 2018. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */

// Package commands is the one containing all the cli commands
package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(VersionCmd)
}

// RootCmd is the root command for dovetail cli
var RootCmd = &cobra.Command{
	Use:   "labskubectl",
	Short: "labskubectl is a flexible cli for multi-cloud kubernetes infrastructure",
	Long:  `labskubectl is a flexible cli for multi-cloud kubernetes infrastructure`,
}

// Execute is the starting point
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
