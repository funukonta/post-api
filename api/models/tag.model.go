package models

type Tag struct {
	ID    int    `gorm:"primaryKey" json:"id"`
	Label string `gorm:"unique" json:"label"`
	Post  []Post `gorm:"many2many:post_tags;" json:"posts"`
}
