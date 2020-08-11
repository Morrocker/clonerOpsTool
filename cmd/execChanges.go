package cmd

import (
	"fmt"

	e "github.com/clonerOpsTool/pkg/confeditor"
	"github.com/spf13/cobra"
)

// execChangesCmd represents the execChanges command
var execChangesCmd = &cobra.Command{
	Use:   "execChanges",
	Short: "Executes a series of changes on the target configuration file",
	Long:  `This tool allows the execution of several actions on the storage configuration at once. It relies on a "instructions" JSON file that executes each change in order. At the end of every instruction the result is checked for possible insconsistencies in order to avoid issues. The possible modifications are: adding a new store, extending an extending store, changing any parameter to a set of stores.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("execChanges called")
		if e := config.Load(stConf); e != nil {
			fmt.Println(e)
			return
		}
		if e := instructions.Load(insFile); e != nil {
			fmt.Println(e)
			return
		}
		if e := instructions.Run(&config); e != nil {
			fmt.Println(e)
			return
		}
		setOutput()
		config.Write(output)
	},
}

var (
	instructions e.Instructions
	insFile      string
)

func init() {
	configEditorCmd.AddCommand(execChangesCmd)

	execChangesCmd.Flags().StringVarP(&insFile, "instructions", "i", "changes.json", "Sets filename that specifies instructions that will be executed.")
	execChangesCmd.MarkFlagRequired("instructions")
}
