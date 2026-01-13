package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"io"
)

func fetchAndPrintLocations(cfg *config, url string) error {
	raw, err := fetchLocationsAreas(cfg, url)
	if err != nil {
		return err
	}

	return decodeAndPrintLocations(cfg, raw)
}

func fetchLocationsAreas(cfg *config, url string) ([]byte, error) {
	// Check Cache first
	if cachedData, ok := cfg.Cache.Get(url); ok {
		return cachedData, nil
	}
	
	// Fetch from API
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//Store in cache
	cfg.Cache.Add(url, body)

	return body, nil
	
}

func decodeAndPrintLocations(cfg *config, raw []byte) error {
	var res locationAreaResponse

	if err := json.Unmarshal(raw, &res); err != nil {
		return err
	}

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}

	if res.Next != nil {
		cfg.Next = *res.Next
	} else {
		cfg.Next = ""
	}

	if res.Previous != nil {
		cfg.Previous = *res.Previous
	} else {
		cfg.Previous = ""
	}

	return nil
}

func getFromCacheOrFetch(cfg *config, url string) ([]byte, error) {
	if data, ok := cfg.Cache.Get(url); ok {
		return data, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	cfg.Cache.Add(url, data)

	return data, nil
}