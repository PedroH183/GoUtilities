package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HttpClient struct {
	client *http.Client
}

func (c *HttpClient) makeGetRequest(
	url string, headers map[string]string, target interface{}) error {

	requestBuilder, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return err
	}
	for key, value := range headers {
		requestBuilder.Header.Set(key, value)
	}
	response, err := c.client.Do(requestBuilder)

	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return json.NewDecoder(response.Body).Decode(target)
	}
	body, _ := io.ReadAll(response.Body)
	return fmt.Errorf("HTTP %d: %s", response.StatusCode, string(body))
}

func (c *HttpClient) makePostRequest(
	url string, headers map[string]string, target interface{}, payload interface{}) error {

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	requestBuilder, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))

	if err != nil {
		return err
	}
	for key, value := range headers {
		requestBuilder.Header.Set(key, value)
	}
	response, err := c.client.Do(requestBuilder)

	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return json.NewDecoder(response.Body).Decode(target)
	}
	body, _ := io.ReadAll(response.Body)
	return fmt.Errorf("HTTP %d: %s", response.StatusCode, string(body))
}
