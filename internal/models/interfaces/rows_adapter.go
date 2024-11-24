package interfaces

type RowsAdapter interface {
	Next() bool
	Scan(dest ...interface{}) error
	Close() error
	Err() error
}
