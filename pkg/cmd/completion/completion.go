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
package completion

import (
	"github.sec.samsung.com/m5-kim/hdm/pkg/cmd/utils"
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion [bash]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

$ source <(hdm completion bash)

# To load completions for each session, execute once:
Linux:
  $ yourprogram completion bash > /etc/bash_completion.d/hdm
MacOS:
  $ yourprogram completion bash > /usr/local/etc/bash_compleion.d/hdm

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ hdm completion zsh > "${fpath[1]}/_yourprogram"

# You will need to start a new shell for this setup to take effect.
`,
	DisableFlagsInUseLine: true,
	Args: cobra.ExactArgs(1),
	ValidArgs:             []string{"bash", "zsh"},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			err := cmd.Root().GenBashCompletion(os.Stdout)
			utils.CheckError(err)
		case "zsh":
			err := cmd.Root().GenZshCompletion(os.Stdout)
			utils.CheckError(err)
		}
	},
}

func CreateCompletionCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(completionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
