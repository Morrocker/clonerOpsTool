package cmd

import (
	"fmt"
	"os"

	ns "github.com/clonerOpsTool/pkg/netscan"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var cfg ns.ServerList

var rootCmd = &cobra.Command{
	Use:   "clonerOpsTool",
	Short: "",
	Long: `clonerOpsTool es una herramienta desarrollada para asistir en las tareas 
rutinarias realizadas por el equipo de operaciones. Esta compuesto de multiples herramientas
individuales que han sido necesarias para automatizar las tareas de operaciones en la empresa
Cloner Spa.

Para su uso se estabecen los siguientes supuestos:
- El dispositivo donde se ejecuta la herramienta cuenta con acceso SSH a todos los dispositivos
sobre los cuales se realizan operaciones. Adicionalmente, muchas de las operaciones requieren 
sudo para su ejecucion
- Los dispositivos sobre los cuales se pueden realizar operaciones se deben especificar en un 
archivo de configuracion y se les hace referencia por nombre
`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Archivo de configuracion basal, por defecto es 'clonerOpsTool.yaml'.")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".clonerOpsTool" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".clonerOpsTool")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	// Unmarshal config from viper into a configuration variable
	err := viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Printf("unable to decode into struct: %v", err)
	}
}
