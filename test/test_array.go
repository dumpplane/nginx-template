package main

import (
    "fmt"
)

var (
    allowedCIDRs = [2]string{"10.1.10.0/24", "192.168.1.101"}
)

func main() {

    fmt.Println(allowedCIDRs)
    fmt.Println(&allowedCIDRs)
}
