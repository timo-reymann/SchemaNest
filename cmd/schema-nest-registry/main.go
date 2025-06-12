package main

import (
	"context"
	"fmt"
	"github.com/timo-reymann/SchemaNest/pkg/api"
	"github.com/timo-reymann/SchemaNest/pkg/buildinfo"
	"github.com/timo-reymann/SchemaNest/pkg/config"
	"github.com/timo-reymann/SchemaNest/pkg/openapi"
	"github.com/timo-reymann/SchemaNest/pkg/persistence/database"
	"github.com/timo-reymann/SchemaNest/pkg/persistence/json_schema"
	"github.com/timo-reymann/SchemaNest/pkg/ui"
	"github.com/urfave/cli/v3"
	"net"
	"net/http"
	"os"
	"strconv"
)

func main() {
	cli.VersionPrinter = buildinfo.PrintVersionInfo

	app := &cli.Command{
		Name:    "schema-nest-registry",
		Version: buildinfo.Version,
		Usage:   "Registry for storing and managing schemas.",
		Commands: []*cli.Command{
			{
				Name:        "serve-http",
				Description: "Start the registry HTTP server",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "port",
						Aliases:     []string{"p"},
						DefaultText: "8080",
					},
					&cli.StringFlag{
						Name:    "config-file",
						Aliases: []string{"C"},
						Validator: func(path string) error {
							if path == "" {
								return nil
							}

							stat, err := os.Stat(path)
							if err != nil {
								return fmt.Errorf("file %s does not exist", path)
							}

							if stat.IsDir() {
								return fmt.Errorf("%s it not a valid file", path)
							}

							return nil
						},
					},
				},
				Action: func(ctx context.Context, command *cli.Command) error {
					cfg, err := config.ParseFromFile(command.String("config-file"), config.ParseFromToml)

					db, err := database.Connect(cfg.DBConnectionString)
					if err != nil {
						return err
					}

					if err = db.MigrateUp(); err != nil {
						return err
					}

					port := command.Int("port")
					if port == 0 {
						port = 8080
					}

					r, err := api.NewServeMux(&api.SchemaNestApiContext{
						JsonSchemaRepository:        &json_schema.JsonSchemaRepositoryImpl{DB: db},
						JsonSchemaVersionRepository: &json_schema.JsonSchemaVersionRepositoryImpl{DB: db},
						Config:                      cfg,
					})
					if err != nil {
						return err
					}
					r.HandleFunc("/", ui.CreateHandler())
					r.HandleFunc("/api-spec.yml", openapi.SpecHandler)

					addr := net.JoinHostPort("0.0.0.0", strconv.Itoa(port))
					fmt.Println("Listening on", addr)
					s := &http.Server{
						Handler: r,
						Addr:    addr,
					}
					return s.ListenAndServe()
				},
			},
		},
	}
	err := app.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
