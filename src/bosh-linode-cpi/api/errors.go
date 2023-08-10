package api

import "fmt"

type CloudError interface {
	error

	Type() string
}

type NotSupportedError struct{}

type VMCreationFailedError struct {
	reason string
}

func NewVMCreationFailedError(reason string) VMCreationFailedError {
	return VMCreationFailedError{reason: reason}
}

func (e VMCreationFailedError) Error() string { return fmt.Sprintf("VM failed to create: %v", e.reason) }
