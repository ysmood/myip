package myip

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/ysmood/got"
)

type T struct {
	got.G
}

func Test(t *testing.T) {
	got.Each(t, T{})
}

func (t T) GetInterfaceIP() {
	ip, err := New().GetInterfaceIP()

	if err != nil {
		panic(err)
	}

	t.Regex(`\A(^127\.)|(^10\.)|(^172\.1[6-9]\.)|(^172\.2[0-9]\.)|(^172\.3[0-1]\.)|(^192\.168\.)|(^198\.18\.)`, ip)
}

func (t T) GetInterfaceIPMultipleTimes() {
	ipA, err := New().GetInterfaceIP()
	if err != nil {
		panic(err)
	}

	ipB, err := New().GetInterfaceIP()
	if err != nil {
		panic(err)
	}

	t.Eq(ipA, ipB)
}

func (t T) GetPublicIP() {
	var ip string
	var err error
	ip, err = New().GetPublicIP()

	if err != nil {
		panic(err)
	}

	var resp *http.Response
	resp, err = http.Get("https://ipinfo.io/ip")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	t.Eq(ip, strings.TrimSpace(string(body)))
}

func (t T) GetPublicIPMultipleTimes() {
	ipA, err := New().GetPublicIP()
	if err != nil {
		panic(err)
	}

	ipB, err := New().GetPublicIP()
	if err != nil {
		panic(err)
	}

	t.Eq(ipA, ipB)
}

func (t T) DialError() {
	im := New()
	im.NameServer = "a.com"
	_, err := im.GetInterfaceIP()
	t.Eq(err.Error(), "dial udp4: address a.com: missing port in address")
}
