package myip

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInterfaceIP(t *testing.T) {
	ip, err := GetInterfaceIP()

	if err != nil {
		panic(err)
	}

	assert.Regexp(t, `\A(^127\.)|(^10\.)|(^172\.1[6-9]\.)|(^172\.2[0-9]\.)|(^172\.3[0-1]\.)|(^192\.168\.)`, ip)
}

func TestGetInterfaceIPMultipleTimes(t *testing.T) {
	ipA, err := GetInterfaceIP()
	if err != nil {
		panic(err)
	}

	ipB, err := GetInterfaceIP()
	if err != nil {
		panic(err)
	}

	assert.Equal(t, ipA, ipB)
}

func TestGetPublicIP(t *testing.T) {
	var ip string
	var err error
	ip, err = GetPublicIP()

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

	assert.Equal(t, ip, strings.TrimSpace(string(body)))
}

func TestGetPublicIPMultipleTimes(t *testing.T) {
	ipA, err := GetPublicIP()
	if err != nil {
		panic(err)
	}

	ipB, err := GetPublicIP()
	if err != nil {
		panic(err)
	}

	assert.Equal(t, ipA, ipB)
}

func TestDialError(t *testing.T) {
	tmp := NameServer
	NameServer = "a.com"
	_, err := GetInterfaceIP()
	NameServer = tmp
	assert.EqualError(t, err, "dial udp: address a.com: missing port in address")
}
