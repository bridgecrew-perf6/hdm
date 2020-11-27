/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package apply

import (
	"github.com/spf13/cobra"
	"github.sec.samsung.com/m5-kim/hdm/pkg/cmd/common"
	"github.sec.samsung.com/m5-kim/hdm/pkg/cmd/config"
	"github.sec.samsung.com/m5-kim/hdm/pkg/cmd/exec"
	"github.sec.samsung.com/m5-kim/hdm/pkg/cmd/utils"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply [INDEX|ALIAS]",
	Short: "apply (install & upgrade) with helm command",
	Long: `apply (install & upgrade) with helm command.

Example:
  $ hdm apply 1

  or if you want to use alias, change index to alias
  
  $ hdm apply amf
`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := common.ConvertToIndexTarget(args[0])
		utils.CheckError(err)
		exec.ExecuteCommand(config.GetConfig().ApplyCommand(index))
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return config.GetTargetsToCompletion(), cobra.ShellCompDirectiveNoFileComp
	},
}

func CreateApplyCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(applyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
