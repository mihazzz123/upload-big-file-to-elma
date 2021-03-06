// Code generated by "enumer -json -transform=snake -type=AccessAtom -trimprefix=AccessAtom -output=access_string.go access.go"; DO NOT EDIT.

//
package permissions

import (
	"encoding/json"
	"fmt"
)

const (
	_AccessAtomName_0 = "readcreate"
	_AccessAtomName_1 = "update"
	_AccessAtomName_2 = "delete"
	_AccessAtomName_3 = "assign"
	_AccessAtomName_4 = "bpmanage"
	_AccessAtomName_5 = "export"
	_AccessAtomName_6 = "import"
)

var (
	_AccessAtomIndex_0 = [...]uint8{0, 4, 10}
	_AccessAtomIndex_1 = [...]uint8{0, 6}
	_AccessAtomIndex_2 = [...]uint8{0, 6}
	_AccessAtomIndex_3 = [...]uint8{0, 6}
	_AccessAtomIndex_4 = [...]uint8{0, 8}
	_AccessAtomIndex_5 = [...]uint8{0, 6}
	_AccessAtomIndex_6 = [...]uint8{0, 6}
)

func (i AccessAtom) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _AccessAtomName_0[_AccessAtomIndex_0[i]:_AccessAtomIndex_0[i+1]]
	case i == 4:
		return _AccessAtomName_1
	case i == 8:
		return _AccessAtomName_2
	case i == 16:
		return _AccessAtomName_3
	case i == 32:
		return _AccessAtomName_4
	case i == 64:
		return _AccessAtomName_5
	case i == 128:
		return _AccessAtomName_6
	default:
		return fmt.Sprintf("AccessAtom(%d)", i)
	}
}

var _AccessAtomValues = []AccessAtom{1, 2, 4, 8, 16, 32, 64, 128}

var _AccessAtomNameToValueMap = map[string]AccessAtom{
	_AccessAtomName_0[0:4]:  1,
	_AccessAtomName_0[4:10]: 2,
	_AccessAtomName_1[0:6]:  4,
	_AccessAtomName_2[0:6]:  8,
	_AccessAtomName_3[0:6]:  16,
	_AccessAtomName_4[0:8]:  32,
	_AccessAtomName_5[0:6]:  64,
	_AccessAtomName_6[0:6]:  128,
}

// AccessAtomString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func AccessAtomString(s string) (AccessAtom, error) {
	if val, ok := _AccessAtomNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to AccessAtom values", s)
}

// AccessAtomValues returns all values of the enum
func AccessAtomValues() []AccessAtom {
	return _AccessAtomValues
}

// IsAAccessAtom returns "true" if the value is listed in the enum definition. "false" otherwise
func (i AccessAtom) IsAAccessAtom() bool {
	for _, v := range _AccessAtomValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for AccessAtom
func (i AccessAtom) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccessAtom
func (i *AccessAtom) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AccessAtom should be a string, got %s", data)
	}

	var err error
	*i, err = AccessAtomString(s)
	return err
}
