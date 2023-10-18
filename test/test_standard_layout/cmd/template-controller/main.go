package main

import (
        "fmt"
        "runtime"

        "github.com/dumpplane/template-controller/internal/nginx"
        "github.com/dumpplane/template-controller/pkg/configuration"
        "github.com/dumpplane/template-controller/api/user"
)

var version string

func main() {

    commitHash, commitTime, dirtyBuild := getBuildInfo()
    fmt.Printf("NGINX Ingress Controller Version=%v Commit=%v Date=%v DirtyState=%v Arch=%v/%v Go=%v\n", version, commitHash, commitTime, dirtyBuild, runtime.GOOS, runtime.GOARCH, runtime.Version())

    nginx.MyFunction()
    configuration.MyFunction()
    user.MyFunction()
}

