package api

import (
	"compress/gzip"
	"encoding/json"
	"os"
)

// ApiConfiguration contains all information to connect with the api.
type ApiConfiguration struct {
	BaseUrl     string  `json:"baseurl"`
	AccessToken string  `json:"accesstoken"`
	AppId       string  `json:"appid"`
}

// ReadApiConfiguration reads the api configuration from a file.
func ReadApiConfiguration(filename string) (*ApiConfiguration, error) {
	var configuration ApiConfiguration
	if err := readConfig(filename, &configuration); err != nil {
		return nil, err
	}
	return &configuration, nil
}

// readConfig reads a configuration file: gzipped json file.
func readConfig(filename string, target interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(gzipReader)
	if err := decoder.Decode(target); err != nil {
		return err
	}

	return nil
}

// WriteApiConfiguration writes a configuration file: gzipped json file.
func WriteApiConfiguration(filename string, config *ApiConfiguration) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	writer := gzip.NewWriter(file)
	defer writer.Close()

	encoder := json.NewEncoder(writer)
	if err := encoder.Encode(config); err != nil {
		return err
	}

	return nil
}
