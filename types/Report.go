package types

import (
	"time"
)

type Report struct {
	Datetime        time.Time
	PacketsCount    int
	ReportedAIPDB   bool
	ReportedWebhook bool
	Handshake       bool
	Ping            bool
}
