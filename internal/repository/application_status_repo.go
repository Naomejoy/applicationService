package repository

import (
	"github.com/Naomejoy/app-service/domain"
	"gorm.io/gorm"
)

type ApplicationStatusRepository interface {
	Add(status *domain.ApplicationStatus) error
	ListByApplication(appID uint64, page, pageSize int) ([]domain.ApplicationStatus, int64, error)
}

type statusRepo struct {
	db *gorm.DB
}

func NewApplicationStatusRepository(db *gorm.DB) ApplicationStatusRepository {
	return &statusRepo{db: db}
}

func (r *statusRepo) Add(status *domain.ApplicationStatus) error {
	return r.db.Create(status).Error
}

func (r *statusRepo) ListByApplication(appID uint64, page, pageSize int) ([]domain.ApplicationStatus, int64, error) {
	var statuses []domain.ApplicationStatus
	var total int64

	if page < 1 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	query := r.db.Model(&domain.ApplicationStatus{}).Where("application_id = ?", appID)

	query.Count(&total)

	err := query.Order("created_at desc").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&statuses).Error

	return statuses, total, err
}
