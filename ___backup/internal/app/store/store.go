package store

// Store ...
type Store interface {
	BigFile() BigFileRepository
}
