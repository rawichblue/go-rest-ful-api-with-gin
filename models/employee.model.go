package models

import (
	"github.com/uptrace/bun"
)

type Employee struct {
	bun.BaseModel `bun:"table:employees"`

	ID       int64  `bun:",type:serial,autoincrement,pk"`
	UserId   string `json:"userId"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Images   string `json:"images"`
	Role     string `json:"role"`
	Address  string `json:"address"`
	Phone    int64  `json:"phone"`
	CreatedBy
	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
