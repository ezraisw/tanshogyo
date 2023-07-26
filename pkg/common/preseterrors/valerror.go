package preseterrors

type (
	ValidationError struct {
		Message     string        `json:"message"`
		FieldErrors []*FieldError `json:"fieldErrors"`
	}

	FieldError struct {
		Field   string `json:"field"`
		Code    string `json:"code"`
		Message string `json:"message"`
	}
)

func (e ValidationError) Error() string {
	return e.Message
}

func (e ValidationError) MarshalAs() (map[string]any, error) {
	var fes []map[string]any

	for _, fieldErr := range e.FieldErrors {
		fes = append(fes, map[string]any{
			"field":   fieldErr.Field,
			"code":    fieldErr.Code,
			"message": fieldErr.Message,
		})
	}

	return map[string]any{
		"message":     e.Message,
		"fieldErrors": fes,
	}, nil
}

func IsValidationError(err error) bool {
	_, ok := err.(*ValidationError)
	return ok
}
