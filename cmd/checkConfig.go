package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// checkConfigCmd represents the checkConfig command
var checkConfigCmd = &cobra.Command{
	Use:   "checkConfig",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking config for issues")
		config.Load(stConf)
		fmt.Println(stConf)
		if issues := config.Check(); issues {
			fmt.Println("Issues found with config")
		}

		fmt.Println("No issues found on target config")
		// spew.Dump(config)
		return
	},
}

func init() {
	configEditorCmd.AddCommand(checkConfigCmd)

}
