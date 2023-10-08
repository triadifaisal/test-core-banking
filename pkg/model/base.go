package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	UUID        uuid.UUID      `gorm:"primarykey;type:uuid;default:uuid_generate_v4()"`
	CreatedDate time.Time      `gorm:"default:now()"`
	CreatedBy   uuid.NullUUID  `gorm:"type:uuid;"`
	UpdatedDate time.Time      `gorm:"default:now()"`
	UpdatedBy   uuid.NullUUID  `gorm:"type:uuid;"`
	DeletedDate gorm.DeletedAt `gorm:"index"`
	DeletedBy   uuid.NullUUID  `gorm:"type:uuid;"`
}
