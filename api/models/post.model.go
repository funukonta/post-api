package models

type Post struct {
	ID      int    `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    []Tag  `gorm:"many2many:post_tags;" json:"tags"`
}

type PostRes struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    []Tag  `json:"tags"`
}

type PostReq struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}
