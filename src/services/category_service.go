package services

import (
	"errors"
	"github.com/alireza-fa/blog-go/src/api/dto"
	"github.com/alireza-fa/blog-go/src/data/db"
	"github.com/alireza-fa/blog-go/src/data/models"
	"github.com/alireza-fa/blog-go/src/pkg/logging"
	"gorm.io/gorm"
)

type CategoryService struct {
	logger   logging.Logger
	database *gorm.DB
}

func NewCategoryService() *CategoryService {
	return &CategoryService{
		logger:   logging.NewLogger(),
		database: db.GetDb(),
	}
}

func (service *CategoryService) Create(categoryCreate *dto.CategoryCreate) (*dto.CategoryOutput, error) {
	exists, err := service.checkCategoryExists(categoryCreate.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("this category already exists")
	}

	var category models.Category
	category.Name = categoryCreate.Name

	if err := service.database.Create(&category).Error; err != nil {
		service.logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}

	categoryOutput := dto.CategoryOutput{Id: category.Id, Name: category.Name}

	return &categoryOutput, nil
}

func (service *CategoryService) checkCategoryExists(name string) (bool, error) {
	var exists bool
	if err := service.database.
		Model(models.Category{}).
		Select("count(*) > 0").
		Where("name = ?", name).
		Find(&exists).Error; err != nil {
		service.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (service *CategoryService) Update(categoryUpdate *dto.CategoryUpdate, id int) (*dto.CategoryOutput, error) {
	tx := service.database.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.
		Model(models.Category{}).
		Where("id = ?", id).
		Update("name", categoryUpdate.Name).Error; err != nil {
		tx.Rollback()
		service.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		service.logger.Error(logging.Postgres, logging.Commit, err.Error(), nil)
		return nil, err
	}

	categoryOutput := dto.CategoryOutput{Id: 0, Name: categoryUpdate.Name}
	return &categoryOutput, nil
}
