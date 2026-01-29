package domain

import "time"

type Application struct {
	ID          uint64 `gorm:"primaryKey;column:id" json:"id"`
	UserID      uint64 `gorm:"column:user_id;not null;index" json:"userId"`
	Name        string `gorm:"column:name;size:255;not null" json:"name"`
	Description string `gorm:"column:description" json:"description"`
	Code        string `gorm:"column:code;size:50;not null;uniqueIndex" json:"code"`

	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index" json:"deletedAt,omitempty"`

	Statuses  []ApplicationStatus           `gorm:"foreignKey:ApplicationID" json:"statuses,omitempty"`
	FileTypes []ApplicationUploadedFileType `gorm:"foreignKey:ApplicationID" json:"fileTypes,omitempty"`
}

func (Application) TableName() string {
	return "applications"
}
