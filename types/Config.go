package types

type Config struct {
	Debug bool `json:"debug"`

	WriteLogs bool   `json:"writeLogs"`
	LogFile   string `json:"logFile"`

	WriteHistory bool   `json:"writeHistory"`
	HistoryFile  string `json:"historyFile"`

	Port           int `json:"port"`
	PingDelayMinMs int `json:"pingDelayMinMs"`
	PingDelayMaxMs int `json:"pingDelayMaxMs"`
	IdleTimeoutS   int `json:"IdleTimeoutS"`

	ReportThreshold int `json:"reportThreshold"`

	AbuseIPDBReport    bool   `json:"abuseIPDBReport"`
	AbuseIPDBKey       string `json:"abuseIPDBKey"`
	AbuseIPDBCooldownH int    `json:"abuseIPDBCooldownH"`

	WebhookReport     bool   `json:"webhookReport"`
	WebhookUrl        string `json:"webhookUrl"`
	WebhookCooldownH  int    `json:"webhookCooldownH"`
	WebhookEmbedColor string `json:"webhookEmbedColor"`

	StatusResponse     bool         `json:"statusResponse"`
	StatusResponseData ServerStatus `json:"statusResponseData"`
	FaviconPath        string       `json:"faviconPath"`
	RandomVersion      bool         `json:"randomVersion`
}
