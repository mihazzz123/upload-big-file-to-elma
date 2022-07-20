package permissions

//go:generate ../../tooling/bin/enumer -json -transform=snake -type=GroupType -trimprefix=GroupType -output=group_type_string.go group_type.go

// GroupType тип группы владельца доступа
type GroupType int8

const (
	// GroupTypeUser — пользователь
	GroupTypeUser GroupType = iota
	// GroupTypeGroup — группа
	GroupTypeGroup
	// GroupTypeOrgstruct — элемент оргуструктуры
	GroupTypeOrgstruct
	// GroupTypeRole — роль (группа с единственным участником
	GroupTypeRole
)
