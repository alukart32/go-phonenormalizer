package app

import (
	"context"
	"log"
	"time"

	"alukart32.com/phoneNormalizer/config"
	"alukart32.com/phoneNormalizer/internal/usecase"
	"alukart32.com/phoneNormalizer/internal/usecase/repo"
	"alukart32.com/phoneNormalizer/pkg/postgres"
	"github.com/google/uuid"
)

// Run prepares all dependecies and lunches the main use case.
func Run(cfg *config.Config) error {
	// DB: Postgres
	db, err := postgres.New(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}
	// Repository
	phoneRepo := repo.NewPhoneRepo(db, func() string { return uuid.New().String() })

	// Usecase
	phoneService := usecase.NewPhoneService(cfg.PhoneNormalizer)

	// Normalizer manager
	manager := usecase.NewManager(phoneRepo, phoneService)
	ctx, cancelCtx :=
		context.WithTimeout(context.Background(), time.Duration(cfg.PhoneNormalizer.Timeout)*time.Millisecond)
	defer cancelCtx()

	notify := make(chan error)
	go manager.Normalize(ctx, notify)

	// Waiting
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err = <-notify:
		return err
	}
}
