package cli

import (
	"fmt"
	"os"

	"github.com/k0kubun/go-ansi"
	"github.com/urfave/cli/v2"
	"github.com/voidint/g/internal/version"
)

func list(*cli.Context) (err error) {
	dirs, err := os.ReadDir(versionsDir)
	if err != nil || len(dirs) <= 0 {
		fmt.Printf("No version installed yet\n\n")
		return nil
	}
	items := make([]*version.Version, 0, len(dirs))
	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}

		v, err := version.New(d.Name())
		if err != nil || v == nil {
			continue
		}
		items = append(items, v)
	}

	render(installed(), items, ansi.NewAnsiStdout())
	return nil
}
