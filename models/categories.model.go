package models

import (
	"github.com/uptrace/bun"
)

type Categories struct {
	bun.BaseModel `bun:"table:categories"`

	Id   int64  `bun:",type:serial,autoincrement,pk"`
	Name string `json:"name"`
	CreatedBy
	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
