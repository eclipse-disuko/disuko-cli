// SPDX-FileCopyrightText: 2025 Mercedes-Benz Group AG and Mercedes-Benz AG
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"os"

	"github.com/eclipse-disuko/disuko-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Printf("%s. ", err)
		os.Exit(1)
	}
}
