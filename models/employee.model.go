package models

import (
	"github.com/uptrace/bun"
)

type Employee struct {
	bun.BaseModel `bun:"table:employees"`

	Id       int64  `bun:",type:serial,autoincrement,pk"`
	Email    string `bun:"email,unique"`
	UserId   string `json:"user_id"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Images   string `json:"images"`
	RoleId   int64  `json:"role_id"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	CreatedBy
	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
