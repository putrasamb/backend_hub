package database

// Kind represents database kind.
type Kind[T any] struct {
	Read  T
	Write T
}
