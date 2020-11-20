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
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.sec.samsung.com/m5-kim/hdm/pkg/cmd/apply"
	"github.sec.samsung.com/m5-kim/hdm/pkg/cmd/completion"
	"github.sec.samsung.com/m5-kim/hdm/pkg/cmd/config"
	hdmdelete "github.sec.samsung.com/m5-kim/hdm/pkg/cmd/delete"
	"github.sec.samsung.com/m5-kim/hdm/pkg/cmd/list"
	"github.sec.samsung.com/m5-kim/hdm/pkg/cmd/status"
	"github.sec.samsung.com/m5-kim/hdm/pkg/cmd/ui"
	"github.sec.samsung.com/m5-kim/hdm/pkg/cmd/version"
	"os"
)

var cfgFile string


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hdm",
	Short: "helm-deploy-manager",
	Long: `helm deploy manager in golang.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

func GetRootCmd() *cobra.Command {
	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hdm/config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	list.CreateListCommand(rootCmd)
	apply.CreateApplyCommand(rootCmd)
	hdmdelete.CreateDeleteCommand(rootCmd)
	status.CreateStatusCommand(rootCmd)
	ui.CreateUICommand(rootCmd)
	version.CreateVersionCommand(rootCmd)
	completion.CreateCompletionCommand(rootCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	config.NewConfig(cfgFile)
}
