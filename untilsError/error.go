package untilsError

import (
	"errors"
	"fmt"
)

type Error struct {
	sentence string
}

func NewString(parameterSentence string, params ...string) error {
	err := Error{}
	err.sentence += fmt.Sprintf(parameterSentence, params)
	return errors.New(err.sentence)
}

func NewError(parameterSentence string, errs ...error) error {
	err := Error{}
	err.sentence += fmt.Sprintf(parameterSentence, errs)
	return errors.New(err.sentence)
}
