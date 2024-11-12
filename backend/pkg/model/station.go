package model

import "time"

type Station struct {
	Id        int       `json:"-" gorm:"primaryKey"`
	UserId    int       `json:"user_id" gorm:""`
	CreatedAt time.Time `json:"created_at" gorm:""`
	UpdatedAt time.Time `json:"updated_at" gorm:""`
}

type ChargingPoint struct {
	Id          int       `json:"-" gorm:"primaryKey"`
	IsActive    bool      `json:"is_active" gorm:""`
	SupportType string    `json:"support_type" gorm:""`
	UpdatedAt   time.Time `json:"updated_at" gorm:""`
}
