package main

import (
    "flag"
    "fmt"
    "os"
)

var (
    healthStatus = flag.Bool("test-status", false, `a sample bool argument`)
    healthStatusURI = flag.String("test-uri", "/health",`a sample string argument`)
    testPort = flag.Int("listen-port", 9113, "a sample int argument. [1024 - 65535]")

    versionFlag = flag.Bool("version", false, "Print the version, git-commit hash and build date and exit")
)

func parseFlags() {
        
    flag.Parse()

    if *versionFlag {
        os.Exit(0)
    }


}
