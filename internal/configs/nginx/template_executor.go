package nginx

import (
	"bytes"
	"path"
	"text/template"
)

// TemplateExecutor executes NGINX configuration templates.
type TemplateExecutor struct {
	mainTemplate    *template.Template
}

// NewTemplateExecutor creates a TemplateExecutor.
func NewTemplateExecutor(mainTemplatePath string) (*TemplateExecutor, error) {
	// template name must be the base name of the template file https://golang.org/pkg/text/template/#Template.ParseFiles
	nginxTemplate, err := template.New(path.Base(mainTemplatePath)).ParseFiles(mainTemplatePath)
	if err != nil {
		return nil, err
	}

	return &TemplateExecutor{
		mainTemplate:    nginxTemplate,
	}, nil
}

// UpdateMainTemplate updates the main NGINX template.
func (te *TemplateExecutor) UpdateMainTemplate(templateString *string) error {
	newTemplate, err := template.New("nginxTemplate").Parse(*templateString)
	if err != nil {
		return err
	}

	te.mainTemplate = newTemplate

	return nil
}

// ExecuteMainConfigTemplate generates the content of the main NGINX configuration file.
func (te *TemplateExecutor) ExecuteMainConfigTemplate(cfg *MainConfig) ([]byte, error) {
	var configBuffer bytes.Buffer
	err := te.mainTemplate.Execute(&configBuffer, cfg)

	return configBuffer.Bytes(), err
}

