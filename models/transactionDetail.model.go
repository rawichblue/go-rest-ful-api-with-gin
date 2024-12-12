package models

import (
	"github.com/uptrace/bun"
)

type TransactionsDetail struct {
	bun.BaseModel `bun:"table:transactions_detail"`

	Id            int64 `bun:",type:serial,autoincrement,pk"`
	TransactionId int64 `json:"transaction_id"`
	ProductId     int64 `json:"product_id"`
	Quantity      int64 `json:"quantity"`
	Price         int64 `json:"price"`
	CreatedBy
	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
