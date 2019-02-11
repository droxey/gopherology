package utils

import "github.com/go-playground/validator"

type (
	// PathQuery represents a user date submission.
	PathQuery struct {
		Year  uint16 `json:"year" form:"year" validate:"required"`
		Month uint16 `json:"month" form:"month" validate:"required,gte=1,lte=12"`
		Day   uint16 `json:"day" form:"day" validate:"required,gte=1,lte=31"`
	}

	// PathResponse returns the path.
	PathResponse struct {
		Number int    `json:"number"`
		URL    string `json:"more_info"`
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

// CalculateLifePath sums each individual date portion to calculate
// the user's life path number.
func CalculateLifePath(d uint16, m uint16, y uint16) int {
	day := total(d)
	month := total(m)
	year := total(y)
	final := (day + month + year)
	if final == 11 || final == 22 || final == 33 {
		return final
	}
	return total(uint16(final))
}

func process(i uint16) (sum int) {
	if i < 10 {
		return int(i)
	}
	b64 := uint16(10)
	for ; i > 0; i /= b64 {
		sum += int(i % b64)
	}
	return
}

func total(i uint16) (sum int) {
	sum = process(i)
	if sum > 9 {
		sum = process(uint16(sum))
	}
	return
}
