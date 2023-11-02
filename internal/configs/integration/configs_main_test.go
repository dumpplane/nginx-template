package configs

import (
    "testing"
    "strings"
    "io/ioutil"
    "github.com/dumpplane/template-controller/internal/configs"
    "github.com/dumpplane/template-controller/internal/configs/nginx"
    "github.com/golang/glog"
)

var (
    nginxPlus               = true
    disableIPV6             = false
    healthStatus            = false
    healthStatusURI         = "/nginx-health"
    nginxStatus             = true
    nginxStatusAllowCIDRs   = "10.1.10.0/24,192.168.1.101"
    nginxStatusPort         = 8080
    enablePrometheusMetrics = false
    enableTLSPassthrough    = false
    enableSnippets          = false
    spireAgentAddress       = ""
    appProtect              = false
    appProtectDos           = false
    enableLatencyMetrics    = false
    enableOIDC              = false
    enableCertManager       = false
    sslRejectHandshake      = false
    allowedCIDRs []string
)

func TestDefault(t *testing.T) {

    nginxConfTemplatePath := "../../../tmpl/nginx-plus.tmpl"
    ngxTemplateExecutor, err := nginx.NewTemplateExecutor(nginxConfTemplatePath)
    if err != nil {
        glog.Fatalf("Error creating TemplateExecutor: %v", err)
    }

    allowedCIDRs, err = parseNginxStatusAllowCIDRs(nginxStatusAllowCIDRs)
    if err != nil {
        glog.Fatalf(`Invalid value: %v`, err)
    }

    cfgParams := configs.NewDefaultConfigParams(nginxPlus)


    staticCfgParams := &configs.StaticConfigParams{
        DisableIPV6:                    disableIPV6,
        HealthStatus:                   healthStatus,
        HealthStatusURI:                healthStatusURI,
        NginxStatus:                    nginxStatus,
        NginxStatusAllowCIDRs:          allowedCIDRs,
        NginxStatusPort:                nginxStatusPort,
        StubStatusOverUnixSocketForOSS: enablePrometheusMetrics,
        TLSPassthrough:                 enableTLSPassthrough,
        EnableSnippets:                 enableSnippets,
        NginxServiceMesh:               spireAgentAddress != "",
        MainAppProtectLoadModule:       appProtect,
        MainAppProtectDosLoadModule:    appProtectDos,
        EnableLatencyMetrics:           enableLatencyMetrics, 
        EnableOIDC:                     enableOIDC,
        SSLRejectHandshake:             sslRejectHandshake,
        EnableCertManager:              enableCertManager,
    }

    ngxConfig := configs.GenerateNginxMainConfig(staticCfgParams, cfgParams)
    content, err := ngxTemplateExecutor.ExecuteMainConfigTemplate(ngxConfig)
    if err != nil { 
        glog.Fatalf("Error generating NGINX main config: %v", err)
    } 

    output := "/tmp/nginx-default.conf"
    ioutil.WriteFile(output, content, 0644)
    t.Log("output to file ", output)
    
}

func parseNginxStatusAllowCIDRs(input string) (cidrs []string, err error) {
        cidrsArray := strings.Split(input, ",")
        for _, cidr := range cidrsArray {
                trimmedCidr := strings.TrimSpace(cidr)
                cidrs = append(cidrs, trimmedCidr)
        }
        return cidrs, nil 
}
