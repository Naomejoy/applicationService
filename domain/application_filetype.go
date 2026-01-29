package domain

import "time"

type ApplicationUploadedFileType struct {
	ID            uint64    `gorm:"primaryKey;column:id" json:"id"`
	ApplicationID uint64    `gorm:"column:application_id;not null;index" json:"applicationId"`
	FileTypeName  string    `gorm:"column:file_type_name;size:100;not null" json:"fileTypeName"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`

	Application Application `gorm:"foreignKey:ApplicationID" json:"-"`
}

func (ApplicationUploadedFileType) TableName() string {
	return "application_uploaded_file_type"
}
