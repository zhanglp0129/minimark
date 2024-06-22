package dao

import (
	"encoding/json"
	"errors"
	"github.com/shopspring/decimal"
)

type ChargeType byte

const (
	ChargeByNumber ChargeType = iota
	ChargeByWeight
)

func (ct *ChargeType) MarshalJSON() ([]byte, error) {
	if *ct > 1 {
		return nil, errors.New("不支持的计价方式")
	}
	return json.Marshal(ct.String())
}

func (ct *ChargeType) String() string {
	switch *ct {
	case ChargeByNumber:
		return "按个数收费"
	case ChargeByWeight:
		return "散装称重"
	}
	return ""
}

func (ct *ChargeType) UnmarshalJSON(jsonBytes []byte) error {
	str := string(jsonBytes)
	switch str {
	case "\"按个数收费\"":
		*ct = ChargeByNumber
	case "\"散装称重\"":
		*ct = ChargeByWeight
	default:
		return errors.New("不支持的计价方式")
	}
	return nil
}

type Goods struct {
	ID         int             `json:"id"`
	GTIN       *string         `gorm:"size:15" json:"gtin,omitempty"`
	Name       string          `gorm:"size:50;not null;unique" json:"name"`
	CategoryID *int            `json:"categoryId,omitempty"`
	Category   *Category       `json:"category,omitempty"`
	Price      decimal.Decimal `gorm:"type:decimal(10,2);not null;check:price>=0" json:"price"`
	ChargeType ChargeType      `gorm:"not null;check:charge_type<=1" json:"chargeType"`
	Stock      decimal.Decimal `gorm:"type:decimal(10,3);not null;check:stock>=0" json:"stock"`
}
