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
package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"k8s.io/klog"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var verbosity int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tf-crud",
	Short: "A CLI tool to run CRUD operarion in Terraform cloud",
	Long: `A CLI tool for managing workspace creation, deletion, listing and managing 
	in Terraform cloud and Terraform enterprise`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	klog.V(2).Info("Using config file:", cfgFile)
}

func init() {

	var wsname string
	var organisation string

	cobra.OnInitialize(initConfig)
	klog.InitFlags(nil)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Logging flags
	rootCmd.PersistentFlags().AddGoFlag(flag.CommandLine.Lookup("v"))
	rootCmd.PersistentFlags().AddGoFlag(flag.CommandLine.Lookup("logtostderr"))
	rootCmd.PersistentFlags().Set("logtostderr", "true")
	//rootCmd.PersistentFlags().Set("v", "2")
	// Config flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tf-crud.yaml)")

	//TFE crud flags
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVar(&wsname, "wsname", "", "workspace name (required)")
	rootCmd.PersistentFlags().StringVar(&organisation, "organisation", "organisation", "Organisation name (required)")

	defer klog.Flush()
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

		// Search config in home directory with name ".tf-crud" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".tf-crud")
		viper.SetConfigType("yaml")
		viper.BindPFlag("wsname", rootCmd.Flags().Lookup("wsname"))
		viper.BindPFlag("organisation", rootCmd.Flags().Lookup("organisation"))
		rootCmd.MarkPersistentFlagRequired("organisation")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		klog.V(2).Infoln("Using config file", viper.ConfigFileUsed())
	}

}
