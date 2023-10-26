package main

import (
    "flag"
    "fmt"
    "os"
    "regexp"

    "github.com/golang/glog"
)

var (
    healthStatus = flag.Bool("test-status", false, `a sample bool argument`)
    healthStatusURI = flag.String("test-uri", "/health",`a sample string argument`)
    testPort = flag.Int("listen-port", 9113, "a sample int argument. [1024 - 65535]")

    includeYearInLogs = flag.Bool("include-year", false, "Option to include the year in the log header")

    versionFlag = flag.Bool("version", false, "Print the version, git-commit hash and build date and exit")
)

func parseFlags() {
        
    flag.Parse()

    if *versionFlag {
        os.Exit(0)
    }

    initialChecks()

    validationChecks()
}

func validationChecks() {

    healthStatusURIValidationError := validateLocation(*healthStatusURI)
    if healthStatusURIValidationError != nil {
        glog.Fatalf("Invalid value for health-status-uri: %v", healthStatusURIValidationError)
    }
}

func initialChecks() {
    err := flag.Lookup("logtostderr").Value.Set("true")
    if err != nil {
        glog.Fatalf("Error setting logtostderr to true: %v", err)
    }

    glog.Infof("Starting with flags: %+q", os.Args[1:])
        
    unparsed := flag.Args()
    if len(unparsed) > 0 {
        glog.Warningf("Ignoring unhandled arguments: %+q", unparsed)
    }
} 

const ( 
    locationFmt    = `/[^\s{};]*`
    locationErrMsg = "must start with / and must not include any whitespace character, `{`, `}` or `;`"
)

var locationRegexp = regexp.MustCompile("^" + locationFmt + "$")

func validateLocation(location string) error {
    if location == "" || location == "/" {
        return fmt.Errorf("invalid location format: '%v' is an invalid location", location)
    }
    return nil      
} 
