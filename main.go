package main

import (
	"github.com/adrianrudnik/yealink-click2dial-handler/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		TimeFormat: time.RFC3339,
		NoColor: true,
		Out: os.Stderr,
	})

	cmd.Execute()
}
