package generator

import (
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/markbates/pkger"
)

func compile(path string, o *Option) error {
	// create destination file
	fp, err := buildPath(path, o)

	if err != nil {
		return err
	}

	fp = strings.TrimSuffix(fp, ".tmpl")

	w, err := os.OpenFile(fp, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer w.Close()

	// open pkger file
	f, err := pkger.Open(path)

	if err != nil {
		return err
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)

	if err != nil {
		return err
	}

	// parse template file
	return template.
		Must(template.New("").Delims("[[", "]]").Parse(string(b))).
		Execute(w, o)
}
