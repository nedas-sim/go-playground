package entities

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name     string          `json:"name"`
	Price    decimal.Decimal `json:"price"`
	Quantity int             `json:"quantity"`
}