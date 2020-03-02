package main

import (
	"context"
	"flag"
	"os"
	"tf-crud/pkg/tfclient"
	"tf-crud/pkg/tfcrud"

	tfe "github.com/hashicorp/go-tfe"
)

var Tfclient *tfe.Client

func init() {
	Tfclient = tfclient.TfeClient()
}

func main() {

	//
	// Add time stamp
	// List all workspaces
	// List all teams
	// move to klog with debuggging

	ctx := context.Background()

	tfe_workspace_name := flag.String("tfe_workspace_create", "foo", "name of workspace to be created")

	flag.Parse()

	if tfe_workspace_name == nil {
		flag.Usage()
		os.Exit(1)
	}

	tfcrud.CreateWorkspace(*tfe_workspace_name, ctx, Tfclient)
}
