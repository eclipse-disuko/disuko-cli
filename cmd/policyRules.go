// SPDX-FileCopyrightText: 2025 Mercedes-Benz Group AG and Mercedes-Benz AG
//
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/eclipse-disuko/disuko-cli/conf"
	"github.com/eclipse-disuko/disuko-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var policyRulesCmd = &cobra.Command{
	Use:   "policyrules",
	Short: "Returning the project policy rules",
	Long:  `The policy rules of the project`,
	Run: func(cmd *cobra.Command, args []string) {
		msg := helper.DiscoApiGet(helper.GetProjectAPIURL(conf.DefaultApiVersion, "/policyrules"))
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
