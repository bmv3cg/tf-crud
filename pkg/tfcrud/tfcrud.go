package tfcrud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/go-tfe"
)

func GetWorkspaceID(TfeWS string, ctx context.Context, Tfclient *tfe.Client) string {

	//Move to workspace list options
	wl, err := Tfclient.Workspaces.List(ctx, "tf-cloud", tfe.WorkspaceListOptions{})
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

func CreateWorkspace(TfeWsName string, ctx context.Context, Tfclient *tfe.Client) string {

	//Create workspace
	_, err := Tfclient.Workspaces.Create(ctx, "tf-cloud", tfe.WorkspaceCreateOptions{
		Name: tfe.String(TfeWsName),
	})
	if err != nil {
		log.Fatal(err)
	}
	return "Workspace created"
}

func DeleteWorkspace(TfeWsName string, ctx context.Context, Tfclient *tfe.Client) string {

	//Delete  workspace
	err := Tfclient.Workspaces.Delete(ctx, "tf-cloud", TfeWsName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Dleted workpspace")
	return "Workspace Deleted"
}

func DeleteWorkspaceiD(TfeDelWS string, ctx context.Context, Tfclient *tfe.Client) string {

	//Create workspace
	err := Tfclient.Workspaces.DeleteByID(ctx, TfeDelWS)
	if err != nil {
		log.Fatal(err)
	}
	return "Workspace Deleted"
}
