package internal

import (
	"encoding/json"
	"github.com/adrianrudnik/yealink-url-scheme-handler/service/pkg/yealink"
	"github.com/kirsle/configdir"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetConfigFilePath() (string, error) {
	p := configdir.LocalConfig("yealink-url-scheme-handler")

	// Ensure the config path exists

	err := configdir.MakePath(p) // Ensure it exists.
	if err != nil {
		return "", err
	}

	file := filepath.Join(p, "config.json")

	return file, nil
}

func StoreConfig(device yealink.Device) error {
	file, err := GetConfigFilePath()
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(device, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(file, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func LoadConfig() (yealink.Device, error) {
	device := yealink.Device{}

	file, err := GetConfigFilePath()
	if err != nil {
		return device, err
	}

	if _, err := os.Stat(file); err != nil {
		return device, ErrNotConfigured
	}

	b, err := ioutil.ReadFile(file)
	if err != nil {
		return device, err
	}

	err = json.Unmarshal(b, &device)

	return device, err
}

func DeleteConfig() error {
	file, err := GetConfigFilePath()
	if err != nil {
		return err
	}

	if _, err := os.Stat(file); err == nil {
		return os.Remove(file)
	} else {
		return nil
	}
}
