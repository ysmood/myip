package myip

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"testing"
)

func TestGetInterfaceIP(t *testing.T) {
	ip, err := GetInterfaceIP()

	if err != nil {
		panic(err)
	}

	privateIPReg := regexp.MustCompile(
		`(^127\.)|(^10\.)|(^172\.1[6-9]\.)|(^172\.2[0-9]\.)|(^172\.3[0-1]\.)|(^192\.168\.)`,
	)

	if !privateIPReg.MatchString(ip) {
		panic("ip is not private")
	}
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

	if strings.TrimSpace(string(body)) != ip {
		panic("not equal")
	}
}
