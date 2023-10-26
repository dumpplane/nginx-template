package main

import (
    "fmt"
    "runtime"

    "github.com/dumpplane/template-controller/internal/configs/nginx"
    "github.com/golang/glog"
)

var version string

func main() {

    commitHash, commitTime, dirtyBuild := getBuildInfo()
    fmt.Printf("Template Controller Version=%v Commit=%v Date=%v DirtyState=%v Arch=%v/%v Go=%v\n", version, commitHash, commitTime, dirtyBuild, runtime.GOOS, runtime.GOARCH, runtime.Version())

    parseFlags()

    templateExecutor := createTemplateExecutors()
 
    fmt.Printf("Template Executor: %v\n", templateExecutor)
}

func createTemplateExecutors() (*nginx.TemplateExecutor) {

    nginxVirtualServerTemplatePath := "nginx.virtualserver.tmpl"
    nginxTransportServerTemplatePath := "nginx.transportserver.tmpl"

    templateExecutor, err := nginx.NewTemplateExecutor(nginxVirtualServerTemplatePath, nginxTransportServerTemplatePath)
    if err != nil {
        glog.Fatalf("Error creating TemplateExecutor: %v", err)
    }

    return templateExecutor
}
