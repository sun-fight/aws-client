package mdynamodb

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrQueryThanOne   = errors.New("query than one")
)
