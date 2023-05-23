<!--
Title: RedFish API Go
Description: Redfish api to support any server type that is enabled with Redfish standards
Author: kgrvamsi
Maintainer: vnikolin
-->

# Redfish API

This Library supports both Dell and Hp Servers which have Redfish API enabled.

[![GoDoc](https://godoc.org/github.com/vnikolin/redfishapi?status.svg)](https://godoc.org/github.com/vnikolin/redfishapi)
[![Go Report Card](https://goreportcard.com/badge/github.com/vnikolin/redfishapi)](https://goreportcard.com/report/github.com/vnikolin/redfishapi)

## Usage

```go
//main.go
package main

import "fmt"
import "github.com/vnikolin/redfishapi"

func main() {
    client := redfishapi.NewIloClient("https://hostname-0", "username", "password")
    // use lowercase
    //Dell
    biosData, err := client.GetBiosDataDell()
    if err != nil {
        panic(err)
    }

    fmt.Println(biosData)

    //HP
    client2 := redfishapi.NewIloClient("https://hostname-0", "username", "password")

    fwrData, err := client2.GetFirmwareHP()
    if err != nil {
        panic(err)
    }

    fmt.Println(fwrData)
}
```
