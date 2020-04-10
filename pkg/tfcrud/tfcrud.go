package tfcrud

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gosuri/uitable"
	"github.com/hashicorp/go-tfe"
	"k8s.io/klog"
)

// GetWorkspaceID is a function to retrive workspace ID of a workspace
func GetWorkspaceID(ctx context.Context, TfeWS string, TfeOrg string, Tfclient *tfe.Client) string {

	//Move to workspace list options
	wl, err := Tfclient.Workspaces.List(ctx, TfeOrg, tfe.WorkspaceListOptions{})
	if err != nil {
		klog.Fatal(err)
	}

	for _, ws := range wl.Items {
		if ws.Name == TfeWS {
			return ws.ID
		}
	}

	return ""
}

// CreateWorkspace is a function to create a workspace in an organisation.
func CreateWorkspace(ctx context.Context, TfeWsName string, TfeOrg string, Tfclient *tfe.Client) {

	ws := GetWorkspaceID(ctx, TfeWsName, TfeOrg, Tfclient)
	if ws != "" {
		klog.Fatalf("Workspace %s already exists \n", TfeWsName)
		os.Exit(1)
	}

	_, err := Tfclient.Workspaces.Create(ctx, TfeOrg, tfe.WorkspaceCreateOptions{
		Name: tfe.String(TfeWsName),
	})
	if err != nil {
		klog.Fatal(err)
	}
	klog.Info("Created workspace ", TfeWsName)
}

// DeleteWorkspace is a function to delete workspace in an organisation.
func DeleteWorkspace(ctx context.Context, TfeWsName string, TfeOrg string, Tfclient *tfe.Client) {
	err := Tfclient.Workspaces.Delete(ctx, TfeOrg, TfeWsName)
	if err != nil {
		klog.Fatal(err)
	}
	klog.Info("Deleted workspace", TfeWsName)
}

// DeleteWorkspaceID is a function to delete a workspace with workspace ID
func DeleteWorkspaceID(ctx context.Context, TfeDelWS string, Tfclient *tfe.Client) string {

	err := Tfclient.Workspaces.DeleteByID(ctx, TfeDelWS)
	if err != nil {
		klog.Fatal(err)
	}
	return "Workspace Deleted"
}

// ListWorkspace is a function to list worksapce name and workpsace ID in a tablular format
func ListWorkspace(ctx context.Context, TfeOrg string, Tfclient *tfe.Client) {

	type saveWs struct {
		WsName string
		WsID   string
	}

	WsList := make([]saveWs, 0)

	wl, err := Tfclient.Workspaces.List(ctx, TfeOrg, tfe.WorkspaceListOptions{})
	if err != nil {
		klog.Fatal(err)
	}

	for _, ws := range wl.Items {
		WsList = append(WsList, saveWs{WsName: ws.Name, WsID: ws.ID})
	}

	table := uitable.New()
	table.MaxColWidth = 80
	table.Wrap = true
	table.Separator = "|"

	table.AddRow("", "---------------------", "----------------------", "")
	table.AddRow("", "Workspace Name", "Workspace ID", "")
	table.AddRow("", "---------------------", "----------------------", "")
	for _, WsList := range WsList {
		table.AddRow("", WsList.WsName, WsList.WsID, "")
	}
	table.AddRow("", "---------------------", "----------------------", "")
	fmt.Println(table)
}

/*
SortWorkspace function sorts workspaces according to created time and list un-used
workspaces acoording to creation date of workspace.
*/
func SortWorkspace(ctx context.Context, TfeOrg string, delta int, Tfclient *tfe.Client) {

	wl, err := Tfclient.Workspaces.List(ctx, TfeOrg, tfe.WorkspaceListOptions{})
	if err != nil {
		klog.Fatal(err)
	}
	loc, _ := time.LoadLocation("UTC")
	createdAt := time.Now().In(loc).AddDate(0, 0, -delta)

	table := uitable.New()
	table.MaxColWidth = 80
	table.Wrap = true
	table.Separator = "|"
	table.AddRow("", "---------------------", "---------------------------------", "")
	table.AddRow("", "Workspace Name", "Creation time", "")
	table.AddRow("", "---------------------", "---------------------------------", "")
	for _, ws := range wl.Items {
		if ws.CreatedAt.Before(createdAt) && listConfig(ctx, ws.ID, Tfclient) == true {
			table.AddRow("", ws.Name, ws.CreatedAt, "")

		}
	}
	table.AddRow("", "---------------------", "---------------------------------", "")
	fmt.Println(table)
}

/*
listConfig function checks wheteher a configuration version exists for a workspace.
For unused workspace there will be no existing configuration version
For used workspaces there wil be multiple configuration versions.
*/
func listConfig(ctx context.Context, tfeWsID string, Tfclient *tfe.Client) (wsUsage bool) {

	options := tfe.ConfigurationVersionListOptions{}
	cvl, err := Tfclient.ConfigurationVersions.List(ctx, tfeWsID, options)
	if err != nil {
		klog.Fatal(err)
	}

	if cvl.TotalCount == 0 {
		return true
	}

	return false
}
