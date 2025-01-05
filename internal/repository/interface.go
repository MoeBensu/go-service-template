package repository

import (
	"context"
	"yourproject/internal/model"
)

type Repository interface {
	GetExamples(ctx context.Context) ([]model.Example, error)
	CreateExample(ctx context.Context, example *model.Example) (*model.Example, error)
}
