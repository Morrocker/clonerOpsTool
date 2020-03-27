package cmd

import (
	"fmt"

	nt "github.com/clonerOpsTool/methods/netscan"
	"github.com/spf13/cobra"
)

var fullCmd = &cobra.Command{
	Use:   "full",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nDoing full location scan")
		var data [][]string
		header := nt.CreateHeader(cfg, location)
		data = append(data, header)

		for _, mServer := range cfg.Servers {
			if mServer.Location != location {
				continue
			}

			stopMaster, mst, err := nt.StartMaster(cfg, port, scantime, location, mServer.Name)
			if err != nil {
				fmt.Println(err)
				continue
			}
			row := nt.ScanServers(mst, port, scantime, location, cfg)
			data = append(data, row)
			stopMaster()
		}
		writeFile("FullScan", "data", data)
		// fmt.Printf("%v", data)
	},
}

func init() {
	netScanCmd.AddCommand(fullCmd)

}
