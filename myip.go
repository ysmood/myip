package myip

import (
	"context"
	"net"
)

// NameServer the name server to use for this lib
var NameServer = "ns1.google.com:53"

// GetInterfaceIP get the ip of your interface, useful when you want to
// get your ip inside a private network, such as wifi network.
func GetInterfaceIP() (string, error) {
	conn, err := net.Dial("udp", NameServer)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String(), nil
}

// GetPublicIP get the ip that is public to global.
func GetPublicIP() (string, error) {
	r := net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}
			return d.DialContext(ctx, "udp", NameServer)
		},
	}
	ctx := context.Background()
	txt, err := r.LookupTXT(ctx, "o-o.myaddr.l.google.com")

	return txt[0], err
}
