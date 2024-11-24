package interfaces

type ResultAdapter interface {
	RowsAffected() (int64, error)
}
