package main

import (
    "fmt"
    "runtime"

    "github.com/dumpplane/template-controller/internal/configs"
    "github.com/dumpplane/template-controller/internal/configs/nginx"
    "github.com/dumpplane/template-controller/internal/configs/nginx/gateway"
    "github.com/golang/glog"
)

var version string

func main() {

    commitHash, commitTime, dirtyBuild := getBuildInfo()
    fmt.Printf("Template Controller Version=%v Commit=%v Date=%v DirtyState=%v Arch=%v/%v Go=%v\n", version, commitHash, commitTime, dirtyBuild, runtime.GOOS, runtime.GOARCH, runtime.Version())

    parseFlags()

    ngxTemplateExecutor, templateExecutor := createTemplateExecutors()

    nginxPlus := true
    cfgParams := configs.NewDefaultConfigParams(nginxPlus)
 
    fmt.Printf("Template Executor: %v, %v, %s\n", ngxTemplateExecutor, templateExecutor, cfgParams.UpstreamZoneSize)
}

func createTemplateExecutors() (*nginx.TemplateExecutor, *gateway.TemplateExecutor) {

    nginxConfTemplatePath := "tmpl/nginx-plus.tmpl"
    nginxVirtualServerTemplatePath := "tmpl/nginx-plus.virtualserver.tmpl"
    nginxTransportServerTemplatePath := "tmpl/nginx-plus.transportserver.tmpl"

    ngxTemplateExecutor, err := nginx.NewTemplateExecutor(nginxConfTemplatePath)
    if err != nil {
        glog.Fatalf("Error creating TemplateExecutor: %v", err)
    }

    templateExecutor, err := gateway.NewTemplateExecutor(nginxVirtualServerTemplatePath, nginxTransportServerTemplatePath)
    if err != nil {
        glog.Fatalf("Error creating TemplateExecutor: %v", err)
    }

    return ngxTemplateExecutor, templateExecutor
}
