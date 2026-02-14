package strapiModel

import "time"

type Diary struct {
	ID                *uint      `gorm:"primaryKey"`
	Date              *string    `gorm:"type:VARCHAR(255)"`
	Summary           *string    `gorm:"type:LONGTEXT"`
	Story             *string    `gorm:"type:LONGTEXT"`
	GraphGain         *int       `gorm:"type:int"`
	GraphEmotional    *int       `gorm:"type:int"`
	GraphProductivity *int       `gorm:"type:int"`
	Task              *[]byte    `gorm:"type:json"`
	CreatedAt         *time.Time `gorm:"not null"`
	UpdatedAt         *time.Time `gorm:"not null"`
	PublishedAt       *time.Time `gorm:"not null"`
	CreatedByID       *uint      `gorm:"type:int"`
	UpdatedByID       *uint      `gorm:"type:int"`
}
