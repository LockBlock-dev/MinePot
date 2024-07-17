package util

import (
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"time"

	"github.com/LockBlock-dev/MinePot/types"
)

func HandleReport(conn types.ConnWrapper, addr string) {
	if conn.PacketsReceived >= conn.Config.ReportThreshold {
		host, _, err := net.SplitHostPort(addr)
		if err != nil {
			log.Println(addr+" - Failed to read host from address:", err)
		}

		var reportedAIPDB = false
		if conn.Config.AbuseIPDBReport && shouldReport(host, conn.Config.AbuseIPDBCooldownH, true) {
			respCode, err := Report(host, conn.Config.AbuseIPDBKey, conn.Config.Port)
			if err != nil {
				log.Println(addr+" - Failed to report on AbuseIPDB:", err)
			} else if respCode == 200 {
				reportedAIPDB = true
			}
		}

		var reportedWebhook = false
		if conn.Config.WebhookReport && shouldReport(host, conn.Config.WebhookCooldownH, false) {
			err := SendWebhook(
				conn.Config,
				host,
				reportedAIPDB,
				conn.DidHandshake,
				conn.DidPing,
			)
			if err != nil {
				log.Println(addr+" - Failed to report on webhook:", err)
			} else {
				reportedWebhook = true
			}
		}

		AddToCache(
			host,
			// The maximum time between the AbuseIPDB and Webhook report (in hours)
			time.Duration(math.Max(float64(conn.Config.AbuseIPDBCooldownH), float64(conn.Config.WebhookCooldownH)))*time.Hour,
			types.Report{
				Datetime:        time.Now(),
				PacketsCount:    conn.PacketsReceived,
				ReportedAIPDB:   reportedAIPDB,
				ReportedWebhook: reportedWebhook,
				Handshake:       conn.DidHandshake,
				Ping:            conn.DidPing,
			},
		)

		if conn.Config.WriteHistory {
			// Open history file
			file, err := os.OpenFile(conn.Config.HistoryFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // 644 = rw-,r--,r--
			if err != nil {
				log.Println(addr+" - Failed to open history file:", err)
			}
			defer file.Close()

			t := time.Now()

			_, err = file.WriteString(fmt.Sprintf(
				"%s,%s,%d,%t,%t,%t\n",
				t.Format("2006-01-02 15:04:05"),
				host,
				conn.PacketsReceived,
				reportedAIPDB,
				conn.DidHandshake,
				conn.DidPing,
			))
			if err != nil {
				log.Println(addr+" - Failed to write history:", err)
			}
		}
	}
}
