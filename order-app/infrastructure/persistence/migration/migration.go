package migration

import (
	"errors"
	"fmt"

	"github.com/ysfkel/order-app/domain/models"
	"github.com/ysfkel/order-app/infrastructure/data_reader"
	"github.com/ysfkel/order-app/infrastructure/persistence/database/postgresql"
	"github.com/ysfkel/order-app/infrastructure/persistence/repository"
)

//MigrateDatabase performs database migrations
func MigrateDatabase() error {

	//Get a database session
	db, err := postgresql.GetDB()
	if err != nil {
		fmt.Printf("[Error getting DB connection]: %v\n", err)
		return err
	}
	//Migrate the schema
	db.AutoMigrate(
		&models.Order{},
		&models.OrderItem{},
		&models.Delivery{},
	)

	if db.Error != nil {
		message := fmt.Sprintf("[database migrations error]: %v", db.Error)
		return errors.New(message)
	}

	err = seedDatabase()

	if err != nil {
		return err
	}

	return nil

}

func seedDatabase() error {

	orders, err := data_reader.ReadOrders()

	if err != nil {
		return err
	}

	orderItems, err := data_reader.ReadOrderItems()

	if err != nil {
		return err
	}

	deliveries, err := data_reader.ReadDeliveries()

	if err != nil {
		return err
	}

	customers, err := data_reader.ReadCustomers()

	if err != nil {
		return err
	}

	companies, err := data_reader.ReadCustomerCompany()

	if err != nil {
		return err
	}

	dbseeder := NewSeeder(repository.NewOrderRepository(), repository.NewOrderItemRepository(),
		repository.NewDeliveryRepository(), repository.NewCustomerRepository(),
		repository.NewCustomerCompanyRepository())

	err = dbseeder.Seed(orders, orderItems, deliveries, customers, companies)

	if err != nil {
		return err
	}

	return nil
}
