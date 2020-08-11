package cmd

import (
	"time"

	"github.com/spf13/cobra"

	e "github.com/clonerOpsTool/pkg/confeditor"
)

// configEditorCmd represents the configEditor command
var configEditorCmd = &cobra.Command{
	Use:   "configEditor",
	Short: "Block server storage configuration file editor",
	Long: `configEditor allows easier and consistent modification of a Block server storage configuration file. It uses an existing blockserver configuration upon which new stores can be added, existing stores can be extended or modified. Most commands include automatic sorting of the stores and consistency checks, but these methods also exist as individual command in order to assist manual editing of the file. 

	Changes can be made individually. However for multiple additions and modifications it is highly suggested that the user takes advantage of the execChanges command, which takes an Instructions JSON file that contains a list of all changes to be made, that are executed in order, with consistency checks at every step.

	If manual changes are required the tool can do single instance checking of the store consistency, as well as obtaining a few relevant data point: Last used port by a storage related to a single server or last point used by a store on a server.

	Finally, the editor can output a slave configuration file, derived from a master, to make sure that master and slave confs are in sync.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var (
	config                 e.StorageConfig
	stConf, svName, output string
	stNum, ptNum           int
)

func init() {
	rootCmd.AddCommand(configEditorCmd)

	configEditorCmd.PersistentFlags().StringVarP(&stConf, "stConf", "c", "storage_config.json", "input file: Block server storage configuration JSON (required)")
	configEditorCmd.MarkFlagRequired("stConf")
	configEditorCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "storage configuration output filename(path)")
}

func setOutput() {
	if output != "" {
		return
	}
	date := time.Now().Format("2006-01-02")
	output = stConf + "." + date + ".out.json"
}
