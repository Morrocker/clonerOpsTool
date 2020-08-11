package cmd

import (
	"github.com/spf13/cobra"
)

// generateSlaveCmd represents the generateSlave command
var generateSlaveCmd = &cobra.Command{
	Use:   "generateSlave",
	Short: "Generates a slave block server configuration file from a master file.",
	Long:  `A slave storage configuration file is generated using a valid block server configuration file as the base reference. The stores generated`,
	Run: func(cmd *cobra.Command, args []string) {
		config.Load(stConf)
		slave := config.GenerateSlave(svName)
		slave.Write(svName + "SlaveConf.json")
	},
}

func init() {
	configEditorCmd.AddCommand(generateSlaveCmd)
	generateSlaveCmd.Flags().StringVarP(&svName, "server", "s", "", "Target server where store is created")
	generateSlaveCmd.MarkFlagRequired("server")
}
