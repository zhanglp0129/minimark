package dao

import (
	"database/sql/driver"
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

// Time 自定义时间格式
type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = Time{t1}
	return err
}

func (t Time) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *Time) Scan(value interface{}) error {
	if value == nil {
		*t = Time{Time: time.Time{}}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*t = Time{Time: v}
	case []byte:
		parsed, err := time.Parse("2006-01-02 15:04:05", string(v))
		if err != nil {
			return err
		}
		*t = Time{parsed}
	case string:
		parsed, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return err
		}
		*t = Time{parsed}
	default:
		return fmt.Errorf("cannot scan type %T into MyTime", value)
	}
	return nil
}

type Order struct {
	ID          int             `json:"id"`
	TotalPaid   decimal.Decimal `gorm:"type:decimal(10,2);not null;check:total_paid>=0" json:"totalPaid"`
	PayMethodID *int            `json:"payMethodId,omitempty"`
	PayMethod   *PayMethod      `json:"payMethod,omitempty"`
	PayTime     Time            `gorm:"not null" json:"payTime"`
	Goods       []Goods         `gorm:"many2many:order_goods" json:"-"`
	OrderGoods  []OrderGoods    `json:"goods"`
}

type OrderGoods struct {
	GoodsID  int             `gorm:"primaryKey" json:"goodsId"`
	OrderID  int             `gorm:"primaryKey" json:"orderId"`
	Quantity decimal.Decimal `gorm:"type:decimal(10,3);not null;check:quantity>=0" json:"quantity"`
	Price    decimal.Decimal `gorm:"type:decimal(10,2);not null;check:price>=0" json:"price"`
}
