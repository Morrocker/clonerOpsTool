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

		_, err = mdfy.ExecInstr(config.Stores, instructions)
		// Stores, err := mdfy.ExecInstr(config.Stores, instructions)
		// spew.Dump(config)

		// NewStores, err := mdfy.Key(key, value, config.Stores, i)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		// spew.Dump(Stores[0])
		// spew.Dump(i)

	},
}

var instFile, stConf string

func init() {
	configEditorCmd.AddCommand(modifyCmd)

	// modifyCmd.Flags().StringVarP(&key, "key", "k", "", "Help message for toggle")
	// modifyCmd.Flags().StringVarP(&value, "value", "v", "", "Help message for toggle")
	// modifyCmd.MarkFlagRequired("key")
	// modifyCmd.MarkFlagRequired("value")
	modifyCmd.Flags().StringVarP(&instFile, "instruction", "i", "changes.json", "Help message for toggle")
}
