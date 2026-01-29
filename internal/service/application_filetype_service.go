package service

import (
	"github.com/Naomejoy/app-service/internal/repository"

	"github.com/Naomejoy/app-service/domain"
)

type ApplicationFileTypeService struct {
	fileRepo repository.ApplicationFileTypeRepository
}

func NewApplicationFileTypeService(fileRepo repository.ApplicationFileTypeRepository) *ApplicationFileTypeService {
	return &ApplicationFileTypeService{fileRepo: fileRepo}
}

func (s *ApplicationFileTypeService) Add(appID uint64, fileTypeName string) error {
	return s.fileRepo.Add(&domain.ApplicationUploadedFileType{
		ApplicationID: appID,
		FileTypeName:  fileTypeName,
	})
}

func (s *ApplicationFileTypeService) Delete(fileTypeID uint64) error {
	return s.fileRepo.Delete(fileTypeID)
}

func (s *ApplicationFileTypeService) List(appID uint64) ([]domain.ApplicationUploadedFileType, error) {
	return s.fileRepo.ListByApplication(appID)
}
