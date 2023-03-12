package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/LockBlock-dev/MinePot/typings"
)

func SendWebhook(config *typings.Config, ip string, reported bool, didHandshake bool, didPing bool) error {
    // Parse the color string
    var r, g, b int
    fmt.Sscanf(config.WebhookEmbedColor, "#%02x%02x%02x", &r, &g, &b)
    // Combine the red, green, and blue color components into a single int value
    color := (r << 16) | (g << 8) | b

    var hostnamePart string
    hostname, err := os.Hostname()
	if err != nil {
		hostnamePart = ""
	} else {
        hostnamePart = " of " + hostname
    }
    
    // Create a DiscordWebhookPayload struct with the message
    payload := typings.DiscordWebhookPayload{
        Embeds: []typings.DiscordWebhookEmbed{
            {
                Title: "New scan detected on port " + fmt.Sprint(config.Port) + hostnamePart,
                Color: color,
                Fields: []typings.DiscordWebhookField{
                    {
                        Name: "IP",
                        Value: "`" + ip + "`",
                    },
                    {
                        Name: "Reported to AbuseIPDB",
                        Value: fmt.Sprintf("`%t`", reported),
                    },
                    {
                        Name: "Handshake",
                        Value: fmt.Sprintf("`%t`", didHandshake),
                        Inline: true,
                    },
                    {
                        Name: "Ping",
                        Value: fmt.Sprintf("`%t`", didPing),
                        Inline: true,
                    },
                },
            },
        },
    }

    // Marshal the struct to JSON
    payloadJSON, err := json.Marshal(payload)
    if err != nil {
        return fmt.Errorf("error encoding JSON payload: %w", err)
    }

    // Create a new HTTP POST request with the payload as the body
    req, err := http.NewRequest("POST", config.WebhookUrl, bytes.NewBuffer(payloadJSON))
    if err != nil {
        return fmt.Errorf("error creating HTTP request: %w", err)
    }

    // Set the content-type header to application/json
    req.Header.Set("Content-Type", "application/json")

    // Send the request
    client := http.DefaultClient
    resp, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("error making HTTP request: %w", err)
    }
    defer resp.Body.Close()

    return nil
}