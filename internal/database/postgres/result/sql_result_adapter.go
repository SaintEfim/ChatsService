package result

import (
	"ChatsService/internal/models/interfaces"

	"database/sql"
)

type SqlResultAdapter struct {
	sqlResult sql.Result
}

func NewSqlResultAdapter(sqlResult sql.Result) interfaces.ResultAdapter {
	return &SqlResultAdapter{
		sqlResult: sqlResult,
	}
}

func (r *SqlResultAdapter) RowsAffected() (int64, error) {
	return r.sqlResult.RowsAffected()
}
