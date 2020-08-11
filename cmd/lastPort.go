package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lastPortCmd represents the lastPort command
var lastPortCmd = &cobra.Command{
	Use:   "lastPort",
	Short: "Retrieve last port available for use in the given storage configuration file.",
	Long:  `Simple tool to retrieve the last available port used by a server in a given storage configuration file. Last parameter given must be the target server.`,
	Run: func(cmd *cobra.Command, args []string) {
		config.Load(stConf)
		fmt.Printf("Getting last port for: %v\n\n", args)
		for _, server := range args {
			port, e := config.GetLastPort(server)
			if e != nil {
				fmt.Println(e)
			}
			fmt.Printf("Last port for %s is:%d\n", server, port)

		}
	},
}

func init() {
	configEditorCmd.AddCommand(lastPortCmd)
}
