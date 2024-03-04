package tracker

import "fmt"

type Error struct {
	Errors        map[string]interface{} `json:"errors"`
	ErrorMessages []string               `json:"errorMessages"`
	StatusCode    int                    `json:"statusCode"`
}

func (e Error) Error() string {
	if len(e.Errors) > 0 {
		return fmt.Sprint(e.Errors)
	}
	return fmt.Sprint(e.ErrorMessages)
}
