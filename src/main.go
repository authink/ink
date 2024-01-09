package main

import (
	"flag"
)

func main() {
	var cmd string
	flag.StringVar(&cmd, "cmd", "run", "Specify a command, such as [run, migrate]")
	flag.Parse()

	switch cmd {
	case "migrate":
		ink := newInk()
		defer ink.Close()

		migrateSchema(ink)

	case "run":
		ink := newInk()
		defer ink.Close()

		setupGracefulShutdown(
			ink,
			createServer(ink),
		)

	default:
		flag.PrintDefaults()
	}
}
