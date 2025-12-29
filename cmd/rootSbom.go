// SPDX-FileCopyrightText: 2025 Mercedes-Benz Group AG and Mercedes-Benz AG
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/spf13/cobra"
)

var sbomCmd = &cobra.Command{
	Use:   "sbom",
	Short: "Execute a project version sbom command",
	Long:  `Executes the given command to project version sbom`,
	Args:  cobra.MinimumNArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {},
}
