package main

import (
	"github.com/kenmoini/seesaw-sensor/pkg/http"
	"github.com/kenmoini/seesaw-sensor/pkg/ping"
)

func main() {
	// Ping Test
	ping.Ping("google.com")

	// HTTP Test
	http.HTTPRequest("GET", "https://google.com", "", map[string]string{}, 200)
}
