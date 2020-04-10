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
	"github.com/bmv3cg/tf-crud/pkg/tfclient"
	"github.com/bmv3cg/tf-crud/pkg/tfcrud"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var delta int

// emptyWsCmd represents the sort command
var emptyWsCmd = &cobra.Command{
	Use:     "emptyws",
	Aliases: []string{"ews"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tfcrud.SortWorkspace(tfclient.Ctx, viper.GetString("organisation"), viper.GetInt("delta"), tfclient.Tfclient)
	},
}

func init() {
	rootCmd.AddCommand(emptyWsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	emptyWsCmd.Flags().IntVar(&delta, "delta", 0, "List workspace older than delta")

}
