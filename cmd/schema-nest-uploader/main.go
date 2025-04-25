package main

import (
	"github.com/timo-reymann/SchemaNest/pkg/buildinfo"
	"os"
)
import "github.com/urfave/cli/v2"

func main() {
	cli.VersionPrinter = buildinfo.PrintVersionInfo

	app := &cli.App{
		Name:     "schema-nest-uploader",
		Version:  buildinfo.Version,
		Compiled: buildinfo.BuildTimeParsed,
		Usage:    "Upload schemas to the SchemaNest",
	}
	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
