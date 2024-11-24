package rows

import (
	"ChatsService/internal/models/interfaces"

	"database/sql"
)

type SqlRowsAdapter struct {
	rows *sql.Rows
}

func NewSqlRowsAdapter(rows *sql.Rows) interfaces.RowsAdapter {
	return &SqlRowsAdapter{rows}
}

func (r *SqlRowsAdapter) Next() bool {
	return r.rows.Next()
}

func (r *SqlRowsAdapter) Scan(dest ...interface{}) error {
	return r.rows.Scan(dest...)
}

func (r *SqlRowsAdapter) Close() error {
	return r.rows.Close()
}

func (r *SqlRowsAdapter) Err() error {
	return r.rows.Err()
}
