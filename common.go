package main

import (
	"strings"
)

func filter[T any](items []T, match func(T) bool) []T {
	var filtered []T
	for _, item := range items {
		if match(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func searchSubstring(item string, query string) bool {
	return strings.Contains(strings.ToLower(item), strings.ToLower(query))
}
