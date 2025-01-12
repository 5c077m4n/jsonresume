package main

import (
	"embed"
	"encoding/json"
	"html/template"
	"jsonresume/cliargs"
	"jsonresume/schema"
	"os"
	"path/filepath"
	"strings"
)

//go:embed themes/*.html.tmpl
var themes embed.FS

func main() {
	args := cliargs.New()

	content, err := os.ReadFile(args.FilePath)
	if err != nil {
		panic(err)
	}

	resume := &schema.Resume{}
	if err := json.Unmarshal(content, resume); err != nil {
		panic(err)
	}

	if err := resume.Validate(); err != nil {
		panic(err)
	}
	if args.ValidateOnly {
		return
	}

	templateFileName := args.TemplateName + ".html.tmpl"
	tmpl, err := template.New(templateFileName).
		Funcs(template.FuncMap{
			"lower": func(s string) string {
				return strings.ToLower(s)
			},
			"default": func(s1 string, s2 string) string {
				if s1 == "" {
					return s2
				}
				return s2
			},
			"join": func(sign string, s ...string) string {
				return strings.Join(s, sign)
			},
		}).
		ParseFS(
			themes,
			filepath.Join("themes", templateFileName),
		)
	if err != nil {
		panic(err)
	}

	var output *os.File
	if args.OutputPath == "-" {
		output = os.Stdout
	} else {
		output, err = os.Create(args.OutputPath)
		if err != nil {
			panic(err)
		}
		defer output.Close()
	}

	if err := tmpl.Execute(output, resume); err != nil {
		panic(err)
	}
}
