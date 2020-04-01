package cmd

import (
	"github.com/spf13/cobra"
)

// configEditorCmd represents the configEditor command
var configEditorCmd = &cobra.Command{
	Use:   "configEditor",
	Short: "Editor de configuracion de Block Server",
	Long: `configEditor permite la edicion de archivos de configuracion 
utilizados por el servicio Block Server. Requiere de un archivo de configuracion como
entrada, a partir del cual se pueden realizar operaciones, como el cambio de un parametro
a lo largo del todo el archivo. 

La herramienta permite cierta granularidad en los cambios a realizar pudiendo restringir 
los storagepoints cuyos parametros son afectados.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("configEditor called")
	},
}

var stConf, svName string

func init() {
	rootCmd.AddCommand(configEditorCmd)

	configEditorCmd.PersistentFlags().StringVarP(&stConf, "storeConf", "c", "storage_config.json", "Archivo de configuracion de Block Server de entrada (requerido)")
	configEditorCmd.MarkFlagRequired("storeConf")
}
