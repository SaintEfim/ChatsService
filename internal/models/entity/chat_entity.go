package entity

import "github.com/google/uuid"

type ChatEntity struct {
	Id          uuid.UUID   `db:"id"`
	Name        string      `db:"name"`
	EmployeeIds []uuid.UUID `db:"employee_Ids"`
}
