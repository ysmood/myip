No dependency, lightweight and fast way to get your ip.

The OpenDNS is not as stable as google's dns server, so here I use the 8.8.8.8.

Example

```go
package main

import (
  "fmt"
  "github.com/ysmood/myip"
)

func main() {
    // GetInterfaceIP get the ip of your interface, useful when you want to
    // get your ip inside a private network, such as wifi network.
    fmt.Println(myip.GetInterfaceIP())

    // GetPublicIP get the ip that is public to global.
    fmt.Println(myip.GetPublicIP())
}
```