package repository

import (
	"github.com/Naomejoy/app-service/domain"
	"gorm.io/gorm"
)

type ApplicationFileTypeRepository interface {
	Add(fileType *domain.ApplicationUploadedFileType) error
	Delete(id uint64) error
	ListByApplication(appID uint64) ([]domain.ApplicationUploadedFileType, error)
}

type fileTypeRepo struct {
	db *gorm.DB
}

func NewApplicationFileTypeRepository(db *gorm.DB) ApplicationFileTypeRepository {
	return &fileTypeRepo{db: db}
}

func (r *fileTypeRepo) Add(fileType *domain.ApplicationUploadedFileType) error {
	return r.db.Create(fileType).Error
}

func (r *fileTypeRepo) Delete(id uint64) error {
	return r.db.Delete(&domain.ApplicationUploadedFileType{}, "id = ?", id).Error
}

func (r *fileTypeRepo) ListByApplication(appID uint64) ([]domain.ApplicationUploadedFileType, error) {
	var fileTypes []domain.ApplicationUploadedFileType
	err := r.db.Where("application_id = ?", appID).Find(&fileTypes).Error
	return fileTypes, err
}
