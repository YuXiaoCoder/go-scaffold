package models

import (
	"time"

	"gorm.io/gorm"
)

// BasicMode 基础模型
type BasicMode struct {
	ID        int64          `gorm:"column:id;"`         // 主键，利用雪花算法生成
	CreatedAt time.Time      `gorm:"column:created_at;"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;"` // 删除时间
}
