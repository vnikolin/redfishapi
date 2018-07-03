# Redfish API

This Library should support both Dell and Hp Servers which have Redfish API enabled.


[![Go Report Card](https://goreportcard.com/badge/github.com/kgrvamsi/redfishapi)](https://goreportcard.com/report/github.com/kgrvamsi/redfishapi)

## Usage

```go
//main.go
package main

import "fmt"
import "github.com/kgrvamsi/redfishapi"

func main() {
    client := redfishapi.NewIloClient("hostname-0", "username", "password")
    // use lowercase
    mac_address, err := client.GetMacAddress("dell")
    if err != nil {
        panic(err)
    }

    fmt.Println(mac_address)
}
```
