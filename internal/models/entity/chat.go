package entity

import (
	"database/sql/driver"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type UUIDArray []uuid.UUID

func (ua UUIDArray) Contains(target uuid.UUID) bool {
	for _, id := range ua {
		if id == target {
			return true
		}
	}
	return false
}

func (ua UUIDArray) Value() (driver.Value, error) {
	var strs []string
	for _, u := range ua {
		strs = append(strs, u.String())
	}
	return "{" + strings.Join(strs, ",") + "}", nil
}

func (ua *UUIDArray) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("scan source was not a string")
	}
	str = strings.Trim(str, "{}")
	parts := strings.Split(str, ",")
	for _, p := range parts {
		id, err := uuid.Parse(p)
		if err != nil {
			return err
		}
		*ua = append(*ua, id)
	}
	return nil
}

type Chat struct {
	gorm.Model
	Id          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primary_key"`
	Name        string
	IsGroup     bool      `gorm:"default:false"`
	EmployeeIds UUIDArray `gorm:"type:uuid[]"`
	Messages    []Message
}
