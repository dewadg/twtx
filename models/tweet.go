package models

// Tweet speaks for itself.
type Tweet struct {
	ID        string `json:"id_str"`
	Text      string `json:"text"`
	Truncated bool   `json:"truncated"`
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
}
