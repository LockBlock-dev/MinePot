package util

import (
	"time"

	"github.com/LockBlock-dev/MinePot/types"
	"github.com/muesli/cache2go"
)

func AddToCache(key interface{}, lifeSpan time.Duration, data interface{}) bool {
	exists := cache2go.Cache("MinePot").Exists(key)
	if !exists {
		cache2go.Cache("MinePot").Add(key, lifeSpan, data)
		return true
	}

	return false
}

func shouldReport(host string, cooldown int, reportType bool) bool {
	item, err := cache2go.Cache("MinePot").Value(host)
	if err != nil {
		return true
	}

	if reportType {
		// Check if the report was reported to AIPDB and if it's older than the cooldown
		report := item.Data().(types.Report)
		if !report.ReportedAIPDB && time.Since(report.Datetime) > (time.Duration(cooldown)*time.Hour) {
			return true
		}
	} else {
		// Check if the report was reported to the webhook and if it's older than the cooldown
		report := item.Data().(types.Report)
		if !report.ReportedWebhook && time.Since(report.Datetime) > (time.Duration(cooldown)*time.Hour) {
			return true
		}
	}

	return false
}
