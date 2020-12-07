package models

import (
	"fmt"
	"regexp"
)

const (
	errEmptyField            = "Fields are empty: %v"
	errInvalidPositionFormat = "Invalid position format"
)

var (
	positionRegexp = regexp.MustCompile(`\d+(\-\d+)*`)
)

type Task struct {
	Tag      string
	Title    string
	Position string
	Count    int
}

func (t *Task) Validate() (err error) {
	emptyFields := make([]string, 0, 5)

	if t.Tag == "" {
		emptyFields = append(emptyFields, "tag")
	}

	if t.Title == "" {
		emptyFields = append(emptyFields, "title")
	}

	if t.Position == "" {
		emptyFields = append(emptyFields, "position")
	}

	if t.Count == 0 {
		emptyFields = append(emptyFields, "count")
	}

	if len(emptyFields) != 0 {
		err = fmt.Errorf(errEmptyField, emptyFields)
		return
	}

	if !positionRegexp.MatchString(t.Position) {
		err = fmt.Errorf(errInvalidPositionFormat)
	}

	return
}

type TaskResponse struct {
	Tag   *string
	Title string
	Date  string
}
