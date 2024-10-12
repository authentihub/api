package models

import (
	"database/sql"
)

type User struct {
	BaseModel
	Uid      string         `json:"uid" gorm:"unique"`
	Username string         `json:"username" gorm:"unique"`
	Password sql.NullString `json:"password"`
}
