package mdynamodb

import "errors"

var (
	ErrGetItemNotFound = errors.New("record not found")
)
