package entities

type Image struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Name         string `json:"name"`
	OriginalName string `json:"originalName"`
	MimeType     string `json:"mimeType"`
	Url          string `json:"url"`
	CreatedAt    int64  `gorm:"autoCreateTime"`
	UpdatedAt    int64  `gorm:"autoUpdateTime"`
}
