package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Content is a map of fields
type Content map[string]Field

// NewContent is a constructor
func NewContent(fields []Field) *Content {
	c := Content(map[string]Field{})

	for _, field := range fields {
		c[field.Code] = field
	}

	return &c
}

// Fields return list of fields
func (c *Content) Fields() []Field {
	var fields []Field
	for _, field := range *c {
		fields = append(fields, field)
	}

	return fields
}

// Validate data
func (c *Content) Validate(inputData map[string]json.RawMessage, checkOriginalRequired bool) error {
	var errstrings []string

	// Валидация наличия всех обязательных полей. Используется при создании элемента.
	if checkOriginalRequired {
		for fieldCode, field := range *c {
			if field.Required {
				if _, exists := inputData[fieldCode]; !exists {
					errstrings = append(errstrings, fmt.Sprintf("%s: field is required", fieldCode))
				}
			}
		}
	}

	for fieldCode, value := range inputData {
		field, ok := (*c)[fieldCode]
		if !ok {
			errstrings = append(errstrings, fmt.Sprintf("Field with Code: %q not found", fieldCode))
			continue
		}
		_, err := field.Validate(value)
		if err != nil {
			errstrings = append(errstrings, err.Error())
		}
	}

	// Если возникла хоть одна ошибка - возвращаем её
	if len(errstrings) > 0 {
		return fmt.Errorf(strings.Join(errstrings, "\n"))
	}

	return nil
}
