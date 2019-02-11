package utils

import "github.com/go-playground/validator"

type (
	// PathQuery represents a user date submission.
	PathQuery struct {
		Year    uint16 `json:"year" validate:"required"`
		Month   uint16 `json:"month" validate:"required"`
		Day     uint16 `json:"day" validate:"required"`
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
func CalculateLifePath(d uint16, m uint16, y uint16) int {
	day := total(sum(d))
	month := total(sum(m))
	year := total(sum(y))

	println("Day", day, "Month", month, "Year", year)
	final := (day + month + year)
	if final == 11 || final == 22 || final == 33 {
		return final
	}
	return total(final)
}

func sum(i uint16) (sum int) {
	b64 := uint16(10)
	for ; i > 0; i /= b64 {
		sum += int(i % b64)
	}
	return
}
func total(i int) (total int) {
	if i < 10 {
		return i
	}
	for i > 9 {
		if i > 9 {
			total = sum(uint16(i))
		}
		i = total
	}
	return
}
