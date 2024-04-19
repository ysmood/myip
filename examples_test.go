package myip

import "fmt"

func ExampleNew() {
	mi := New()

	// GetInterfaceIP get the ip of your interface, useful when you want to
	// get your ip inside a private network, such as wifi network.
	fmt.Println(mi.GetInterfaceIP())

	// GetPublicIP get the ip that is public to global.
	fmt.Println(mi.GetPublicIP())
}
