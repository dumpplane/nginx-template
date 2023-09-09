package main

import (
   "fmt"
   "runtime"
   "github.com/golang/glog"

   "github.com/dumpplane/nginx-template/core/configs/nginx"
   "github.com/dumpplane/nginx-template/core/configs/gateway"
)

var version = "1.0"

func main() {
    fmt.Println("main start")

    commitHash, commitTime, dirtyBuild := getBuildInfo()
    fmt.Printf("Version=%v Commit=%v Date=%v DirtyState=%v Arch=%v/%v Go=%v\n", version, commitHash, commitTime, dirtyBuild, runtime.GOOS, runtime.GOARCH, runtime.Version())

    parseFlags()


    //templateExecutor, templateExecutorV2 := createTemplateExecutors()

    //fmt.Println(templateExecutor)
    //fmt.Println(templateExecutorV2)
}

func createTemplateExecutors() (*nginx.TemplateExecutor, *gateway.TemplateExecutor) {

        nginxConfTemplatePath := "nginx.tmpl"
        nginxIngressTemplatePath := "nginx.ingress.tmpl"
        nginxVirtualServerTemplatePath := "nginx.virtualserver.tmpl"
        nginxTransportServerTemplatePath := "nginx.transportserver.tmpl"

        templateExecutor, err := nginx.NewTemplateExecutor(nginxConfTemplatePath, nginxIngressTemplatePath)
        if err != nil {
                glog.Fatalf("Error creating TemplateExecutor: %v", err)
        }       
                
        templateExecutorV2, err := gateway.NewTemplateExecutor(nginxVirtualServerTemplatePath, nginxTransportServerTemplatePath)
        if err != nil {
                glog.Fatalf("Error creating TemplateExecutorV2: %v", err)
        }
        
        return templateExecutor, templateExecutorV2

}

func getBuildInfo() (commitHash string, commitTime string, dirtyBuild string) {
        commitHash = "25b2116"
        commitTime = "2023-08-01"
        dirtyBuild = "bee"
        return commitHash, commitTime, dirtyBuild
}

func parseFlags() {

}
