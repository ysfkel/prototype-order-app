package models

import (
	"time"
)

type Order struct {
	ID         uint64       `gorm:"primary_key;" json:"id" csv:"id"`
	CreatedAt  time.Time    `json:"created_at" csv:"created_at"`
	Name       string       `json:"order_name" csv:"order_name"`
	CustomerID string       `json:"customer_id" csv:"customer_id"`
	OrderItems []*OrderItem // has many relationship
}

type IOrderRepository interface {
	Create(m *Order) (*Order, error)
	Get(id uint64) (*Order, error)
	List(offSet int, limit int) ([]*Order, error)
	GetCount(queryModel *Order) (int, error)
	FindLike(value string, offSet int, limit int) ([]*Order, error)
	ListByProductName(productName string, offSet int, limit int) ([]*Order, error)
	ListByDateRange(startDate time.Time, endDate time.Time, offSet int, limit int) ([]*Order, error)
	FindLikeAndDate(value string, startDate time.Time, endDate time.Time, offSet int, limit int) ([]*Order, error)
	ListByProductNameAndDate(productName string, startDate time.Time, endDate time.Time, offSet int, limit int) ([]*Order, error)
}
