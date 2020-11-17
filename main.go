package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/andig/evcc/util"
	"github.com/andig/evcc/vehicle"
)

func usage() {
	fmt.Println(`
evcc-soc

Usage:
  evcc-soc vehicle [--log level] [--param value [...]]
`)
}

func main() {
	if len(os.Args) < 3 {
		usage()
		log.Fatal("not enough arguments")
	}

	typ := strings.ToLower(os.Args[1])
	args := make(map[string]interface{})

	var key string
	for _, arg := range os.Args[2:] {
		if key == "" {
			key = strings.ToLower(strings.TrimLeft(arg, "-"))
		} else {
			if key == "log" {
				util.LogLevel(arg, nil)
			} else {
				args[key] = arg
			}
			key = ""
		}
	}

	if key != "" {
		usage()
		log.Fatal("unexpected number of parameters")
	}

	v, err := vehicle.NewFromConfig(typ, args)
	if err != nil {
		log.Fatal(err)
	}

	soc, err := v.ChargeState()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(soc)
}
