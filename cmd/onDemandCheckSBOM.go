// SPDX-FileCopyrightText: 2025 Mercedes-Benz Group AG and Mercedes-Benz AG
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"os"

	"github.com/mercedes-benz/disclosure-cli/conf"
	"github.com/mercedes-benz/disclosure-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var onDemandCheckSBOM = &cobra.Command{
	Use:   "sbomCheck [fileName]",
	Short: "On demand check for SBOM files",
	Long:  `Uploads a SBOM file and run a check against the project settings`,
	Run: func(cmd *cobra.Command, args []string) {
		fileName := ""
		if len(args) > 0 {
			fileName = args[0]
		} else {
			fmt.Println("Missing filename of SBOM upload")
			os.Exit(1)
		}
		msg := helper.SbomUploadFormData(helper.GetProjectAPIURL(conf.DefaultApiVersion, "/sbomcheck"), fileName, "")
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
