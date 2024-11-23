package exception

import (
	"fmt"

	"net/http"
)

type NotFoundException struct {
	Code        int
	Description string
}

func NewNotFoundException(description string) *NotFoundException {
	return &NotFoundException{
		Code:        http.StatusNotFound,
		Description: description,
	}
}

func (e NotFoundException) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Description)
}
