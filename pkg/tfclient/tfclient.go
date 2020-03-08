package tfclient

import (
	"context"
	"log"

	"github.com/hashicorp/go-tfe"
)

var ctx = context.Background()

func TfeClient() (Tfclient *tfe.Client) {

	//Move token and host to env checks
	config := &tfe.Config{}

	Tfclient, err := tfe.NewClient(config)
	if err != nil {
		log.Print("cant make client")
		log.Print(config)
		log.Fatal(err)
	}
	log.Println("tfe client intailised")

	return
}
