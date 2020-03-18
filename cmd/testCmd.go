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
	"log"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// testCmdCmd represents the testCmd command
var testCmdCmd = &cobra.Command{
	Use:   "testCmd",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("testCmd called")
		callDf()
	},
}

type config struct {
	Servers []server
}

type server struct {
	Name       string
	Category   string
	HostIperf  bool
	Location   string
	Port       int
	LocalIP    string
	VpnIP      string
	ExternalIP string
}

var author string

func init() {
	rootCmd.AddCommand(testCmdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// viper.BindPFlag("toggle", rootCmd.Flags().Lookup("toggle"))
	testCmdCmd.Flags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")
	viper.BindPFlag("author", testCmdCmd.Flags().Lookup("author"))
}

func callDf() {
	var C config
	cmd := exec.Command("ssh", "-p 2279", "192.168.201.127", "hostname")
	out, err := cmd.Output()
	fmt.Printf("output:%s", out)
	log.Printf("Command finished with error: %v", err)
	err = viper.Unmarshal(&C)
	// str := viper.GetString("servers")
	// for _, server := range C.Servers {
	// 	fmt.Println(server)
	// }
	fmt.Println(viper.GetString("author"))
}
