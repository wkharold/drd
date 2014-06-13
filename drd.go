package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/wkharold/dr"
)

var (
	fladdr      = flag.String("addr", ":3080", "registry address")
	fltelemetry = flag.Bool("telemetry", false, "provide service telemetry")
)

func main() {
	flag.Parse()

	dr, err := dr.New(dr.Context{*fltelemetry})
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(*fladdr, dr)
	if err != nil {
		log.Fatal(err)
	}
}
