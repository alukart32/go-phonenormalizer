package usecase

import (
	"testing"

	"alukart32.com/phoneNormalizer/config"
	"alukart32.com/phoneNormalizer/internal/entity"
)

func TestNormalize(t *testing.T) {
	useCase := &PhoneService{
		config.PhoneNormalizer{
			Regex: "\\D",
		},
	}

	testCases := []struct {
		input *entity.Phone
		want  string
	}{
		{
			input: &entity.Phone{ID: "1", Number: "1234567890"},
			want:  "1234567890",
		},
		{
			input: &entity.Phone{ID: "2", Number: "123 456 7891"},
			want:  "1234567891",
		},
		{
			input: &entity.Phone{ID: "3", Number: "(123) 456 7892"},
			want:  "1234567892",
		},
		{
			input: &entity.Phone{ID: "1", Number: "(123) 456-7893"},
			want:  "1234567893",
		},
		{
			input: &entity.Phone{ID: "1", Number: "123-456-7894"},
			want:  "1234567894",
		},
		{
			input: &entity.Phone{ID: "1", Number: "123-456-7890"},
			want:  "1234567890",
		},
		{
			input: &entity.Phone{ID: "1", Number: "1234567892"},
			want:  "1234567892",
		},
		{
			input: &entity.Phone{ID: "1", Number: "(123)456-7892"},
			want:  "1234567892",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input.Number, func(t *testing.T) {
			actual, err := useCase.Normalize(tc.input)
			if err != nil {
				t.Error(err)
			}
			if actual.Number != tc.want {
				t.Errorf("get: %s; want: %s", actual.Number, tc.want)
			}
		})
	}
}
