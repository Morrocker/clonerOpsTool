package cmd

import (
	"github.com/spf13/cobra"
)

// modifyCmd represents the modify command
var modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modifica los parametros de uno o mas storages",
	Long: `modify permite modificar los parametros de los storages
en un archivo de configuracion de Block Server, entregando como resultado el archivo 
modificado o un archivo distinto, dependiendo de la opcion utilizada.

Para realizar las modificaciones se debe proveer a la herramienta un archivo con las
instrucciones a seguir. Este permite especificar el rango de storages que seran 
modificados y sus valores correspondientes. Las instrucciones permiten realizar varias 
modificaciones en cadena.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Printf("Modifying storage config %s according to instructions in %s\n", stConf, instFile)
		// config, err := cm.UploadStorageConf(stConf)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// instructions, err := cm.UploadInstructions(instFile)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		// Stores, err := mdfy.ExecInstr(config.Stores, instructions)
		// config.Stores = Stores

		// if outFile == "" {
		// 	cm.WriteJSON(stConf, config)
		// } else {
		// 	cm.WriteJSON(outFile, config)
		// }

		// spew.Dump(config)
		// spew.Dump(i)
	},
}

// var instFile, outFile string

func init() {
	configEditorCmd.AddCommand(modifyCmd)
	// modifyCmd.Flags().StringVarP(&instFile, "instruction", "i", "changes.json", "Archivo de intrucciones para el cambio en las configuraciones (requerido)")
	// modifyCmd.MarkFlagRequired("instruction")
	// modifyCmd.Flags().StringVarP(&outFile, "outputTo", "o", "", "Nombre del archivo resultante del cambio. Por defecto sobreescribe el archivo de origen")
}
