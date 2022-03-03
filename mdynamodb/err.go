package mdynamodb

import "errors"

var (
	ErrRecordNotFound      = errors.New("record not found")
	ErrQueryOneOnlyOneItem = errors.New("query than one")

	ErrUpdateItemNoSet         = errors.New("set updateItem to update")
	ErrUpdateItemOperationMode = errors.New("operationMode must set")
	ErrConditionMode           = errors.New("conditionMode not set")
	ErrLogicalMode             = errors.New("logicalMode not set")
)

type ErrNameToVal struct {
	Name string
}

func NewErrNamtToVal(name string) *ErrNameToVal {
	return &ErrNameToVal{
		Name: name,
	}
}

func (e *ErrNameToVal) Error() string {
	return "can't find name:" + e.Name
}

type ErrParamsValid struct {
	Name string
}

func NewErrParamsValid(name string) *ErrParamsValid {
	return &ErrParamsValid{
		Name: name,
	}
}

func (e *ErrParamsValid) Error() string {
	return "params valid:" + e.Name
}
