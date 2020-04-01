package cmd

import (
	"fmt"

	cm "github.com/clonerOpsTool/pkg/common"
	mdfy "github.com/clonerOpsTool/pkg/modify"
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
		fmt.Println("extendStorage called")
		config, err := cm.UploadStorageConf(stConf)
		if err != nil {
			fmt.Println(err)
			return
		}

		port, err := mdfy.GetLastPort(svName, config.Stores)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Last port for %s is:%d\n", svName, port)

		stores, err := mdfy.GetStoreCluster(svName, 11, config.Stores)
		if err != nil {
			fmt.Println(err)
			return
		}
		point, err := mdfy.GetLastPoint(stores)
		fmt.Printf("Last point for server %s, store 11 is:%d\n", svName, point)

	},
}

func init() {
	configEditorCmd.AddCommand(extendStorageCmd)

	extendStorageCmd.Flags().StringVarP(&svName, "server", "s", "", "FILL explanation")
	extendStorageCmd.MarkFlagRequired("server")
}
