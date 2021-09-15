package yealink

import (
	"fmt"
	"net/url"
	"strings"
)

func ConfigureDND(device *Device, toggle bool) error {
	var key string

	if toggle {
		key = "key=DNDOn"
	} else {
		key = "key=DNDOff"
	}

	return sendActionURICommand(device, key, nil)
}

func TriggerAutoProvision(device *Device) error {
	return sendActionURICommand(device, "key=AutoP", nil)
}

func Call(device *Device, number string) error {
	// Prepare correct query parameters
	q := url.Values{}
	q.Set("outgoing_uri", device.DefaultOutgoingURI)

	number = strings.TrimSpace(number)

	// Remove any known scheme handlers
	number = strings.TrimPrefix(number, "tel:")
	number = strings.TrimPrefix(number, "callto:")

	// Remove all whitespace from the phone number before escaping
	number = strings.ReplaceAll(number, " ", "")

	// Ensure query parameters are escaped, this is important for international
	// numbers beginning with "+"
	number= url.QueryEscape(number)

	return sendActionURICommand(device, fmt.Sprintf("key=number=%s", number), q)
}
