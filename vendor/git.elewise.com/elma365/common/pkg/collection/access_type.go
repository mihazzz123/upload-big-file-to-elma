package collection

//go:generate ../../tooling/bin/enumer -json -sql -transform=snake -type=AccessType -trimprefix=AccessType -output=access_type_string.go access_type.go

// AccessType типы ограничения доступа к коллекции
type AccessType int8

const (
	// AccessTypeNone не ограничивать доступ к коллекции
	AccessTypeNone AccessType = iota
	// AccessTypeCollection права назначаются на всю коллекцию целиком
	AccessTypeCollection
	// AccessTypeRow права назначаются на каждую строчку коллекции
	AccessTypeRow
	// AccessTypeDirectory права назначаются папке коллекции
	AccessTypeDirectory
)
