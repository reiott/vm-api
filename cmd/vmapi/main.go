package main

import (
	"context"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/reiott/vm-api/http"
)

func main() {
	srv := http.Server{}

	parser := flags.NewParser(&srv, flags.Default)
	parser.ShortDescription = `VM APIs`
	parser.LongDescription = `Options for VM APIs`

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	srv.Serve(context.Background())
}
