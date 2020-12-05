package global

import (
	"time"
)

type GVA_MODEL struct {
	ID        uint `gorm:"primarykey" json "id"`
	CreatedAt time.Time `gorm:"created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updatedAt"`
	DeletedAt time.Time `gorm:"index" json:"-"`
}
