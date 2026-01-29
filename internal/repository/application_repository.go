package repository

import (
	"strings"
	"time"

	"github.com/Naomejoy/app-service/domain"
	"gorm.io/gorm"
)

type ApplicationRepository interface {
	Create(app *domain.Application) error
	GetByID(id uint64) (*domain.Application, error)
	Update(app *domain.Application) error
	Delete(id uint64) error
	List(params ApplicationListParams) ([]domain.Application, int64, error)
}

type appRepo struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) ApplicationRepository {
	return &appRepo{db: db}
}

type ApplicationListParams struct {
	Page     int
	PageSize int
	Q        string
	UserID   uint64
	From     *time.Time
	To       *time.Time
	Sort     string
	Order    string
}

func (r *appRepo) Create(app *domain.Application) error {
	return r.db.Create(app).Error
}

func (r *appRepo) GetByID(id uint64) (*domain.Application, error) {
	var app domain.Application
	err := r.db.Preload("Statuses").Preload("FileTypes").
		First(&app, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &app, nil
}

func (r *appRepo) Update(app *domain.Application) error {
	return r.db.Save(app).Error
}

func (r *appRepo) Delete(id uint64) error {
	return r.db.Delete(&domain.Application{}, "id = ?", id).Error
}

func (r *appRepo) List(params ApplicationListParams) ([]domain.Application, int64, error) {
	var apps []domain.Application
	var total int64

	if params.Page < 1 {
		params.Page = 1
	}
	if params.PageSize <= 0 || params.PageSize > 100 {
		params.PageSize = 20
	}

	query := r.db.Model(&domain.Application{})

	if params.Q != "" {
		q := "%" + params.Q + "%"
		query = query.Where(
			"name ILIKE ? OR description ILIKE ? OR code ILIKE ?", q, q, q,
		)
	}
	if params.UserID > 0 {
		query = query.Where("user_id = ?", params.UserID)
	}
	if params.From != nil {
		query = query.Where("created_at >= ?", params.From)
	}
	if params.To != nil {
		query = query.Where("created_at <= ?", params.To)
	}

	query.Count(&total)

	sortColumn := "created_at"
	if params.Sort == "name" || params.Sort == "code" {
		sortColumn = params.Sort
	}
	order := "desc"
	if strings.ToLower(params.Order) == "asc" {
		order = "asc"
	}

	err := query.Preload("Statuses").Preload("FileTypes").
		Order(sortColumn + " " + order).
		Offset((params.Page - 1) * params.PageSize).
		Limit(params.PageSize).
		Find(&apps).Error

	return apps, total, err
}
