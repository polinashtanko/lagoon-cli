package cmd

import (
	"fmt"
	"os"

	"github.com/amazeeio/lagoon-cli/lagoon/environments"
	"github.com/amazeeio/lagoon-cli/output"
	"github.com/spf13/cobra"
)

// @TODO re-enable this at some point if more environment based commands are made availab;e
// EnvironmentFlags .
// type EnvironmentFlags struct {
// 	Name string `json:"name,omitempty"`
// }

// func parseEnvironmentFlags(flags pflag.FlagSet) EnvironmentFlags {
// 	configMap := make(map[string]interface{})
// 	flags.VisitAll(func(f *pflag.Flag) {
// 		if flags.Changed(f.Name) {
// 			configMap[f.Name] = f.Value
// 		}
// 	})
// 	jsonStr, _ := json.Marshal(configMap)
// 	parsedFlags := EnvironmentFlags{}
// 	json.Unmarshal(jsonStr, &parsedFlags)
// 	return parsedFlags
// }

var deleteEnvCmd = &cobra.Command{
	Use:     "environment",
	Aliases: []string{"e"},
	Short:   "Delete an environment",
	Run: func(cmd *cobra.Command, args []string) {
		// environmentFlags := parseEnvironmentFlags(*cmd.Flags()) //@TODO re-enable this at some point if more environment based commands are made availab;e
		if cmdProjectName == "" || cmdProjectEnvironment == "" {
			fmt.Println("Not enough arguments. Requires: project name and environment name")
			cmd.Help()
			os.Exit(1)
		}
		fmt.Println(fmt.Sprintf("Deleting %s-%s", cmdProjectName, cmdProjectEnvironment))

		if yesNo() {
			projectByName, err := environments.DeleteEnvironment(cmdProjectName, cmdProjectEnvironment)
			if err != nil {
				output.RenderError(err.Error(), outputOptions)
				os.Exit(1)
			}
			resultData := output.Result{
				Result: string(projectByName),
			}
			output.RenderResult(resultData, outputOptions)
		}
	},
}

func init() {
	deleteCmd.AddCommand(deleteEnvCmd)
}