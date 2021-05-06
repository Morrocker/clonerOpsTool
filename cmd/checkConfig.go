package cmd

import (
	"github.com/morrocker/log"
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
		log.Task("Checking config %s for issues", stConf)
		config.Load(stConf)
		if err := config.Check(); err != nil {
			log.Info("Issues found on target config")
		}

		log.Info("No issues found on target config")
		return
	},
}

func init() {
	configEditorCmd.AddCommand(checkConfigCmd)

}
