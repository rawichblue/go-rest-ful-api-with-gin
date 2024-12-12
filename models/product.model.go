package models

import "github.com/uptrace/bun"

type Product struct {
	bun.BaseModel `bun:"table:products"`

	Id          int64   `bun:",type:serial,autoincrement,pk"`
	Name        string  `bun:"name"`
	Price       float64 `bun:"price"`
	Stock       float64 `bun:"stock"`
	CategoryId  int64   `bun:"category_id"`
	Description string  `bun:"description"`
	IsActived   bool    `bun:"is_actived,default:false"`
	CreatedBy
	CreateUnixTimestamp
	UpdateUnixTimestamp
	SoftDelete
}
