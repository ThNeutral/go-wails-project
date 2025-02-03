package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Endpoint string

const BASE_URL Endpoint = "https://mhw-db.com/"

func fetchData[T any](endpoint Endpoint) (T, error) {
	var data T

	response, err := http.Get(string(BASE_URL + endpoint))
	if err != nil {
		return data, err
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return data, err
	}

	json.Unmarshal(bytes, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
