package resource

import "errors"

var (
	ErrResourceNotFound      = errors.New("Resource not found")
	ErrOptimisticLockFailure = errors.New("Database concurrent update fail")
)
