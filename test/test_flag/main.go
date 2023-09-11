package main

import (
    "flag"
    "fmt"
)

var (
    mainTemplatePath = flag.String("main-template-path", "", `Path to the main NGINX configuration template. (default for NGINX "nginx.tmpl"; default for NGINX Plus "nginx-plus.tmpl")`)
    name = flag.String("name", "Kylin", `Name`)
    age = flag.Int("age", 0, "age")
    verbose = flag.Bool("verbose", false, "Enable support")
)

func main() {
    // Parse the command-line arguments and flags
    flag.Parse()

    // Access the parsed flag values
    fmt.Printf("mainTemplatePath: %s\n", mainTemplatePath)
    fmt.Printf("Name: %s\n", name)
    fmt.Printf("Age: %d\n", age)
    fmt.Printf("Verbose Mode: %v\n", verbose)

    // Access non-flag command-line arguments (if any)
    fmt.Printf("Non-flag Arguments: %v\n", flag.Args())
}

