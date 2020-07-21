package repository

import (
	"fmt"
	"time"

	"github.com/micro/go-log"

	"github.com/ysfkel/order-app/domain/models"
	persistence "github.com/ysfkel/order-app/infrastructure/persistence/database/postgresql"
)

//OrderRepository operation
type OrderRepository struct {
}

//NewOrderRepository instantiates and returns OrderRepository
func NewOrderRepository() models.IOrderRepository {
	return &OrderRepository{}
}

//Create OrderRepository model
func (o *OrderRepository) Create(m *models.Order) (*models.Order, error) {
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
func (o *OrderRepository) Get(orderID uint64) (*models.Order, error) {
	//Get a database session
	dbSession, err := persistence.GetDB()

	if err != nil {
		return nil, err
	}

	order := models.Order{
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

func (o *OrderRepository) GetCount(queryModel *models.Order) (int, error) {

	dbSession, err := persistence.GetDB()

	if err != nil {
		return 0, err
	}

	var count int
	//select the item that matches the id
	db := dbSession.Model(&models.Order{}).Where(queryModel).Count(&count)

	if db.Error != nil {
		log.Log(db.Error)
		return 0, db.Error
	}

	return count, nil

}

//List retrieve all items
func (o *OrderRepository) List(offSet int, limit int) ([]*models.Order, error) {
	//Get a database session
	dbSession, err := persistence.GetDB()

	if err != nil {
		return nil, err
	}

	orders := []*models.Order{}

	fmt.Println(" off ", offSet, " limit ", limit)

	dbSession = dbSession.Offset(offSet)

	if limit > 0 {
		dbSession = dbSession.Limit(limit)
	}

	//retrieve all items
	db := dbSession.
		Find(&orders)

	if db.Error != nil {
		log.Log(db.Error)
		return nil, db.Error
	}

	return orders, nil
}

func (o *OrderRepository) FindLike(value string, offSet int, limit int) ([]*models.Order, error) {
	//Get a database session
	dbSession, err := persistence.GetDB()

	if err != nil {
		return nil, err
	}

	dbSession = dbSession.Offset(offSet)

	if limit > 0 {
		dbSession = dbSession.Limit(limit)
	}

	orders := []*models.Order{}
	//select the item that matches the id
	db := dbSession.Model(&models.Order{}).Where("name LIKE ?", value+"%").Find(&orders)

	if db.Error != nil {
		log.Log(db.Error)
		return nil, db.Error
	}

	return orders, nil
}

func (o *OrderRepository) FindLikeAndDate(value string, startDate time.Time, endDate time.Time, offSet int, limit int) ([]*models.Order, error) {
	//Get a database session
	dbSession, err := persistence.GetDB()

	if err != nil {
		return nil, err
	}

	dbSession = dbSession.Offset(offSet)

	if limit > 0 {
		dbSession = dbSession.Limit(limit)
	}

	orders := []*models.Order{}
	//select the item that matches the id
	db := dbSession.Model(&models.Order{}).
		Where("name LIKE ?", value+"%").
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Find(&orders)

	if db.Error != nil {
		log.Log(db.Error)
		return nil, db.Error
	}

	return orders, nil
}

//List retrieve all items
func (o *OrderRepository) ListByProductName(productName string, offSet int, limit int) ([]*models.Order, error) {
	//Get a database session
	dbSession, err := persistence.GetDB()

	if err != nil {
		return nil, err
	}

	dbSession = dbSession.Offset(offSet)

	if limit > 0 {
		dbSession = dbSession.Limit(limit)
	}

	orders := []*models.Order{}
	//retrieve all items

	db := dbSession.Where("id IN (?)", dbSession.Table("order_items").
		Where("product LIKE ?", productName+"%").
		Select("order_id").
		SubQuery()).
		Find(&orders)

	if db.Error != nil {
		log.Log(db.Error)
		return nil, db.Error
	}

	return orders, nil
}

//List retrieve all items
func (o *OrderRepository) ListByProductNameAndDate(productName string, startDate time.Time, endDate time.Time, offSet int, limit int) ([]*models.Order, error) {
	//Get a database session
	dbSession, err := persistence.GetDB()

	if err != nil {
		return nil, err
	}

	dbSession = dbSession.Offset(offSet)

	if limit > 0 {
		dbSession = dbSession.Limit(limit)
	}

	orders := []*models.Order{}
	//retrieve all items

	db := dbSession.Where("id IN (?)", dbSession.Table("order_items").
		Where("product LIKE ?", productName+"%").
		Select("order_id").
		SubQuery()).
		Where("created_at BETWEEN ? AND ?", startDate, endDate).
		Find(&orders)

	if db.Error != nil {
		log.Log(db.Error)
		return nil, db.Error
	}

	return orders, nil
}

func (o *OrderRepository) ListByDateRange(startDate time.Time, endDate time.Time, offSet int, limit int) ([]*models.Order, error) {

	dbSession, err := persistence.GetDB()

	if err != nil {
		return nil, err
	}

	dbSession = dbSession.Offset(offSet)

	if limit > 0 {
		dbSession = dbSession.Limit(limit)
	}

	orders := []*models.Order{}
	//retrieve all items
	db := dbSession.Where("created_at BETWEEN ? AND ?", startDate, endDate).Find(&orders)

	if db.Error != nil {
		log.Log(db.Error)
		return nil, db.Error
	}

	return orders, nil
}
