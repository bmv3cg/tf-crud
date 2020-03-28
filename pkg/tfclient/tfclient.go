package tfclient

import (
	"context"
	"os"

	"github.com/hashicorp/go-tfe"
	"k8s.io/klog"
)

// Ctx for managing TFE client connections
var Ctx = context.Background()

// Tfclient instance
var Tfclient *tfe.Client

func init() {
	Tfclient = TfeClient()
}

// TfeClient singleton instance for creating client
func TfeClient() *tfe.Client {

	if os.Getenv("TFE_TOKEN") == "" {
		klog.Error("TFe token not found in env variable")
	}

	//Move token and host to env checks
	config := &tfe.Config{}

	Tfclient, err := tfe.NewClient(config)
	if err != nil {
		klog.Error("cant make client")
		klog.Info(config)
		klog.Fatal(err)
	}

	klog.V(2).Info("tfe client intailised")

	return Tfclient
}
