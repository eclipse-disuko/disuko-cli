// SPDX-FileCopyrightText: 2025 Mercedes-Benz Group AG and Mercedes-Benz AG
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"os"

	"github.com/eclipse-disuko/disuko-cli/conf"
	"github.com/eclipse-disuko/disuko-cli/pkg/domain"
	"github.com/eclipse-disuko/disuko-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var sbomTagCmd = &cobra.Command{
	Use:   "tag [sbomId] [tag]",
	Short: "Add tag to a sbom",
	Long:  `Add tag to a sbom `,
	Run: func(cmd *cobra.Command, args []string) {
		projectVersion := conf.Config.ProjectVersion
		data := domain.RequestCreateTag{}
		sbomId := ""

		if len(args) > 1 {
			sbomId = args[0]
			data.Tag = args[1]
		} else {
			helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString("[sbomId] or [tag] is missing in input"))
			os.Exit(1)
		}

		msg := helper.DiscoApiPut(helper.GetProjectVersionAPIURL(conf.DefaultApiVersion, projectVersion, "sboms/"+sbomId+"/tag"), data)
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
