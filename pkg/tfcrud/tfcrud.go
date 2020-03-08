package tfcrud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/go-tfe"
)

// GetWorkspaceID is a function to retrive workspace ID of a workspace
func GetWorkspaceID(ctx context.Context, TfeWS string, Tfclient *tfe.Client) string {

	//Move to workspace list options
	wl, err := Tfclient.Workspaces.List(ctx, "tfe-cloud-spike", tfe.WorkspaceListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, ws := range wl.Items {
		if ws.Name == TfeWS {
			fmt.Println(ws.ID)
			return ws.ID
		}
	}

	return "Workspace not found"
}

// CreateWorkspace is a funciton to create a workspace in an organisation.
func CreateWorkspace(ctx context.Context, TfeWsName string, Tfclient *tfe.Client) string {

	//Create workspace
	_, err := Tfclient.Workspaces.Create(ctx, "tfe-cloud-spike", tfe.WorkspaceCreateOptions{
		Name: tfe.String(TfeWsName),
	})
	if err != nil {
		log.Fatal(err)
	}
	return "Workspace created"
}

// DeleteWorkspace is a fucntion to delete workspace in an organisation.
func DeleteWorkspace(ctx context.Context, TfeWsName string, Tfclient *tfe.Client) string {

	//Delete  workspace
	err := Tfclient.Workspaces.Delete(ctx, "tfe-cloud-spike", TfeWsName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Dleted workpspace")
	return "Workspace Deleted"
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
func ListWorkspace(ctx context.Context, Tfclient *tfe.Client) {

	wl, err := Tfclient.Workspaces.List(ctx, "tfe-cloud-spike", tfe.WorkspaceListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(" -----------------------------------")
	fmt.Printf("|%12s|%-20s|\n", "Workspace name", "Workspace ID")
	fmt.Println(" -----------------------------------")

	for _, ws := range wl.Items {

		fmt.Printf("|%-14s|%20s|\n", ws.Name, ws.ID)
	}

}
