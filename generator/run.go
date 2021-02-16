package generator

import (
	"os"
	"text/tabwriter"

	"github.com/markbates/pkger"
)

// Run a generator
func Run(o *Option) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.Debug)
	defer w.Flush()

	return pkger.Walk("/templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return mkdir(path, o)
		}

		return compile(path, o)
	})
}
