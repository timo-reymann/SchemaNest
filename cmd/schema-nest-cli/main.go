package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/semver"
	"github.com/timo-reymann/SchemaNest/pkg/buildinfo"
	"github.com/timo-reymann/SchemaNest/pkg/client"
	"os"
	"strings"
)
import "github.com/urfave/cli/v3"

func main() {
	cli.VersionPrinter = buildinfo.PrintVersionInfo

	app := &cli.Command{
		Name:    "schema-nest-cli",
		Version: buildinfo.Version,
		Usage:   "Upload schemas to the SchemaNest",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "base-url",
				Usage:    "Base URL including protocol for SchemaNest instance",
				Required: true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "upload-json-schema",
				Usage: "Upload a JSON schema file",
				Flags: []cli.Flag{
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

					c, err := client.NewClient(command.String("base-url"))
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
