package cmd

/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// netScanCmd represents the netScan command
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

var port, scantime int
var fullscan bool
var location string

func init() {
	rootCmd.AddCommand(netScanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	netScanCmd.PersistentFlags().IntVarP(&port, "port", "p", 8000, "A help for foo")
	viper.BindPFlag("port", netScanCmd.PersistentFlags().Lookup("port"))
	netScanCmd.PersistentFlags().IntVarP(&scantime, "scantime", "t", 20, "A help for foo")
	viper.BindPFlag("scantime", netScanCmd.PersistentFlags().Lookup("scantime"))
	netScanCmd.PersistentFlags().BoolVarP(&fullscan, "fullscan", "f", false, "A help for foo")
	viper.BindPFlag("fullscan", netScanCmd.PersistentFlags().Lookup("fullscan"))
	netScanCmd.PersistentFlags().StringVarP(&location, "location", "l", "", "A help for foo")
	viper.BindPFlag("location", netScanCmd.PersistentFlags().Lookup("location"))
	netScanCmd.MarkPersistentFlagRequired("location")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// netScanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
