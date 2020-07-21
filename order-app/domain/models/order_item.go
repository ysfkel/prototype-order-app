package models

type OrderItem struct {
	ID       uint64  `gorm:"primary_key;" json:"id" csv:"id"`
	OrderID  uint64  `json:"order_id" csv:"order_id"`
	Price    float64 `json:"price_per_unit" csv:"price_per_unit"`
	Quantity uint64  `json:"quantity" csv:"quantity"`
	Product  string  `json:"product" csv:"product"`
	Delivery *Delivery
	Order    *Order
}

type IOrderItemRepository interface {
	Create(m *OrderItem) (*OrderItem, error)
	Get(id uint64) (*OrderItem, error)
	List() ([]*OrderItem, error)
	FindLike(value string) ([]*OrderItem, error)
	FindByOrderID(orderID uint64) ([]*OrderItem, error)
}
