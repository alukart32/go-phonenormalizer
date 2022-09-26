package usecase

import (
	"context"

	"alukart32.com/phoneNormalizer/internal/entity"
)

type (
	// manager
	NormalizeManager interface {
		Normalize(ctx context.Context, notify chan<- error)
	}

	// use case
	NormalizeService interface {
		Normalize(*entity.Phone) (*entity.Phone, error)
	}

	// repo
	PhoneRepo interface {
		FetchAll(ctx context.Context) ([]*entity.Phone, error)
		BulkUpdate(ctx context.Context, phones ...*entity.Phone) error
	}
)
