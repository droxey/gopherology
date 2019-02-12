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
		PathNumber     int    `json:"pathNumber"`
		URL            string `json:"detailsUrl"`
		IsMasterNumber bool   `json:"isMasterNumber"`
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
func CalculateLifePath(d uint16, m uint16, y uint16) (int, bool) {
	pathResult := (numerology(d) + numerology(m) + numerology(y))
	isMasterNumber := pathResult%11 == 0
	if isMasterNumber {
		return pathResult, isMasterNumber
	}
	return numerology(uint16(pathResult)), isMasterNumber
}

// Recursive function. Applys numerological rules when adding dates.
// If sum isn't a master number, call numerology a final time to sum double-digit ints.
func numerology(i uint16) (sum int) {
	if i < 10 {
		return int(i)
	}
	b64 := uint16(10)
	for ; i > 0; i /= b64 {
		sum += int(i % b64)
	}
	if sum > 9 {
		return numerology(uint16(sum))
	}
	return
}
