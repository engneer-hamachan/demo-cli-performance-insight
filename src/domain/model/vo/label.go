package vo

import (
	"fmt"
)

type Label string

func NewLabel(value string) (*Label, error) {

	if value == "" {
		err := fmt.Errorf("%s", "empty arg:Label New()")
		return nil, err
	}

	label := Label(value)

	return &label, nil
}
