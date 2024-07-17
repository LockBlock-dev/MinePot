package util

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func Report(ip string, key string, port int) (int, error) {
	URI := "https://api.abuseipdb.com/api/v2/report"
	t := time.Now()
	var hostnamePart string
	hostname, err := os.Hostname()
	if err != nil {
		hostnamePart = ""
	} else {
		hostnamePart = " of " + hostname
	}

	payload := fmt.Sprintf(
		"ip=%s&categories=%s&comment=%s",
		url.QueryEscape(ip),
		url.QueryEscape("14"),
		url.QueryEscape(fmt.Sprintf(
			"%s: Minecraft server scan detected from %s on port %d%s",
			t.Format("2006-01-02 15:04:05"),
			ip,
			port,
			hostnamePart,
		)),
	)

	req, err := http.NewRequest("POST", URI, strings.NewReader(payload))
	if err != nil {
		return -1, fmt.Errorf("error creating HTTP request: %w", err)
	}

	req.Header.Set("Key", key)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return -1, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
