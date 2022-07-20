package config

import (
	"encoding/json"
	"fmt"
	"strings"
)

//go:generate ../../tooling/bin/enumer -transform=snake -type=Solution -trimprefix=Solution -output=solution_string.go solution.go

// Solution редакция конкретной инсталляции
type Solution int8

const (
	// Saas в облаке
	Saas Solution = iota
	// Onpremise на сервере клиента
	Onpremise
)

// MarshalJSON implements the json.Marshaler interface for Solution
func (i Solution) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Solution
func (i *Solution) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Solution should be a string, got %s", data)
	}

	var err error
	*i, err = SolutionString(strings.ToLower(s))
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for Solution
func (i Solution) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Solution
func (i *Solution) UnmarshalText(text []byte) error {
	var err error
	var s = strings.ToLower(string(text))
	*i, err = SolutionString(s)
	return err
}
