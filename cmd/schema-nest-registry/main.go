package main

import (
	"fmt"
	"github.com/timo-reymann/SchemaNest/pkg/api"
	"github.com/timo-reymann/SchemaNest/pkg/buildinfo"
	"github.com/timo-reymann/SchemaNest/pkg/persistence/database"
	"github.com/urfave/cli/v2"
	"net"
	"net/http"
	"os"
	"strconv"
)

func main() {
	cli.VersionPrinter = buildinfo.PrintVersionInfo

	app := &cli.App{
		Name:     "schema-nest-registry",
		Version:  buildinfo.Version,
		Compiled: buildinfo.BuildTimeParsed,
		Usage:    "Registry for storing and managing schemas.",
		Commands: []*cli.Command{
			{
				Name:        "serve-http",
				Description: "Start the registry HTTP server",
				Flags: []cli.Flag{
					&cli.IntFlag{Name: "port", Aliases: []string{"p"}, DefaultText: "8080"},
				},
				Action: func(context *cli.Context) error {
					db, err := database.Connect("sqlite3://schema_nest.sqlite?cache=shared&mode=memory")
					if err != nil {
						return err
					}

					if err = db.MigrateUp(); err != nil {
						return err
					}

					port := context.Int("port")
					if port == 0 {
						port = 8080
					}

					r, err := api.NewServeMux()
					if err != nil {
						return err
					}
					r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
						_, _ = writer.Write([]byte(`Welcome to SchemaNest! Here you will soon see a fancy dancy UI!`))
					})

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
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
