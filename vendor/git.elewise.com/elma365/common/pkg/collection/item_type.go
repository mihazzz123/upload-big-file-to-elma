package collection

//go:generate ../../tooling/bin/enumer -json -sql -transform=snake -type=Type -trimprefix=Type -output=item_type_string.go item_type.go

// Type - тип элементов коллекции
type Type int8

const (
	// TypeApplication приложение, обрабатывается в main
	TypeApplication Type = iota
	// TypeContract - контракт, обрабатывается сервисом contractor
	TypeContract
)
