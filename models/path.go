package path

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
