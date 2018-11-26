package myip

import (
	"context"
	"net"
	"regexp"
)

// GetInterfaceIP get the ip of your interface, useful when you want to
// get your ip inside a private network, such as wifi network.
func GetInterfaceIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
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
			return d.DialContext(ctx, "udp", "8.8.8.8:53")
		},
	}
	ctx := context.Background()
	txt, err := r.LookupTXT(ctx, "o-o.myaddr.l.google.com")

	return regexp.MustCompile(`\d+\.\d+\.\d+\.\d+`).FindString(txt[1]), err
}
