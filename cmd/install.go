/*
* Copyright © 2020. TIBCO Software Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */

// Package cmd is the one containing all the cli commands
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// installCmd installs the application
	installCmd = &cobra.Command{
		Use:   "install",
		Short: "Install the application",
		Run:   printVersion,
	}
)
