package nginx

import (
        "bytes"
        "path"
        "text/template"
)

type TemplateExecutor struct {
        virtualServerTemplate       *template.Template
        transportServerTemplate     *template.Template
}

func NewTemplateExecutor(virtualServerTemplatePath string, transportServerTemplatePath string) (*TemplateExecutor, error) {

    vsTemplate, err := template.New(path.Base(virtualServerTemplatePath)).Funcs(helperFunctions).ParseFiles(virtualServerTemplatePath)
    if err != nil {
        return nil, err
    }

    tsTemplate, err := template.New(path.Base(transportServerTemplatePath)).ParseFiles(transportServerTemplatePath)
    if err != nil {
        return nil, err
    }

    return &TemplateExecutor{
        virtualServerTemplate:       vsTemplate,
        transportServerTemplate:     tsTemplate,
    }, nil
}

func (te *TemplateExecutor) UpdateVirtualServerTemplate(templateString *string) error {
        newTemplate, err := template.New("virtualServerTemplate").Funcs(helperFunctions).Parse(*templateString)
        if err != nil {
                return err
        }
        te.virtualServerTemplate = newTemplate

        return nil
}

func (te *TemplateExecutor) ExecuteVirtualServerTemplate(cfg *VirtualServerConfig) ([]byte, error) {
        var configBuffer bytes.Buffer
        err := te.virtualServerTemplate.Execute(&configBuffer, cfg)

        return configBuffer.Bytes(), err
}

func (te *TemplateExecutor) ExecuteTransportServerTemplate(cfg *TransportServerConfig) ([]byte, error) {
        var configBuffer bytes.Buffer
        err := te.transportServerTemplate.Execute(&configBuffer, cfg)

        return configBuffer.Bytes(), err
}
