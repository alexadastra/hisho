package models

import (
	"github.com/pkg/errors"
)

var (
	stringToIntMap = map[string]int64{
		"TERM_TODAY": 0,
		"TERM_WEEK":  1,
		"TERM_OTHER": 2,
	}

	intToStringMap = map[int64]string{
		0: "TERM_TODAY",
		1: "TERM_WEEK",
		2: "TERM_OTHER",
	}
)

// Term is a wrapper around Term Enum
type Term struct {
	Value    string
	ValueInt int64
}

// NewTermFromString constructs new Term
func NewTermFromString(value string) (*Term, error) {
	if _, ok := stringToIntMap[value]; !ok {
		return nil, errors.Errorf("undefined value: %s", value)
	}

	return &Term{
		Value:    value,
		ValueInt: stringToIntMap[value],
	}, nil
}

// NewTermFromInt constructs new Term
func NewTermFromInt(value int64) (*Term, error) {
	if _, ok := intToStringMap[value]; !ok {
		return nil, errors.Errorf("undefined value: %d", value)
	}

	return &Term{
		Value:    intToStringMap[value],
		ValueInt: value,
	}, nil
}
