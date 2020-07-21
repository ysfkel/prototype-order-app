package repository

import (
	"github.com/micro/go-log"
	"github.com/ysfkel/order-app/domain/models"
	persistence "github.com/ysfkel/order-app/infrastructure/persistence/database/postgresql"
)

//DeliveryRepository operation
type DeliveryRepository struct {
}

//NewDeliveryRepository instantiates and returns DeliveryRepository
func NewDeliveryRepository() models.IDeliveryRepository {
	return &DeliveryRepository{}
}

//Create Delivery model
func (o *DeliveryRepository) Create(m *models.Delivery) (*models.Delivery, error) {
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
func (o *DeliveryRepository) Get(orderID uint64) (*models.Delivery, error) {
	//Get a database session
	dbSession, err := persistence.GetDB()

	if err != nil {
		return nil, err
	}

	order := models.Delivery{
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
func (o *DeliveryRepository) List() ([]*models.Delivery, error) {
	//Get a database session
	dbSession, err := persistence.GetDB()

	if err != nil {
		return nil, err
	}

	orders := []*models.Delivery{}
	//retrieve all items
	db := dbSession.
		Find(&orders)

	if db.Error != nil {
		log.Log(db.Error)
		return nil, db.Error
	}

	return orders, nil
}

//FindByOrderItemID retrieve  items by orderItemID
func (o *DeliveryRepository) FindByOrderItemID(orderItemID uint64) (*models.Delivery, error) {
	//Get a database session
	dbSession, err := persistence.GetDB()

	if err != nil {
		return nil, err
	}
	delivery := &models.Delivery{}
	//retrieve all items
	db := dbSession.
		Where(&models.Delivery{OrderItemID: orderItemID}).Find(delivery)

	if db.Error != nil {
		log.Log(db.Error)
		return nil, db.Error
	}

	return delivery, nil
}
