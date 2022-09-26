package usecase

import (
	"fmt"
	"regexp"

	"alukart32.com/phoneNormalizer/config"
	"alukart32.com/phoneNormalizer/internal/entity"
)

type PhoneService struct {
	normalizeRule config.PhoneNormalizer
}

func NewPhoneService(pn config.PhoneNormalizer) *PhoneService {
	return &PhoneService{
		normalizeRule: pn,
	}
}

// Normalize updates the phone by regex if necessary. A successful call returns
// err == nil and updated phone entity.
func (puc *PhoneService) Normalize(phone *entity.Phone) (*entity.Phone, error) {
	fail := func(err error) (*entity.Phone, error) {
		return nil, fmt.Errorf("Normalize phone: %v", err)
	}

	if len(phone.Number) == 0 {
		return fail(fmt.Errorf("number length is 0"))
	}

	pattern, err := regexp.Compile(puc.normalizeRule.Regex)
	if err != nil {
		return fail(err)
	}

	if pattern.Match([]byte(phone.Number)) {
		phone.Number = pattern.ReplaceAllString(phone.Number, "")
	}

	return phone, nil
}
