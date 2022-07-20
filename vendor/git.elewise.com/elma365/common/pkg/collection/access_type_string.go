// Code generated by "enumer -json -sql -transform=snake -type=AccessType -trimprefix=AccessType -output=access_type_string.go access_type.go"; DO NOT EDIT.

//
package collection

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const _AccessTypeName = "nonecollectionrowdirectory"

var _AccessTypeIndex = [...]uint8{0, 4, 14, 17, 26}

func (i AccessType) String() string {
	if i < 0 || i >= AccessType(len(_AccessTypeIndex)-1) {
		return fmt.Sprintf("AccessType(%d)", i)
	}
	return _AccessTypeName[_AccessTypeIndex[i]:_AccessTypeIndex[i+1]]
}

var _AccessTypeValues = []AccessType{0, 1, 2, 3}

var _AccessTypeNameToValueMap = map[string]AccessType{
	_AccessTypeName[0:4]:   0,
	_AccessTypeName[4:14]:  1,
	_AccessTypeName[14:17]: 2,
	_AccessTypeName[17:26]: 3,
}

// AccessTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func AccessTypeString(s string) (AccessType, error) {
	if val, ok := _AccessTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to AccessType values", s)
}

// AccessTypeValues returns all values of the enum
func AccessTypeValues() []AccessType {
	return _AccessTypeValues
}

// IsAAccessType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i AccessType) IsAAccessType() bool {
	for _, v := range _AccessTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for AccessType
func (i AccessType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccessType
func (i *AccessType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AccessType should be a string, got %s", data)
	}

	var err error
	*i, err = AccessTypeString(s)
	return err
}

func (i AccessType) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *AccessType) Scan(value interface{}) error {
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

	val, err := AccessTypeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
