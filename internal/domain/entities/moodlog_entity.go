package entities

import "time"

type MoodLogEntity struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	Mood      int       `gorm:"not null" json:"mood"`
	Note      string    `json:"note"`
	Causes    string    `gorm:"not null;index" json:"causes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User UserEntity `gorm:"foreignKey:UserID" json:"-"`
}
