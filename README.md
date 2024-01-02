# Api client for yandex webmaster api

## Install

    go get -u github.com/bzdvdn/yandex-webmaster-go

## Usage

=======

```go
package main

import (
	"encoding/json"
	"fmt"

	yandexwebmaster "github.com/bzdvdn/yandex-webmaster-go/yandex_webmaster"
)

func main() {
	client, err := yandexwebmaster.NewClient("you_token")
	if err != nil {
		fmt.Printf("error - %v", err.Error())
        return
	}
	r, err := client.Hosts.GetHosts()
	if err != nil {
		fmt.Printf("error - %v", err.Error())
		return
	}
	for _, host := range r.Hosts {
		fmt.Println(host)
	}
}
```
