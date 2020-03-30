package cmd

import (
	"fmt"

	nt "github.com/clonerOpsTool/pkg/netscan"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// singleScanCmd represents the singleScan command
var singleScanCmd = &cobra.Command{
	Use:   "singleScan",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		stopMaster, mst, err := nt.StartMaster(cfg, port, scantime, location, master)
		if err != nil {
			fmt.Println(err)
			return
		}

		server, err := nt.GetServer(cfg, client, location)

		nt.RunScan(server, mst, port, scantime)

		stopMaster()
	},
}

func init() {
	netScanCmd.AddCommand(singleScanCmd)

	singleScanCmd.PersistentFlags().StringVarP(&master, "master", "m", "", "A help for foo")
	viper.BindPFlag("master", singleScanCmd.PersistentFlags().Lookup("master"))
	singleScanCmd.MarkPersistentFlagRequired("master")

	singleScanCmd.PersistentFlags().StringVarP(&client, "client", "c", "", "A help for foo")
	viper.BindPFlag("client", singleScanCmd.PersistentFlags().Lookup("client"))
	singleScanCmd.MarkPersistentFlagRequired("client")
}
