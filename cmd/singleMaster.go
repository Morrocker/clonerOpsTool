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
	Short: "Escaneo de una ubicacion, todos contra UN servidor maestro",
	Long: `singleMaster realiza un analisis de trafico de red entre todos los dispositivos
de una red local y un servidor maestro en la misma. El resultado es devuelto a traves de la 
terminal y adicionalmente se escribe un archivo de salida tipo .xlsx con el nombre 
SingleScan.[nombreMaestro].[fechaActual].xlsx`,
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
	},
}

func init() {
	netScanCmd.AddCommand(singleMasterCmd)

	singleMasterCmd.PersistentFlags().StringVarP(&master, "master", "m", "", "Nombre del servidor maestro (requerido)")
	viper.BindPFlag("master", singleMasterCmd.PersistentFlags().Lookup("master"))
	singleMasterCmd.MarkPersistentFlagRequired("master")
}
