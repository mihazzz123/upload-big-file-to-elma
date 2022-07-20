package edition

//go:generate ../../tooling/bin/enumer -json -text -transform=snake -type=Edition -trimprefix=Edition -output=edition_string.go edition.go

// Edition редакция конкретной инсталляции
//
// Возможные значения:
// `lite` - версия, урезанный функционал, в данный момент - это quickbpm.io
// `standard` - стандарт standalone версия, средний набор функционала, 1 нода
// `enterprise` - версия для предприятий, самый большой функционал, множество нод
// `portal` - сборка портала на базе Lite
//
// swagger:strfmt string
type Edition int8

const (
	// Lite версия, урезанный функционал, в данный момент - это quickbpm.io
	Lite Edition = iota
	// Standard стандарт standalone версия, средний набор функционала, 1 нода
	Standard
	// Enterprise версия для предприятий, самый большой функционал, множество нод
	Enterprise
	// Portal сборка портала на базе Lite
	Portal
)
