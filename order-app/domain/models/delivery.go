package models

type Delivery struct {
	ID                uint64 `gorm:"primary_key;" json:"_id" csv:"id"`
	OrderItemID       uint64 `json:"order_item_id" csv:"order_item_id"`
	DeliveredQuantity uint64 `json:"delivered_quantity" csv:"delivered_quantity"`
	OrderItem         *OrderItem
}

type IDeliveryRepository interface {
	Create(m *Delivery) (*Delivery, error)
	Get(id uint64) (*Delivery, error)
	List() ([]*Delivery, error)
	FindByOrderItemID(orderItemID uint64) (*Delivery, error)
}
