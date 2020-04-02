package cmd

import (
	"github.com/spf13/cobra"
)

// singleScanCmd represents the singleScan command
var singleScanCmd = &cobra.Command{
	Use:   "singleScan",
	Short: "Escaneo acotado a dos dispositivos",
	Long: `singleScan realiza un analisis de trafico de red entre 2 dispositivos:
Un maestro que hace de servidor y un cliente que prueba el trafico. El resultado es
devuelto a traves de la terminal. No se escriben archivos de salida.`,
	Run: func(cmd *cobra.Command, args []string) {

		// 		stopMaster, mst, err := nt.StartMaster(cfg, port, scantime, location, master)
		// 		if err != nil {
		// 			fmt.Println(err)
		// 			return
		// 		}

		// 		server, err := nt.GetServer(cfg, client, location)

		// 		nt.RunScan(server, mst, port, scantime)

		// 		stopMaster()
	},
}

func init() {
	netTestCmd.AddCommand(singleScanCmd)

	// singleScanCmd.PersistentFlags().StringVarP(&master, "master", "m", "", "Nombre del servidor maestro (requerido)")
	// viper.BindPFlag("master", singleScanCmd.PersistentFlags().Lookup("master"))
	// singleScanCmd.MarkPersistentFlagRequired("master")

	// singleScanCmd.PersistentFlags().StringVarP(&client, "client", "c", "", "Nombre del cliente (requerido)")
	// viper.BindPFlag("client", singleScanCmd.PersistentFlags().Lookup("client"))
	// singleScanCmd.MarkPersistentFlagRequired("client")
}
