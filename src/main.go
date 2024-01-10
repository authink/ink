package main

import (
	"log"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/migrate"
	"github.com/spf13/cobra"
)

var cmdInk = &cobra.Command{Use: "ink"}

var cmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate schema up or down",
	Run: func(cmd *cobra.Command, args []string) {
		direction, err := cmd.Flags().GetString("direction")

		if err != nil {
			log.Fatalf("CMD migrate: %s\n", err)
		}

		ink := core.NewInk()
		defer ink.Close()

		migrate.Schema(ink, direction)
	},
}

var cmdSeed = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database",
	Run: func(cmd *cobra.Command, args []string) {
		ink := core.NewInk()
		defer ink.Close()

		if err := migrate.Seed(ink); err != nil {
			log.Fatalf("Seed: %s\n", err)
		}
	},
}

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run ink server",
	Run: func(cmd *cobra.Command, args []string) {
		ink := core.NewInk()
		defer ink.Close()

		setupGracefulShutdown(
			ink,
			createServer(ink),
		)
	},
}

func init() {
	cmdMigrate.Flags().StringP("direction", "d", "up", "Specify migrate direction[up, down]")

	cmdInk.AddCommand(cmdMigrate)
	cmdInk.AddCommand(cmdRun)
	cmdInk.AddCommand(cmdSeed)
}

func main() {
	if err := cmdInk.Execute(); err != nil {
		log.Fatalf("main: %s\n", err)
	}
}
