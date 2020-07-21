package repository

import (
	"github.com/micro/go-log"

	"github.com/ysfkel/order-app/domain/models"
	persistence "github.com/ysfkel/order-app/infrastructure/persistence/database/postgresql"
)

//OrderItemRepository operation
type OrderItemRepository struct {
}

//NewOrderItemRepository instantiates and returns OrderItemRepository
func NewOrderItemRepository() models.IOrderItemRepository {
	return &OrderItemRepository{}
}

//Create OrderItemRepository model
func (o *OrderItemRepository) Create(m *models.OrderItem) (*models.OrderItem, error) {
	//Get a database session
	dbSession, err := persistence.GetDB()

	if err != nil {
		return nil, err
	}

	//create item
	db := dbSession.Create(m)

	if db.Error != nil {
		log.Log(db.Error)
		return nil, db.Error
	}
	return m, nil
}

//Get retrieve item by id
func (o *OrderItemRepository) Get(orderID uint64) (*models.OrderItem, error) {
	//Get a database session
	dbSession, err := persistence.GetDB()

	if err != nil {
		return nil, err
	}

	order := models.OrderItem{
		ID: orderID,
	}
	//select the item that matches the id
	db := dbSession.
		Find(&order)

	if db.Error != nil {
		log.Log(db.Error)
		return nil, db.Error
	}

	return &order, nil
}

//List retrieve all items
func (o *OrderItemRepository) List() ([]*models.OrderItem, error) {
	//Get a database session
	dbSession, err := persistence.GetDB()

	if err != nil {
		return nil, err
	}
	orders := []*models.OrderItem{}
	//retrieve all items
	db := dbSession.
		Find(&orders)

	if db.Error != nil {
		log.Log(db.Error)
		return nil, db.Error
	}

	return orders, nil
}

func (o *OrderItemRepository) FindLike(value string) ([]*models.OrderItem, error) {
	//Get a database session
	dbSession, err := persistence.GetDB()

	if err != nil {
		return nil, err
	}

	orderItems := []*models.OrderItem{}
	//select the item that matches the id
	db := dbSession.Model(&models.OrderItem{}).Where("product LIKE ?", value+"%").Find(&orderItems)

	if db.Error != nil {
		log.Log(db.Error)
		return nil, db.Error
	}

	return orderItems, nil
}

//FindByOrderID retrieve  items by orderID
func (o *OrderItemRepository) FindByOrderID(orderID uint64) ([]*models.OrderItem, error) {
	//Get a database session
	dbSession, err := persistence.GetDB()

	if err != nil {
		return nil, err
	}
	orderItems := []*models.OrderItem{}
	//retrieve all items
	db := dbSession.
		Where(&models.OrderItem{OrderID: orderID}).Find(&orderItems)

	if db.Error != nil {
		log.Log(db.Error)
		return nil, db.Error
	}

	return orderItems, nil
}
