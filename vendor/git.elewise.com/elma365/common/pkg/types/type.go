package types

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"strings"

	_ "git.elewise.com/elma365/easylocalizer/generator" // nolint:golint

	"github.com/pkg/errors"
)

// Type is an enum of available types
type Type int

const (
	_ Type = iota

	// String is a plain string
	String

	// Float used for numbers by default
	Float

	// Integer used only for system collections
	Integer

	// Boolean is a boolean
	Boolean

	// DateTime is a RFC3339 timestamp
	DateTime

	// Duration is a number in seconds
	Duration

	// Category is a hierarchical tree linked with collection
	Category

	// Tag is an enum
	//
	// Tag can be static or enhancable (non-privileged user may add tags on item editing)
	Tag

	// Money is a pair of currency and cents value (int)
	Money

	// File is a hash code of file in storage
	File

	// Phone record
	Phone

	// Email record
	Email

	// Image is a file with preview cache?
	//
	// Warning: not implemented yet
	Image

	// Status is an static enum. It must be no more than one per collection.
	Status

	// Version semver compatible?
	//
	// Warning: not implemented yet
	Version

	// JSON object
	JSON

	// SysUser reference to user
	SysUser

	// FullName record
	FullName

	// Link record
	Link

	// SysOSNode reference to orgstruct node
	SysOSNode

	// SysCollection reference to collection (not element of collection)
	SysCollection

	// RefItem references to element of any collection
	RefItem

	// Enum type
	Enum

	// Table type
	Table

	// Account type
	Account

	// Role type
	Role
)

//nolint: gochecknoglobals // это должна быть константа, но Go так не умеет
var typeNames = map[Type]string{
	String:        "STRING",
	Float:         "FLOAT",
	Integer:       "INTEGER",
	Boolean:       "BOOLEAN",
	DateTime:      "DATETIME",
	Duration:      "DURATION",
	Category:      "CATEGORY",
	Tag:           "TAG",
	Money:         "MONEY",
	File:          "FILE",
	Phone:         "PHONE",
	Email:         "EMAIL",
	Image:         "IMAGE",
	Status:        "STATUS",
	Version:       "VERSION",
	JSON:          "JSON",
	SysUser:       "SYS_USER",
	FullName:      "FULL_NAME",
	Link:          "LINK",
	SysOSNode:     "SYS_OSNODE",
	SysCollection: "SYS_COLLECTION",
	RefItem:       "REF_ITEM",
	Enum:          "ENUM",
	Table:         "TABLE",
	Account:       "ACCOUNT",
	Role:          "ROLE",
}

//nolint: gochecknoglobals // это должна быть константа, но Go так не умеет
var typeValues = map[string]Type{
	"STRING":         String,
	"FLOAT":          Float,
	"INTEGER":        Integer,
	"BOOLEAN":        Boolean,
	"DATETIME":       DateTime,
	"DURATION":       Duration,
	"CATEGORY":       Category,
	"TAG":            Tag,
	"MONEY":          Money,
	"FILE":           File,
	"PHONE":          Phone,
	"EMAIL":          Email,
	"IMAGE":          Image,
	"STATUS":         Status,
	"VERSION":        Version,
	"JSON":           JSON,
	"SYS_USER":       SysUser,
	"FULL_NAME":      FullName,
	"LINK":           Link,
	"SYS_OSNODE":     SysOSNode,
	"SYS_COLLECTION": SysCollection,
	"REF_ITEM":       RefItem,
	"ENUM":           Enum,
	"TABLE":          Table,
	"ACCOUNT":        Account,
	"ROLE":           Role,
}

// FromString сформировать Type из строки
func FromString(s string) (Type, error) {
	v, ok := typeValues[strings.ToUpper(s)]
	if !ok {
		return 0, errors.New("no such type")
	}

	return v, nil
}

// String implements sys.Stringer interface
func (t Type) String() string {
	if s, ok := typeNames[t]; ok {
		return s
	}

	return fmt.Sprintf("Unknown type %d", t)
}

// MarshalJSON implements json.Marshaler interface
func (t Type) MarshalJSON() ([]byte, error) {
	if s, ok := typeNames[t]; ok {
		res := make([]byte, len(s)+2)
		res[0] = '"'
		copy(res[1:], s)
		res[len(res)-1] = '"'

		return res, nil
	}

	return nil, errors.Errorf("unknown type %d", t)
}

// UnmarshalJSON implements json.Unmarshaler interface
func (t *Type) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return errors.Errorf("cannot unmarshal %s as type", data)
	}
	tt, ok := typeValues[s]
	if !ok {
		return errors.Errorf("unknown type %q", s)
	}
	*t = tt

	return nil
}

// Generate implements testing/quick.Generator
//
// Deprecated: only for tests
func (Type) Generate(r *rand.Rand, _ int) reflect.Value {
	i := r.Intn(len(typeNames)) + 1
	t := Type(i)
	return reflect.ValueOf(t)
}
