// SPDX-FileCopyrightText: 2025 Mercedes-Benz Group AG and Mercedes-Benz AG
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"os"

	"github.com/eclipse-disuko/disuko-cli/conf"
	"github.com/eclipse-disuko/disuko-cli/pkg/domain"
	"github.com/eclipse-disuko/disuko-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var createVersionCmd = &cobra.Command{
	Use:   "create [name] [description]",
	Short: "Create version",
	Long:  `Creates a project version by name and description`,
	Run: func(cmd *cobra.Command, args []string) {
		data := domain.RequestCreateVersion{}
		if len(args) > 0 {
			data.Name = args[0]
		} else {
			fmt.Println("Version name is missing")
			os.Exit(1)
		}
		if len(args) > 1 {
			data.Description = args[1]
		}

		msg := helper.DiscoApiPost(helper.GetProjectAPIURL(conf.DefaultApiVersion, "/versions"), data)
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
