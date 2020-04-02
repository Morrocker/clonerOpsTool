package cmd

import (
	"fmt"

	nt "github.com/clonerOpsTool/pkg/netscan"
	x "github.com/clonerOpsTool/pkg/xlsx"
	"github.com/spf13/cobra"
)

var fullCmd = &cobra.Command{
	Use:   "full",
	Short: "Escaneo cruzado entre todos los dispositivos de una red local",
	Long: `singleMaster realiza un analisis cruzado de trafico de red entre todos los dispositivos
de una red local. El resultado es devuelto a traves de la terminal y adicionalmente se escribe un 
archivo de salida tipo .xlsx con el nombre FullScan.[fechaActual].xlsx`,
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
	netTestCmd.AddCommand(fullCmd)

}
