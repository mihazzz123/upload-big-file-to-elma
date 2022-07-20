// Code generated by "enumer -json -sql -transform=snake -type=Type -trimprefix=Type -output=item_type_string.go item_type.go"; DO NOT EDIT.

//
package collection

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const _TypeName = "applicationcontract"

var _TypeIndex = [...]uint8{0, 11, 19}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_TypeIndex)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _TypeName[_TypeIndex[i]:_TypeIndex[i+1]]
}

var _TypeValues = []Type{0, 1}

var _TypeNameToValueMap = map[string]Type{
	_TypeName[0:11]:  0,
	_TypeName[11:19]: 1,
}

// TypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TypeString(s string) (Type, error) {
	if val, ok := _TypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Type values", s)
}

// TypeValues returns all values of the enum
func TypeValues() []Type {
	return _TypeValues
}

// IsAType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Type) IsAType() bool {
	for _, v := range _TypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Type
func (i Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Type
func (i *Type) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Type should be a string, got %s", data)
	}

	var err error
	*i, err = TypeString(s)
	return err
}

func (i Type) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Type) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	val, err := TypeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}