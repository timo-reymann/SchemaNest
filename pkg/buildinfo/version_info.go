package buildinfo

import (
	"fmt"
	"github.com/urfave/cli/v3"
	"os"
	"runtime"
	"text/tabwriter"
	"time"
)

func addLine(w *tabwriter.Writer, heading string, val string) {
	_, _ = fmt.Fprintf(w, heading+"\t%s\n", val)
}

func PrintVersionInfo(cmd *cli.Command) {
	fmt.Println(cmd.Name, cmd.Version)
	println()
	println("Build information")
	w := tabwriter.NewWriter(os.Stderr, 10, 1, 10, byte(' '), tabwriter.TabIndent)
	addLine(w, "GitSha", GitSha)
	addLine(w, "Version", Version)
	addLine(w, "BuildTime", BuildTimeParsed.Format(time.DateTime))
	addLine(w, "Go-Version", runtime.Version())
	addLine(w, "OS/Arch", runtime.GOOS+"/"+runtime.GOARCH)
	_ = w.Flush()
}
