package tracker

import "fmt"

type Error struct {
	Errors        interface{} `json:"errors"`
	ErrorMessages []string    `json:"errorMessages"`
	StatusCode    int         `json:"statusCode"`
}

func (e Error) Error() string {
	return fmt.Sprint(e.ErrorMessages)
}
