package main

import (
	"log"
	"os"

	"github.com/authink/ink.go/src/core"
	"github.com/authink/ink.go/src/migrate"
	"github.com/cosmtrek/air/runner"
	"github.com/spf13/cobra"
	"github.com/swaggo/swag"
	"github.com/swaggo/swag/format"
	"github.com/swaggo/swag/gen"
)

var cmdInk = &cobra.Command{Use: "ink"}

var cmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate schema up or down",
	Run: func(cmd *cobra.Command, args []string) {
		direction, err := cmd.Flags().GetString("direction")

		if err != nil {
			panic(err)
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

		migrate.Seed(ink)
	},
}

var cmdSwag = &cobra.Command{
	Use:   "swag",
	Short: "Generate swagger docs",
	Run: func(cmd *cobra.Command, args []string) {
		err := format.New().Build(&format.Config{
			SearchDir: "./src",
			MainFile:  "router/setup.go",
		})
		if err != nil {
			panic(err)
		}

		err = gen.New().Build(&gen.Config{
			SearchDir:   "./src",
			MainAPIFile: "router/setup.go",
			OutputDir:   "./src/docs",

			PropNamingStrategy: swag.CamelCase,
			OutputTypes:        []string{"go", "json", "yaml"},

			ParseDepth: 100,

			OverridesFile: gen.DefaultOverridesFile,
			ParseGoList:   true,

			LeftTemplateDelim:  "{{",
			RightTemplateDelim: "}}",

			Debugger:         log.New(os.Stdout, "", log.LstdFlags),
			CollectionFormat: swag.TransToValidCollectionFormat("csv"),
		})
		if err != nil {
			panic(err)
		}
	},
}

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run ink server",
	Run: func(cmd *cobra.Command, args []string) {
		hotReload, err := cmd.Flags().GetBool("live-reload")

		if err != nil {
			panic(err)
		}

		if hotReload {
			core.AssertEnvDev("live-reload")

			cfg, err := runner.InitConfig(".air.toml")
			if err != nil {
				panic(err)
			}

			r, err := runner.NewEngineWithConfig(cfg, true)
			if err != nil {
				panic(err)
			}

			r.Run()
		} else {
			ink := core.NewInk()
			defer ink.Close()

			setupGracefulShutdown(
				ink,
				createServer(ink),
			)
		}
	},
}

func init() {
	cmdMigrate.Flags().StringP("direction", "d", "up", "Specify migrate direction[up, down]")

	cmdRun.Flags().BoolP("live-reload", "l", false, "Enable live reload")

	cmdInk.AddCommand(cmdMigrate)
	cmdInk.AddCommand(cmdSeed)
	cmdInk.AddCommand(cmdSwag)
	cmdInk.AddCommand(cmdRun)
}

func main() {
	if err := cmdInk.Execute(); err != nil {
		panic(err)
	}
}
