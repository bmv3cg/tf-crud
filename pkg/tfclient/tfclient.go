package tfclient

import (
	"context"
	"log"

	"github.com/hashicorp/go-tfe"
)

var Ctx = context.Background()

var Tfclient *tfe.Client

func init() {
	Tfclient = TfeClient()
}

func TfeClient() *tfe.Client {

	//Move token and host to env checks
	config := &tfe.Config{}

	Tfclient, err := tfe.NewClient(config)
	if err != nil {
		log.Print("cant make client")
		log.Print(config)
		log.Fatal(err)
	}
	//commenting for now, will be moving to debug log level
	//log.Println("tfe client intailised")

	return Tfclient
}
