package main

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/segmentio/go-tableize"
	"github.com/tj/docopt"
)

var version = "0.0.1"
var usage = `
  Usage:
    tableize
    tableize -h | --help
    tableize -v | --version

  Examples:

    $ echo '{"user": { "id": 1 }}' | tableize
    { "user_id": 1 }

  Options:
    -h, --help      show help information
    -v, --version   show version information

`

func main() {
	_, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	for {
		var m map[string]interface{}

		err := dec.Decode(&m)
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}

		err = enc.Encode(tableize.Tableize(&tableize.Input{Value: m}))
		if err != nil {
			log.Fatal(err)
		}
	}
}
