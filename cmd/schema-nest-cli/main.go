package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Masterminds/semver"
	schemanest_files "github.com/timo-reymann/SchemaNest"
	"github.com/timo-reymann/SchemaNest/pkg/buildinfo"
	"github.com/timo-reymann/SchemaNest/pkg/client"
)
import "github.com/urfave/cli/v3"

func main() {
	cli.VersionPrinter = buildinfo.PrintVersionInfo

	app := &cli.Command{
		Name:    "schema-nest-cli",
		Version: buildinfo.Version,
		Usage:   "Interact with the SchemaNest API with ease.",
		Commands: []*cli.Command{
			{
				Name:  "license",
				Usage: "Show the NOTICE for SchemaNest with all relevant license information",
				Action: func(ctx context.Context, command *cli.Command) error {
					fmt.Printf("%s", schemanest_files.Notice)
					return nil
				},
			},
			{
				Name:  "upload-json-schema",
				Usage: "Upload a JSON schema file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "base-url",
						Usage:    "Base URL including protocol for SchemaNest instance",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "api-key",
						Usage:   "API-Key for authentication",
						Sources: cli.NewValueSourceChain(cli.EnvVar("SCHEMA_NEST_CLI_API_KEY")),
					},
					&cli.StringFlag{
						Name:     "version",
						Usage:    "Version of the schema to upload in format major.minor.patch",
						Required: true,
						Validator: func(s string) error {
							if strings.TrimSpace(s) == "" {
								return errors.New("version should be not empty")
							}

							_, err := semver.NewVersion(s)
							if err != nil {
								return fmt.Errorf("invaild version: %s", err)
							}

							return nil
						},
					},
					&cli.StringFlag{
						Name:     "identifier",
						Usage:    "Identifier of the JSON schema",
						Required: true,
						Validator: func(s string) error {
							if strings.TrimSpace(s) == "" {
								return errors.New("identifier should be not empty")
							}

							return nil
						},
					},
					&cli.StringFlag{
						Name:     "path",
						Usage:    "Path to JSON schema file",
						Required: true,
						Validator: func(s string) error {
							stat, err := os.Stat(s)
							if err != nil {
								return fmt.Errorf("cant stat file: %s", err)
							}

							if stat.IsDir() {
								return fmt.Errorf("path parameter must be file")
							}

							return nil
						},
					},
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					version := command.String("version")
					identifier := command.String("identifier")
					path := command.String("path")

					c, err := client.NewClient(command.String("base-url"), command.String("api-key"))
					if err != nil {
						return fmt.Errorf("failed to create API client: %s", err)
					}

					return c.UploadJsonSchema(identifier, version, path)
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		msg := err.Error()
		println(strings.ToUpper(msg[:1]) + strings.ToLower(msg[1:]))
		os.Exit(1)
	}
}
