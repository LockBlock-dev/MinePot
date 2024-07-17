package types

type DiscordWebhookPayload struct {
	Content   string                `json:"content,omitempty"`
	Username  string                `json:"username,omitempty"`
	AvatarURL string                `json:"avatar_url,omitempty"`
	Embeds    []DiscordWebhookEmbed `json:"embeds,omitempty"`
}

type DiscordWebhookEmbed struct {
	Title       string                  `json:"title,omitempty"`
	Description string                  `json:"description,omitempty"`
	URL         string                  `json:"url,omitempty"`
	Color       int                     `json:"color,omitempty"`
	Fields      []DiscordWebhookField   `json:"fields,omitempty"`
	Author      DiscordWebhookAuthor    `json:"author,omitempty"`
	Footer      DiscordWebhookFooter    `json:"footer,omitempty"`
	Image       DiscordWebhookImage     `json:"image,omitempty"`
	Thumbnail   DiscordWebhookThumbnail `json:"thumbnail,omitempty"`
	Video       DiscordWebhookVideo     `json:"video,omitempty"`
}

type DiscordWebhookField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type DiscordWebhookAuthor struct {
	Name    string `json:"name,omitempty"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type DiscordWebhookFooter struct {
	Text         string `json:"text,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type DiscordWebhookImage struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type DiscordWebhookThumbnail struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type DiscordWebhookVideo struct {
	URL    string `json:"url,omitempty"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}
