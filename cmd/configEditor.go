package cmd

import (
	"fmt"

	cm "github.com/clonerOpsTool/methods/common"
	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

// configEditorCmd represents the configEditor command
var configEditorCmd = &cobra.Command{
	Use:   "configEditor",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("configEditor called")
		config, err := cm.UploadJSON(stConf)
		if err != nil {
			fmt.Println(err)
			return
		}
		spew.Dump(config)
	},
}

var stConf string

func init() {
	rootCmd.AddCommand(configEditorCmd)

	configEditorCmd.PersistentFlags().StringVarP(&stConf, "storageconf", "s", "storage_config.json", "A help for foo")
	// viper.BindPFlag("storageconf", configEditorCmd.PersistentFlags().Lookup("storageconf"))
}
