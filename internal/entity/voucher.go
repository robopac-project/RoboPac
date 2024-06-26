package entity

import (
	"gorm.io/gorm"
)

type Voucher struct {
	ID      uint   `gorm:"primaryKey;unique"`
	Creator uint   `gorm:"size:255"`
	Code    string `gorm:"size:8"`
	Amount  uint
	TxHash  string

	gorm.Model
}

func (Voucher) TableName() string {
	return "voucher"
}
