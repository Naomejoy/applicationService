package service

import (
	"errors"
	"math"

	"github.com/Naomejoy/app-service/internal/repository"

	"github.com/Naomejoy/app-service/domain"
)

// ApplicationService handles business logic for Applications
type ApplicationService struct {
	appRepo repository.ApplicationRepository
}

// Constructor
func NewApplicationService(appRepo repository.ApplicationRepository) *ApplicationService {
	return &ApplicationService{
		appRepo: appRepo,
	}
}

// Create a new application
func (s *ApplicationService) Create(app *domain.Application) error {
	if app.Name == "" || app.Code == "" {
		return errors.New("name and code are required")
	}
	return s.appRepo.Create(app)
}

// Get application by ID
func (s *ApplicationService) GetByID(id uint64) (*domain.Application, error) {
	return s.appRepo.GetByID(id)
}

// Update application
func (s *ApplicationService) Update(app *domain.Application) error {
	if app.ID == 0 {
		return errors.New("invalid application ID")
	}
	return s.appRepo.Update(app)
}

// Soft delete application
func (s *ApplicationService) Delete(id uint64) error {
	return s.appRepo.Delete(id)
}

// List applications with pagination
func (s *ApplicationService) List(params repository.ApplicationListParams) (*ListResponse, error) {
	apps, total, err := s.appRepo.List(params)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(params.PageSize)))

	data := make([]domain.Application, len(apps))
	copy(data, apps)

	return &ListResponse{
		Data: data,
		Meta: PaginationMeta{
			Page:       params.Page,
			PageSize:   params.PageSize,
			Total:      total,
			TotalPages: totalPages,
		},
	}, nil
}
