package models

import "time"

// Article représente la structure des données pour un article
type Article struct {
    ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    Title         string    `gorm:"size:255;not null" json:"title"`
    Content       string    `gorm:"type:text;not null" json:"content"`
    Author        string    `gorm:"size:100;not null" json:"author"`
    PublishedDate time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"published_date"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
}
