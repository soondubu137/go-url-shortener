package errors

import "errors"

var (
	ErrInvalidURL          = errors.New("invalid url")
	ErrURLAlreadyShortened = errors.New("url has already been shortened")
)
