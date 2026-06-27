package apperr

import "strconv"

type AppErr struct {
	Code int
}

func (ae *AppErr) Error() string {
	return strconv.Itoa(ae.Code)
}

func NewAppError(code int) *AppErr {
	return &AppErr{
		Code: code,
	}
}