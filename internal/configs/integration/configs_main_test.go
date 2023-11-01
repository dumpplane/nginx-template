package configs

import (
    "testing"
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
    allowedCIDRs            = "0.0.0.0/0"
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
    defaultHTTPPort         = 80
    defaultHTTPSPort        = 443
    tlsPassthroughPort      = 5001
)

func TestDefault(t *testing.T) {

    nginxConfTemplatePath := "../../../tmpl/nginx-plus.tmpl"
    ngxTemplateExecutor, err := nginx.NewTemplateExecutor(nginxConfTemplatePath)
    if err != nil {
        glog.Fatalf("Error creating TemplateExecutor: %v", err)
    }


    cfgParams := configs.NewDefaultConfigParams(*nginxPlus)


    staticCfgParams := &configs.StaticConfigParams{
        DisableIPV6:                    *disableIPV6,
        HealthStatus:                   *healthStatus,
        HealthStatusURI:                *healthStatusURI,
        NginxStatus:                    *nginxStatus,
        NginxStatusAllowCIDRs:          *allowedCIDRs,
        NginxStatusPort:                *nginxStatusPort,
        StubStatusOverUnixSocketForOSS: *enablePrometheusMetrics,
        TLSPassthrough:                 *enableTLSPassthrough,
        EnableSnippets:                 *enableSnippets,
        NginxServiceMesh:               *spireAgentAddress != "",
        MainAppProtectLoadModule:       *appProtect,
        MainAppProtectDosLoadModule:    *appProtectDos,
        EnableLatencyMetrics:           *enableLatencyMetrics, 
        EnableOIDC:                     *enableOIDC,
        SSLRejectHandshake:             *sslRejectHandshake,
        EnableCertManager:              *enableCertManager,
        DefaultHTTPListenerPort:        *defaultHTTPPort,
        DefaultHTTPSListenerPort:       *defaultHTTPSPort,      
        TLSPassthroughPort:             *tlsPassthroughPort,
    }

    ngxConfig := configs.GenerateNginxMainConfig(staticCfgParams, cfgParams)
    content, err := ngxTemplateExecutor.ExecuteMainConfigTemplate(ngxConfig)
    if err != nil { 
        glog.Fatalf("Error generating NGINX main config: %v", err)
    } 
    
}

