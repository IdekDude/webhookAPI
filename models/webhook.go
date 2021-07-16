package models


type Webhook struct {
	Username  string   `json:"username"`
	AvatarURL string   `json:"avatar_url"`
	Content   string   `json:"content"`
	Embeds    []Embeds `json:"embeds"`
}
type Embeds struct {
	Author      Author    `json:"author"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	Color       int       `json:"color"`
	Fields      []Fields  `json:"fields"`
	Thumbnail   Thumbnail `json:"thumbnail"`
	Image       ImageDiscord `json:"image"`
	Footer      Footer    `json:"footer"`
}
type Author struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}
type Fields struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}
type Thumbnail struct {
	URL string `json:"url"`
}
type ImageDiscord struct {
	URL string `json:"url"`
}
type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}