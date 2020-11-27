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
package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

const readme = `
HISTORY

v1.0

* hdm is migrated into golang.
  * no need to install python dependency anymore.
* hdm config file is moved to ~/.hdm/config.yaml
* hdm ui [INDEX] command is added/
* hdm install [INDEX] command is changed to hdm apply [INDEX]

v1.1

* support bash and zsh completion.
  * bash
    * source <(hdm completion bash) in your shell like kubectl.
  * zsh  
    * echo "autoload -U compinit; compinit" >> ~/.zshrc
    * hdm completion zsh > "${fpath[1]}/_hdm"
    * new shell takes effect.

v1.2

* add alias for deploy targets.
  * usage
alias:
  - amf
  - amf-test
* alias validation check when reading config file`

// versionCmd represents the version command

const Version = "v1.2"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show hdm version",
	Long: `show helm deploy manager version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Helm Deploy Manager %s made by m5.kim\n", Version)
		fmt.Println(readme)
	},
}

func CreateVersionCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
