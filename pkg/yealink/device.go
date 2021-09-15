package yealink

import (
	"fmt"
	"net"
	"net/url"
)

type Device struct {
	IP                     net.IP `yaml:"ip"`
	IsHTTPS                bool `yaml:"https"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DefaultOutgoingURI string `yaml:"default_outgoing_uri"`
}

func NewDevice(ip, username, password string) (*Device, error) {
	d := &Device{
		IsHTTPS:                false,
		Username:               username,
		Password:               password,
	}

	err := d.SetIP(ip)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Device) SetIP(in string) error {
	ip := net.ParseIP(in)

	if ip == nil {
		return ErrCouldNotParseIpString
	}

	d.IP = ip

	return nil
}

func (d *Device) GetRootUrl(path string) (*url.URL, error) {
	out, err := url.Parse(fmt.Sprintf("http://10.3.20.10%s", path))

	if err != nil {
		return nil, err
	}

	if d.IsHTTPS {
		out.Scheme = "https"
	}

	out.Host = d.IP.String()

	return out, nil
}

func (d *Device) GetCallEndpointUrl() (*url.URL, error) {
	// Should fit for all yealink devices?
	// @see https://support.yealink.com/en/portal/knowledge/show?id=f8994ddaabfd7dbd59576b17
	// Endpoint example http://10.3.20.10/servlet?key=number=1234&outgoing_uri=1006@10.2.1.48
	out, err := d.GetRootUrl("/servlet")
	return out, err
}
