package cmd

import (
	"github.com/spf13/cobra"
)

var netTestCmd = &cobra.Command{
	Use:   "netTest",
	Short: "Scan to assess transfer rates between servers.",
	Long: `The netTest tool allows basic analisis to determine the data transfer rate between two or more targets in order to assess local network status. 
	
The scan makes use of the Iperf4 analisis tool, available on MacOsX and Linux systems and as such assumes that all devices involved on a test already have the tool installed.

The toolnetTest permite realizar un analisis basico de
la calidad de transferencia de informacion, a traves de la red,
entre 2 o mas dispositivos ubicados en la misma red local.

El analisis utiliza la herramienta de analisis Iperf, disponible en 
MacOsX y Linux y asume que todos los dispositivos involucrados en el
analisis pueden ejecutar la aplicacion.

La herramienta obtiene los dispositivos sobre los cuales potencialmente
se puede realizar un escaneo a partir del archivo de configuracion 
basico de la suite de herramientas. 

Por defecto el resultado del analisis es copiado a un archivo 
.xlsx de nombre FullScan.[fechaActual].xlsx o SingleScan[Maestro].[fechaActual].xlsx. 
Analisis entre 2 servidores no generan un archivo de salida.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("netTest called")
	},
}

var port, scantime, location string
var master, client string

func init() {
	rootCmd.AddCommand(netTestCmd)

	netTestCmd.PersistentFlags().StringVarP(&port, "port", "p", "8000", "puerto utilizado para realizar el analisis")
	netTestCmd.PersistentFlags().StringVarP(&scantime, "scantime", "t", "20", "tiempo que dura cada escaneo individual")
	netTestCmd.PersistentFlags().StringVarP(&location, "location", "l", "", "nombre de la red local donde se realiza el analisis (requerido)")
	netTestCmd.MarkPersistentFlagRequired("location")
}
