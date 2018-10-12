package main

import (
	"log"

	"strconv"

	"github.com/apex/gateway"
	"github.com/impactasaurus/soq-api/cmd"
)

func main() {
	networkConfig := cmd.MustSetup()
	if err := gateway.ListenAndServe(":"+strconv.Itoa(networkConfig.Port), nil); err != nil {
		log.Fatal(err, nil)
	}
}
