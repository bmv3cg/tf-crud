package tfcrud

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bmv3cg/tf-crud/pkg/tfclient"
	"github.com/gosuri/uitable"
	"github.com/hashicorp/go-tfe"
	"github.com/spf13/viper"
)

// GetWorkspaceID is a function to retrive workspace ID of a workspace
func GetWorkspaceID(ctx context.Context, TfeWS string, TfeOrg string, Tfclient *tfe.Client) string {

	//Move to workspace list options
	wl, err := Tfclient.Workspaces.List(ctx, TfeOrg, tfe.WorkspaceListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, ws := range wl.Items {
		if ws.Name == TfeWS {
			return ws.ID
		}
	}

	return ""
}

// CreateWorkspace is a funciton to create a workspace in an organisation.
func CreateWorkspace(ctx context.Context, TfeWsName string, TfeOrg string, Tfclient *tfe.Client) {

	wsname := viper.GetString("wsname")
	org := viper.GetString("organisation")
	ws := GetWorkspaceID(tfclient.Ctx, wsname, org, tfclient.Tfclient)
	if ws != "" {
		fmt.Printf("Workspace %s already exists \n", wsname)
		os.Exit(1)
	}

	_, err := Tfclient.Workspaces.Create(ctx, TfeOrg, tfe.WorkspaceCreateOptions{
		Name: tfe.String(TfeWsName),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created workspace", TfeWsName)
}

// DeleteWorkspace is a fucntion to delete workspace in an organisation.
func DeleteWorkspace(ctx context.Context, TfeWsName string, TfeOrg string, Tfclient *tfe.Client) {

	//Delete  workspace
	err := Tfclient.Workspaces.Delete(ctx, TfeOrg, TfeWsName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted workspace", TfeWsName)
}

// DeleteWorkspaceID is a fucntion to delete a workspace with workspace ID
func DeleteWorkspaceID(ctx context.Context, TfeDelWS string, Tfclient *tfe.Client) string {

	//Create workspace
	err := Tfclient.Workspaces.DeleteByID(ctx, TfeDelWS)
	if err != nil {
		log.Fatal(err)
	}
	return "Workspace Deleted"
}

// ListWorkspace is a fucntion to list worksapce name and workpsace ID in a table
func ListWorkspace(ctx context.Context, TfeOrg string, Tfclient *tfe.Client) {

	type saveWs struct {
		WsName string
		WsID   string
	}

	WsList := make([]saveWs, 0)

	wl, err := Tfclient.Workspaces.List(ctx, TfeOrg, tfe.WorkspaceListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, ws := range wl.Items {
		WsList = append(WsList, saveWs{WsName: ws.Name, WsID: ws.ID})
	}

	table := uitable.New()
	table.MaxColWidth = 80
	table.Wrap = true
	table.Separator = "|"

	table.AddRow("", "-------------------------------", "-------------------------------", "")
	table.AddRow("", "Workspace Name", "Workspace ID", "")
	table.AddRow("", "-------------------------------", "-------------------------------", "")
	for _, WsList := range WsList {
		table.AddRow("", WsList.WsName, WsList.WsID, "")
	}
	table.AddRow("", "-------------------------------", "-------------------------------", "")
	fmt.Println(table)
}
