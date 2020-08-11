package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lastPointCmd represents the lastPoint command
var lastPointCmd = &cobra.Command{
	Use:   "lastPoint",
	Short: "Retrieve last point of the given store/server combination",
	Long:  `Simple tool to retrieve the last used point for a given storage in a storage configuration file. Last parameter given must be the target server.`,
	Run: func(cmd *cobra.Command, args []string) {
		config.Load(stConf)
		svName = args[0]
		fmt.Printf("Getting last point for store %d on %s\n", stNum, svName)
		point, e := config.GetLastPoint(svName, stNum)
		if e != nil {
			fmt.Println(e)
			return
		}
		fmt.Printf("Last point is:%d\n", point)
	},
}

func init() {
	configEditorCmd.AddCommand(lastPointCmd)
	lastPointCmd.Flags().StringVarP(&svName, "server", "s", "", "Target server where store is created")
	lastPointCmd.MarkFlagRequired("server")
	lastPointCmd.Flags().IntVarP(&stNum, "store", "t", 0, "Storage number to be checked")
	lastPointCmd.MarkFlagRequired("store")
}
