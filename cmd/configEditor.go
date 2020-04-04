package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	json "github.com/clonerOpsTool/pkg/json"
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
		config, _ := json.UploadStorageConf(stConf)
		config.GetStoresData()
		config.SortStores()
		// port, _ := config.GetLastPort("alpha")
		err := config.AddStore("alpha", 35, 5, 0, true)
		fmt.Println(err)
		// config.Check()

		// spew.Dump(config)

	},
}

var stConf, svName string

func init() {
	rootCmd.AddCommand(configEditorCmd)

	configEditorCmd.PersistentFlags().StringVarP(&stConf, "storeConf", "c", "storage_config.json", "Archivo de configuracion de Block Server de entrada (requerido)")
	configEditorCmd.MarkFlagRequired("storeConf")
}
