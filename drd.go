package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/wkharold/dr"
)

var (
	fladdr      = flag.String("addr", ":3080", "registry address")
	fltelemetry = flag.Bool("telemetry", false, "provide service telemetry")
	fllogfile   = flag.String("log", "", "access log")
)

func main() {
	flag.Parse()

	ctx := dr.Context{Telemetry: *fltelemetry}

	if len(*fllogfile) > 0 {
		switch *fllogfile {
		case "stdout":
			ctx.AccessOut = os.Stdout
		default:
			lf, err := os.Open(*fllogfile)
			if err != nil {
				log.Fatal(err)
			}
			ctx.AccessOut = lf
		}
	}

	dr, err := dr.New(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(*fladdr, dr)
	if err != nil {
		log.Fatal(err)
	}
}
