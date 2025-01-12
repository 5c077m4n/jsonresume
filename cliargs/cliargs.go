package cliargs

import "flag"

type CliArgs struct {
	FilePath     string
	TemplateName string
	OutputPath   string
	ValidateOnly bool
}

func New() CliArgs {
	filePath := flag.String("file", "./resume.json", "JSON resume file path")
	templateName := flag.String("template", "black-white-min-pro", "JSON resume file path")
	outputPath := flag.String("output-path", "-", "Path to output to")
	validateOnly := flag.Bool("validate-only", false, "Validate Only")
	flag.Parse()

	return CliArgs{
		FilePath:     *filePath,
		TemplateName: *templateName,
		OutputPath:   *outputPath,
		ValidateOnly: *validateOnly,
	}
}
