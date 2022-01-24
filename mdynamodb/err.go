package mdynamodb

import "errors"

var (
	ErrPutItemMustAttributeNotExists = errors.New("putItem must have attribute_not_exists")
	ErrGetItemNotFound               = errors.New("getItem not found")
)
