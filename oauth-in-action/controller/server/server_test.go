package server

import (
	"fmt"
	"testing"
)

func TestGetClient(t *testing.T) {
	client := getClient("oauth-client-1")
	client.ClientSecret = "f**k"
	fmt.Println(client.ClientSecret)
	fmt.Println(clients[0].ClientSecret)
}

func TestToken(t *testing.T) {
	info, ok := approveInfos["h"]
	if ok {
		fmt.Println(info)
	}
}
