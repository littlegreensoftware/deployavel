package config

import "time"

const (
	// URL to connect to forge with
	URL = "https://forge.laravel.com/api/v1"

	// Timeout default client timeout
	Timeout = 60 * time.Second
)
