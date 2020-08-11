package cmd

import (
	"fmt"

	nt "github.com/clonerOpsTool/pkg/netscan"
	x "github.com/clonerOpsTool/pkg/xlsx"
	"github.com/spf13/cobra"
)

var fullScanCmd = &cobra.Command{
	Use:   "fullScan",
	Short: "Crosswise scan across the local network",
	Long:  `fullScan launches a scan amongst all servers marked on the same local network. The scan uses an everyone vs everyone approach, so it's the most extensive scan available. The result is output through the terminal, as well as an output file which consists of an .xlsx file with a matrix showing all pairings done. The output file by default has the form FullScan.[date].xlsx.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\nDoing full location scan on %s\n", location)
		var data [][]string
		var m nt.ScanMaster
		var excel x.Xlsx

		excel.SetName("Fullscan")

		header := nt.CreateHeader(cfg, location)
		data = append(data, header)

		for _, mServer := range cfg.Servers {
			if mServer.Site != location {
				continue
			}
			err := m.SetIMaster(port, scantime, location, mServer.Name, cfg.Servers)
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = m.StartMaster()
			if err != nil {
				fmt.Println(err)
				continue
			}

			row, err := m.RunTest(cfg.Servers)
			if err != nil {
				fmt.Println(err)
				continue
			}
			data = append(data, row)
			err = m.StopMaster()
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

		excel.AddSheet(location+" scan", data)
		err := excel.WriteXlsx()
		if err != nil {
			fmt.Println(err)
		}
		// writeFile("FullScan", "data", data)
		// spew.Dump(m)
		// spew.Dump(data)
	},
}

func init() {
	netTestCmd.AddCommand(fullScanCmd)

}
