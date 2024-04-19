package myip

import (
	"context"
	"errors"
	"net"
)

type MyIP struct {
	Protocol   string
	NameServer string
	IPServer   string
}

func New() *MyIP {
	return &MyIP{
		Protocol:   "udp4",
		NameServer: "ns1.google.com:53",
		IPServer:   "o-o.myaddr.l.google.com",
	}
}

// GetInterfaceIP get the ip of your interface, useful when you want to
// get your ip inside a private network, such as wifi network.
func (mi *MyIP) GetInterfaceIP() (string, error) {
	conn, err := net.Dial(mi.Protocol, mi.NameServer)
	if err != nil {
		return "", err
	}

	defer func() { _ = conn.Close() }()

	localAddr := conn.LocalAddr().(*net.UDPAddr) //nolint: forcetypeassert

	return localAddr.IP.String(), nil
}

var errNoPublicIP = errors.New("[myip] can't get a ip")

// GetPublicIP get the ip that is public to global.
func (mi *MyIP) GetPublicIP() (string, error) {
	r := net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{}

			return d.DialContext(ctx, mi.Protocol, mi.NameServer)
		},
	}

	txt, err := r.LookupTXT(context.Background(), mi.IPServer)
	if err != nil {
		return "", err
	}

	if len(txt) == 0 {
		return "", errNoPublicIP
	}

	return txt[0], nil
}
