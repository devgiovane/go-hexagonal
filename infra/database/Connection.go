package database

type Connection[T any] interface {
	GetConnection() T
}
