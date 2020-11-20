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
package list

import (
	"github.com/spf13/cobra"
	"github.sec.samsung.com/m5-kim/hdm/pkg/cmd/config"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "helm list managed by hdm",
	Long: `Show helm list managed by hdm.
It will display every commands which used by this program.

Example:
  hdm list
`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		config.GetConfig().ListAllCommand()
	},
}

func CreateListCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(listCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
