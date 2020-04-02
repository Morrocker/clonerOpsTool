package cmd

import (
	"github.com/spf13/cobra"
)

// extendStorageCmd represents the extendStorage command
var extendStorageCmd = &cobra.Command{
	Use:   "extendStorage",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("extendStorage called")
		// config, err := cm.UploadStorageConf(stConf)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		// config.Stores, err = mdfy.ExtendStore(svName, stNum, toPoint, isMaster, config.Stores)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		// cm.WriteJSON(stConf, config)
	},
}

// var stNum, toPoint int
// var isMaster bool

func init() {
	configEditorCmd.AddCommand(extendStorageCmd)

	// extendStorageCmd.Flags().StringVarP(&svName, "server", "s", "", "FILL explanation")
	// extendStorageCmd.MarkFlagRequired("server")
	// extendStorageCmd.Flags().IntVarP(&stNum, "store", "t", 0, "FILL explanation")
	// extendStorageCmd.MarkFlagRequired("store")
	// extendStorageCmd.Flags().IntVarP(&toPoint, "toPoint", "p", 0, "FILL explanation")
	// extendStorageCmd.MarkFlagRequired("toPoint")
	// extendStorageCmd.Flags().BoolVarP(&isMaster, "isMaster", "m", false, "FILL explanation")
}
