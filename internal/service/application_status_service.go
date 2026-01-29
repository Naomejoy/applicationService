package service

import (
	"github.com/Naomejoy/app-service/internal/repository"

	"math"

	"github.com/Naomejoy/app-service/domain"
)

type ApplicationStatusService struct {
	statusRepo repository.ApplicationStatusRepository
}

func NewApplicationStatusService(statusRepo repository.ApplicationStatusRepository) *ApplicationStatusService {
	return &ApplicationStatusService{statusRepo: statusRepo}
}

func (s *ApplicationStatusService) Add(appID, userID uint64, status string) error {
	return s.statusRepo.Add(&domain.ApplicationStatus{
		ApplicationID: appID,
		UserID:        userID,
		Status:        status,
	})
}

func (s *ApplicationStatusService) List(appID uint64, page, pageSize int) (*ListResponse, error) {
	statuses, total, err := s.statusRepo.ListByApplication(appID, page, pageSize)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	data := make([]domain.ApplicationStatus, len(statuses))
	copy(data, statuses)

	return &ListResponse{
		Data: data,
		Meta: PaginationMeta{
			Page:       page,
			PageSize:   pageSize,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}
