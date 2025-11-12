package models

type Links struct {
	ID     int    `json:"links_num"`
	URL    string `json:"url"` // "url":"status"
	Status string `json:"status"`
}
