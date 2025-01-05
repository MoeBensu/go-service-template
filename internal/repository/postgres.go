package repository

import (
	"context"

	"yourproject/internal/model"

	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) Repository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) GetExamples(ctx context.Context) ([]model.Example, error) {
	var examples []model.Example
	if err := r.db.WithContext(ctx).Find(&examples).Error; err != nil {
		return nil, err
	}
	return examples, nil
}

func (r *PostgresRepository) CreateExample(ctx context.Context, example *model.Example) (*model.Example, error) {
	if err := r.db.WithContext(ctx).Create(example).Error; err != nil {
		return nil, err
	}
	return example, nil
}
