package models

import (
	"time"
)

type BaseModel struct {
	ID string `gorm:"type:varchar(36);primary_key"`
	//Name      string `gorm:"type:varchar(100)"`
	CreatedAt time.Time
}
