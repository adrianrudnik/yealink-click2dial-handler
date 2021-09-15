package yealink

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func readActionURICommand(device *Device, key string, query url.Values) (string, error) {
	endpoint, err := device.GetCallEndpointUrl()
	if err != nil {
		return "", err
	}

	// Set ConfigMan endpoint
	endpoint.Path = "/servlet"

	// Finalize the query, we NEED to preserve the order of the key command
	// because it will not be executed if it is not in the first position after "?".
	endpoint.RawQuery = fmt.Sprintf("%s&%s", key, query.Encode())

	uri := endpoint.String()

	// Phone client does not like trailing & as query parameters
	uri = strings.TrimRight(uri, "&")

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(device.Username, device.Password)

	resp, err := GetApiClient().Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusUnauthorized {
			return "", ErrCommandFailedUnauthorized
		}

		if resp.StatusCode == http.StatusForbidden {
			return "", ErrCommandFailedBadCredentials
		}

		return "", ErrCommandFailed
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func sendActionURICommand(device *Device, key string, query url.Values) error {
	_, err := readActionURICommand(device, key, query)
	return err
}

func postConfigManApp(device *Device, key string, query url.Values) error {
	endpoint, err := device.GetCallEndpointUrl()
	if err != nil {
		return err
	}

	// Set ConfigMan endpoint
	endpoint.Path = "/cgi-bin/ConfigManApp.com"

	// Compose and merge the query string
	q := endpoint.Query()
	for k, v := range query {
		for _, sv := range v {
			q.Add(k, sv)
		}
	}

	// Finalize query string
	endpoint.RawQuery = q.Encode()

	req, err := http.NewRequest("POST", endpoint.String(), nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(device.Username, device.Password)

	resp, err := GetApiClient().Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	log.Printf("%v+", resp)

	return nil
}

func ReadAccounts(device *Device) ([]Account, error) {
	var list []Account

	q := url.Values{}
	q.Set("accounts", "1")
	q.Set("fw", "0")
	q.Set("dnd", "0")

	resp, err := readActionURICommand(device, "phonecfg=get", q)
	if err != nil {
		return list, err
	}

	scanner := bufio.NewScanner(strings.NewReader(resp))
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "Account") {
			continue
		}

		parts := strings.SplitN(line, "|", 7)

		if parts[5] == "" {
			continue
		}

		list = append(list, Account{
			Label:        parts[1],
			RegisterName: parts[2],
			DisplayName:  parts[3],
			OutgoingURI:  parts[5],
		})
	}

	return list, nil
}

func ReadPhoneConfigParam(device *Device) {
	// phonecfg=get[&accounts=x][&dnd=x][&fw=x]

}

func GetApiClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: time.Second * 5,
	}
}
