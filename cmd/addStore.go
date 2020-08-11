package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addStoreCmd represents the addStore command
var addStoreCmd = &cobra.Command{
	Use:   "addStore",
	Short: "Add a new stores to a server config",
	Long:  `addStore creates a new set of storages on the base storage configuration json file. The store can be a new set in which case the store will be created from point 1 to point X. Extending an existing store to a point N is also possible. If the store is part of the master file the -m flag should be provided lo ensure it's run field is set to true.`,
	Run: func(cmd *cobra.Command, args []string) {
		config.Load(stConf)
		fmt.Printf("addStore called. \nserver: %s\nstore:%d\npoint:%d\nextend:%v\n", svName, stNum, ptNum, extend)
		if extend {
			if e := config.ExtendStore(svName, stNum, ptNum, mflag); e != nil {
				fmt.Println(e)
				return
			}
		} else {
			if e := config.AddStore(svName, stNum, 1, ptNum, mflag); e != nil {
				fmt.Println(e)
				return
			}
		}
		// spew.Dump(config.Stores)
	},
}

var extend, mflag bool

func init() {
	configEditorCmd.AddCommand(addStoreCmd)
	addStoreCmd.Flags().StringVarP(&svName, "server", "s", "", "Target server where store is created")
	addStoreCmd.MarkFlagRequired("server")
	addStoreCmd.Flags().IntVarP(&ptNum, "toPoint", "p", 0, "Target point to reach during creation")
	addStoreCmd.MarkFlagRequired("toPoint")
	addStoreCmd.Flags().IntVarP(&stNum, "store", "t", 0, "Number of store to be created or extended")
	addStoreCmd.MarkFlagRequired("store")
	addStoreCmd.Flags().BoolVarP(&mflag, "master", "m", false, "This flag should be set if the created store is part of the master")
	addStoreCmd.Flags().BoolVarP(&extend, "extend", "e", false, "This flag switches the action from new creation to extend existing store")
}
