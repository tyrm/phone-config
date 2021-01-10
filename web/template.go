package web

import (
	"github.com/markbates/pkger"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

func compileTemplates(dir string) (*template.Template, error) {
	tpl := template.New("")

	tpl.Funcs(template.FuncMap{
		"dec": func(i int) int {
			i--
			return i
		},
		"inc": func(i int) int {
			i++
			return i
		},
	})

	err := pkger.Walk(dir, func(path string, info os.FileInfo, _ error) error {
		if info.IsDir() || !strings.HasSuffix(path, ".tmpl") {
			return nil
		}
		f, err := pkger.Open(path)
		if err != nil {
			logger.Errorf("could not open pkger path %s: %s", path, err.Error())
			return err
		}
		// Now read it.
		sl, err := ioutil.ReadAll(f)
		if err != nil {
			logger.Errorf("could not read pkger file %s: %s", path, err.Error())
		}

		// It can now be parsed as a string.
		_, err = tpl.Parse(string(sl))
		if err != nil {
			logger.Errorf("could not open parse template %s: %s", path, err.Error())
			return err
		}

		return nil
	})

	return tpl, err
}