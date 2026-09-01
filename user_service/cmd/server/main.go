package main

import (
	"log"

	"github.com/teper-ya-pomenyal/privy_stream/jwtmanager"
)

func main() {
	_, err := jwtmanager.LoadPrivateKey("keys/private.pem") // #TODO - заменить _ на privateKey
	if err != nil {
		log.Fatal(err)
	}
}
