package jsondb

import (
	"github.com/rwrr/oauth2-in-action/oauth"
	"testing"
)

func TestWrite(t *testing.T) {
	server := oauth.Server{
		AuthorizationEndpoint: "Hello",
		TokenEndpoint:         "World",
	}
	Write("oauth", "server", server)
}

func TestRead(t *testing.T) {
	server := oauth.Server{}
	if Read("oauth", "server", &server) {
		println(server.TokenEndpoint)
	}
}

func TestDelete(t *testing.T) {
	Delete("oauth", "server")
}
