package utils

import "github.com/go-playground/validator"

type (
	// PathQuery represents a user date submission.
	PathQuery struct {
		Year    uint64 `json:"year" validate:"required"`
		Month   uint64 `json:"month" validate:"required"`
		Day     uint64 `json:"day" validate:"required"`
		Path    int    `json:"path"`
		Message string `json:"message"`
	}

	// PathQueryValidator ensures date inputs are valid.
	PathQueryValidator struct {
		Validator *validator.Validate
	}
)

// Validate ensures data sent to the server is valid,
// and informs the user of invalide submissions.
func (cv *PathQueryValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

// CalculateLifePath sums each individual date portion, then
// sums again if a date portion is greater than 9 and does not equal 11.
func CalculateLifePath(d uint64, m uint64, y uint64) int {
	day := total(sum(d))
	month := total(sum(m))
	year := total(sum(y))
	return day + month + year
}

func sum(i uint64) (sum int) {
	b64 := uint64(10)
	for ; i > 0; i /= b64 {
		sum += int(i % b64)
	}
	return
}
func total(i int) (total int) {
	if i > 9 && i != 11 {
		total = sum(uint64(i))
	} else {
		total = i
	}
	return
}
