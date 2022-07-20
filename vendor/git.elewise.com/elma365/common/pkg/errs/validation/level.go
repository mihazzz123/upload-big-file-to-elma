package validation

//go:generate ../../../tooling/bin/enumer -type=Level -trimprefix=Level -json -transform=snake -output=level_string.go .

// Level of validation error
type Level int8

const (
	// LevelCritical error break follow validation process
	LevelCritical Level = iota
	// LevelError is a standard error shown that data is invalid
	LevelError
	// LevelWarning is a just wraning
	LevelWarning
)
