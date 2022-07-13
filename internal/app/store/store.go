package store

// Store ...
type Store interface {
	Bigfile() BigFileRepository
}
