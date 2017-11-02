package main

import (
	"fmt"
	"github.com/ory-am/hydra/sdk"
)

func main() {

	var hydra, err = sdk.Connect(
		sdk.ClientID("some-consumer9"),
		sdk.ClientSecret("consumer-secret"),
		sdk.ClusterURL("http://localhost:4444"),
		sdk.Scopes("hydra.clients"),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	clients, err := hydra.Clients.GetClients()

	if err != nil {
		fmt.Println(err)
		return
	}

	for c := range clients {
		fmt.Println(c)
	}
}
