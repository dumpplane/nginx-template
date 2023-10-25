package main

import (
        "fmt"
        "runtime"
)

var version string

func main() {

    commitHash, commitTime, dirtyBuild := getBuildInfo()
    fmt.Printf("Template Controller Version=%v Commit=%v Date=%v DirtyState=%v Arch=%v/%v Go=%v\n", version, commitHash, commitTime, dirtyBuild, runtime.GOOS, runtime.GOARCH, runtime.Version())

}

