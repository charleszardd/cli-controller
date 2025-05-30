package io

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func ReadJSONFile[T any](filePath string) (*T, error) {
	file, err  :=  os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("cannot open file %s: %w", filePath, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var data T
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("invalid JSON in file %s: %w", filePath, err)
		
	}

	return &data, nil
}

func BuildRequest(method, url string, headers map[string]string, bodyData any) (*http.Request, error) {
	var body io.Reader

	if bodyData != nil {
		jsonBody, err := json.Marshal(bodyData)
		if err != nil {
			return nil, fmt.Errorf("failed to encode request body: %w", err)
		}
		body = bytes.NewBuffer(jsonBody)
		// fmt.Printf("Sending %s request to %s with headers: %v and body: %v\n", method, url, headers, body)

	}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if bodyData != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	for key, val := range headers {
		request.Header.Set(key, val)
	}

	return  request, nil
}

func SendRequest(request *http.Request) ([]byte, error) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP %d: %s", response.StatusCode, http.StatusText(response.StatusCode))
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	return body, nil
	
}

func DecodeJSON[T any] (data []byte) (*T, error) {
	var result T
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return &result, nil
}

func DoJSONRequest[T any] (method, url string, headers map[string]string, bodyData any) (*T, error) {
	request, err := BuildRequest(method, url, headers, bodyData)
	if err != nil {
		return nil, err
	}

	body, err := SendRequest(request)
	if err != nil {
		return nil, err
	}
	return DecodeJSON[T](body)
}
