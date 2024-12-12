package models

import (
	"github.com/uptrace/bun"
)

type Transactions struct {
	bun.BaseModel `bun:"table:transactions"`

	Id          int64 `bun:",type:serial,autoincrement,pk"`
	TotalAmount int64 `json:"total_amount"`
	CreatedBy
	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
