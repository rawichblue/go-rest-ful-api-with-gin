package models

import "github.com/uptrace/bun"

type Role struct {
	bun.BaseModel `bun:"table:role"`

	Id          int64  `bun:",type:serial,autoincrement,pk"`
	Name        string `bun:"name"`
	Description string `bun:"description"`
	IsActived   bool   `bun:"is_actived,default:false"`
	CreatedBy
	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
