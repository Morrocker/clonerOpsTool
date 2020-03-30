package cmd

import (
	"fmt"
	"strings"

	nt "github.com/clonerOpsTool/pkg/netscan"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// singleMasterCmd represents the singleMaster command
var singleMasterCmd = &cobra.Command{
	Use:   "singleMaster",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nDoing single master scan")
		var data [][]string
		header := nt.CreateHeader(cfg, location)
		data = append(data, header)

		stopMaster, mst, err := nt.StartMaster(cfg, port, scantime, location, master)
		if err != nil {
			fmt.Println(err)
			return
		}

		row := nt.ScanServers(mst, port, scantime, location, cfg)
		data = append(data, row)

		stopMaster()
		writeFile("SingleScan"+strings.Title(mst.Name), "data", data)
		// fmt.Printf("%v", data)
	},
}

var master string

func init() {
	netScanCmd.AddCommand(singleMasterCmd)

	singleMasterCmd.PersistentFlags().StringVarP(&master, "master", "m", "", "A help for foo")
	viper.BindPFlag("master", singleMasterCmd.PersistentFlags().Lookup("master"))
	singleMasterCmd.MarkPersistentFlagRequired("master")
}
