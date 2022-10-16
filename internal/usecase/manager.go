package usecase

import (
	"context"
	"fmt"
	"log"

	"alukart32.com/phoneNormalizer/internal/entity"
	"alukart32.com/phoneNormalizer/internal/usecase/repo"
)

type NormalizePhoneManager struct {
	repo         PhoneRepo
	phoneService NormalizeService
}

func NewManager(phoneRepo *repo.PhoneRepo, phoneService NormalizeService) NormalizeManager {
	return &NormalizePhoneManager{
		repo:         phoneRepo,
		phoneService: phoneService,
	}
}

// Normalize updates the phones in the database if necessary. A successful call sends
// err == nil to the "notify" channel, the context is not canceled.
func (m *NormalizePhoneManager) Normalize(ctx context.Context, notify chan<- error) {
	log.Println("start phone normalization...")
	phones, err := m.repo.FetchAll(ctx)
	if err != nil {
		notify <- err
	}

	log.Println("normalize phones...")
	bulkUpdate := make([]*entity.Phone, 0)
	for _, phone := range phones {
		if phone, err = m.phoneService.Normalize(phone); err == nil {
			fmt.Printf("\t%+v\n", *phone)
			bulkUpdate = append(bulkUpdate, phone)
		}
	}
	notify <- m.repo.BulkUpdate(ctx, bulkUpdate...)
}
