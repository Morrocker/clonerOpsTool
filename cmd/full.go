package cmd

import (
	"fmt"

	nt "github.com/clonerOpsTool/pkg/netscan"
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
	},
}

func init() {
	netScanCmd.AddCommand(fullCmd)

}
