package cmd

import (
	"time"

	xl "github.com/clonerOpsTool/pkg/xlsx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var netScanCmd = &cobra.Command{
	Use:   "netScan",
	Short: "Escaneo del estado de la red.",
	Long: `netScan permite realizar un analisis basico de
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
		// fmt.Println("netScan called")
	},
}

var port, scantime, location string
var master, client string

func init() {
	rootCmd.AddCommand(netScanCmd)

	netScanCmd.PersistentFlags().StringVarP(&port, "port", "p", "8000", "puerto utilizado para realizar el analisis")
	viper.BindPFlag("port", netScanCmd.PersistentFlags().Lookup("port"))
	netScanCmd.PersistentFlags().StringVarP(&scantime, "scantime", "t", "20", "tiempo que dura cada escaneo individual")
	viper.BindPFlag("scantime", netScanCmd.PersistentFlags().Lookup("scantime"))
	netScanCmd.PersistentFlags().StringVarP(&location, "location", "l", "", "nombre de la red local donde se realiza el analisis (requerido)")
	viper.BindPFlag("location", netScanCmd.PersistentFlags().Lookup("location"))
	netScanCmd.MarkPersistentFlagRequired("location")
}

func writeFile(filename, sheetname string, data [][]string) {
	today := time.Now().Format("2006-01-02")
	var xlsx = xl.Xlsx{
		Filename: filename + "." + today + ".xlsx",
	}
	var sheet = xl.Sheet{
		Name: sheetname,
		Data: data,
	}
	xlsx.Sheets = append(xlsx.Sheets, sheet)
	xl.WriteXlsx(xlsx)
}
