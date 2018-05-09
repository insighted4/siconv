package storage

import (
	"errors"
	"regexp"
)

var (
	// Limit default
	Limit = 100

	// ErrNotFound is the error returned by storage if a resource cannot be found.
	ErrNotFound = errors.New("not found")

	// ErrAlreadyExists is the error returned by storage if a resource ID is taken during a create.
	ErrAlreadyExists = errors.New("ID already exists")

	// ErrInvalidUUID is the error returned by storage if ID is not valid UUID.
	ErrInvalidUUID = errors.New("invalid UUID")
)

type Pagination struct {
	Limit  int
	Offset int
}

// NewPagination is passed as a parameter to limit the total of rows.
func NewPagination(perPage, page int) *Pagination {
	return &Pagination{
		Limit:  perPage,
		Offset: page * perPage,
	}
}

// IsValidUUID checks if a given string is a valid UUID v4.
func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
	return r.MatchString(uuid)
}

// IsValidUUID checks if a given string is a valid UUID v4.
func IsValidUUIDV4(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}
