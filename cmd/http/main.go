package main

import (
	"log"
	"net/http"

	"strconv"

	"github.com/impactasaurus/soq-api/cmd"
)

func main() {
	networkConfig := cmd.MustSetup()
	if err := http.ListenAndServe(":"+strconv.Itoa(networkConfig.Port), nil); err != nil {
		log.Fatal(err, nil)
	}
}
