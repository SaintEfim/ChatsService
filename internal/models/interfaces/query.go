package interfaces

type Query[T any] interface {
	Get() string
	GetOneById() string
	Create() string
	Delete() string
	Update() string
}
