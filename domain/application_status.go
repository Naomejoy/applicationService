package domain

import "time"

type ApplicationStatus struct {
	ID            uint64    `gorm:"primaryKey;column:id" json:"id"`
	ApplicationID uint64    `gorm:"column:application_id;not null;index" json:"applicationId"`
	UserID        uint64    `gorm:"column:user_id;not null;index" json:"userId"`
	Status        string    `gorm:"column:status;size:50;not null" json:"status"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`

	Application Application `gorm:"foreignKey:ApplicationID" json:"-"`
}

func (ApplicationStatus) TableName() string {
	return "application_status"
}
