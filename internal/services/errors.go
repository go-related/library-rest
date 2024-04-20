package services

import "fmt"

type ServiceError struct {
	msg string
}

func (r *ServiceError) Error() string {
	return fmt.Sprintf("validation error: %v", r.msg)
}

func NewServiceError(msg string) *ServiceError {
	return &ServiceError{msg: msg}
}
