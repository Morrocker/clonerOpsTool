package cmd

import (
	"fmt"

	cm "github.com/clonerOpsTool/pkg/common"
	mdfy "github.com/clonerOpsTool/pkg/modify"
	"github.com/spf13/cobra"
)

// modifyCmd represents the modify command
var modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Modifying storage config %s according to instructions in %s\n", stConf, instFile)
		config, err := cm.UploadStorageConf(stConf)
		if err != nil {
			fmt.Println(err)
			return
		}
		instructions, err := cm.UploadInstructions(instFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		Stores, err := mdfy.ExecInstr(config.Stores, instructions)
		config.Stores = Stores

		if outFile == "" {
			cm.WriteJSON(stConf, config)
		} else {
			cm.WriteJSON(outFile, config)
		}

		// spew.Dump(config)
		// spew.Dump(i)
	},
}

var instFile, outFile, stConf string

func init() {
	configEditorCmd.AddCommand(modifyCmd)
	modifyCmd.Flags().StringVarP(&instFile, "instruction", "i", "changes.json", "Help message for toggle")
	modifyCmd.Flags().StringVarP(&outFile, "outputTo", "o", "", "Help message for toggle")
}
