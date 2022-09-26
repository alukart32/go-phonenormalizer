package repo

import (
	"context"
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"log"
	"strings"

	"alukart32.com/phoneNormalizer/internal/entity"
)

var (
	//go:embed sql/update_phone.sql
	updatePhoneSql string
	//go:embed sql/select_phone.sql
	selectPhoneSql string
)

type PhoneRepo struct {
	DB           *sql.DB
	GenerateUUID GenerateUUID
}

func NewPhoneRepo(db *sql.DB, g GenerateUUID) *PhoneRepo {
	return &PhoneRepo{
		DB:           db,
		GenerateUUID: g,
	}
}

func (p *PhoneRepo) FetchAll(ctx context.Context) ([]*entity.Phone, error) {
	fail := func(err error) ([]*entity.Phone, error) {
		return nil, fmt.Errorf("phone fetch all: %v", err)
	}

	log.Println("fetchAll phones...")

	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}
	defer tx.Commit()

	phones := make([]*entity.Phone, 0)
	rows, err := tx.QueryContext(ctx, selectPhoneSql)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return phones, nil
		}
		return fail(err)
	}

	for rows.Next() {
		p := &entity.Phone{}
		if err = rows.Scan(&p.ID, &p.Number); err != nil {
			return fail(err)
		}
		phones = append(phones, p)
	}
	if rows.Err() != nil {
		return fail(rows.Err())
	}
	if err = rows.Close(); err != nil {
		return fail(err)
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	return phones, nil
}

func (p *PhoneRepo) BulkUpdate(ctx context.Context, phones ...*entity.Phone) error {
	fail := func(err error) error {
		return fmt.Errorf("phone bulk update: %v", err)
	}

	if len(phones) == 0 {
		return fail(fmt.Errorf("empty phones list"))
	}

	tx, err := p.DB.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()

	placeHolders, args := getBulkUpdateCTE(phones)
	stmtBulk := fmt.Sprintf(updatePhoneSql, strings.Join(placeHolders, ","))

	if _, err := tx.ExecContext(ctx, stmtBulk, args...); err != nil {
		return fail(err)
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	return nil
}

// getBulkUpdateCTE prepares placeholders and args for update CTE.
//
// Placeholders looks like "($%d, $%d)"? because entity.Phone has 2 fields: ID, Number.
//
// Example:
//
//	placeHolders: {
//		"($1, $2)",
//		"($3, $4)",
//	}
//
//	args: {
//		"1", 123456784,
//		"2", 134623123,
//	}
func getBulkUpdateCTE(phones []*entity.Phone) ([]string, []interface{}) {
	placeHolders := make([]string, 0, len(phones))
	args := make([]interface{}, 0, len(phones)*2)
	i := 0
	for _, p := range phones {
		placeHolders = append(placeHolders, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
		args = append(args, p.ID)
		args = append(args, p.Number)
		i++
	}
	return placeHolders, args
}
