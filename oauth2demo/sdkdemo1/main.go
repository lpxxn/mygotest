package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ory/hydra/firewall"
	"github.com/ory/hydra/sdk"
)

var h *sdk.Client

func main() {
	var err error

	if h, err = sdk.Connect(
		sdk.ClientID("some-consumer6"),
		sdk.ClientSecret("consumer-secret"),
		sdk.ClusterURL("http://localhost:4444"),
	); err != nil {
		log.Fatalf("Could not connect to host: %s", err)
	}

	r := httprouter.New()
	r.GET("/protected", handleProtectedEndpoint)

	listen := fmt.Sprintf("%s:%s", "127.0.0.1", "9989")
	if err := http.ListenAndServe(listen, r); err != nil {
		log.Fatalf("Could not listen on %s becase %s", listen, err)
	}
}

func handleProtectedEndpoint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	token := h.Warden.TokenFromRequest(r)

	clientsData, err := h.Clients.GetClients()
	fmt.Println(err)
	for str, cliVal := range clientsData {
		fmt.Println(str, cliVal)
	}

	// Access control using only access token.
	if status, err := h.Introspection.IntrospectToken(context.Background(), token, "some-scope"); err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	} else {
		log.Printf("Token is allowed to perform action, state lookup gave: %v", status)
	}

	// Access control using access token and access control policies.
	if status, err := h.Warden.TokenAllowed(context.Background(), token, &firewall.TokenAccessRequest{
		Resource: "some:resource-name",
		Action:   "some-action",
	}, "some-scope"); err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	} else {
		log.Printf("Token is allowed to perform action, state lookup gave: %v", status)
	}

	// Access control using access control policies only.
	if err := h.Warden.IsAllowed(context.Background(), &firewall.AccessRequest{
		Resource: "some:resource-name",
		Action:   "some-action",
		Subject:  "some-user",
	}); err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}
}
