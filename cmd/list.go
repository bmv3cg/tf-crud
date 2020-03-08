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
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all workspace and workspace ID",
	Long:  `List all workspace in a Terraform cloud account with workspace name and workspace ID`,
	Run: func(cmd *cobra.Command, args []string) {
		//	tfcrud.GetWorkspaceID(cmd.InheritedFlags().Lookup("Wsname").Value.String(), tfclient.Ctx, tfclient.Tfclient)
		tfcrud.ListWorkspace(tfclient.Ctx, tfclient.Tfclient)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
