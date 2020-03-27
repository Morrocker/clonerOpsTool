package cmd

import (
	"fmt"
	"time"

	xl "github.com/clonerOpsTool/methods/xlsx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var netScanCmd = &cobra.Command{
	Use:   "netScan",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("netScan called")
	},
}

var port, scantime, location string

func init() {
	rootCmd.AddCommand(netScanCmd)

	netScanCmd.PersistentFlags().StringVarP(&port, "port", "p", "8000", "A help for foo")
	viper.BindPFlag("port", netScanCmd.PersistentFlags().Lookup("port"))
	netScanCmd.PersistentFlags().StringVarP(&scantime, "scantime", "t", "20", "A help for foo")
	viper.BindPFlag("scantime", netScanCmd.PersistentFlags().Lookup("scantime"))
	netScanCmd.PersistentFlags().StringVarP(&location, "location", "l", "", "A help for foo")
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
