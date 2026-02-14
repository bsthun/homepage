package model

import "time"

type PhotoSection struct {
	Id           *uint64     `gorm:"primaryKey" json:"id"`
	PhotoAlbum   *PhotoAlbum `gorm:"foreignKey:PhotoAlbumId" json:"photo_album"`
	PhotoAlbumId *uint64     `gorm:"not null" json:"photo_album_id"`
	Title        *string     `gorm:"type:VARCHAR(255); not null" json:"title"`
	Subtitle     *string     `gorm:"type:TEXT; not null" json:"subtitle"`
	Path         *string     `gorm:"type:TEXT; not null" json:"path"`
	Date         *time.Time  `gorm:"not null" json:"date"`
	CreatedAt    *time.Time  `gorm:"not null"` // Embedded field
	UpdatedAt    *time.Time  `gorm:"not null"` // Embedded field
}
