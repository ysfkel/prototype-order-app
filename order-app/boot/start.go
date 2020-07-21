package boot

import (
	"fmt"

	"github.com/go-log/log"
	"github.com/ysfkel/order-app/api"
	"github.com/ysfkel/order-app/infrastructure/persistence/migration"
	"github.com/ysfkel/order-app/infrastructure/persistence/repository"
	"github.com/ysfkel/order-app/services"

	controllers "github.com/ysfkel/order-app/controllers"
	"github.com/ysfkel/order-app/infrastructure/persistence/database/postgresql"
)

//Start starts the application
func Start() error {

	err := runDatabaseMigrations()

	if err != nil {
		return err
	}

	err = startWebServer()

	if err != nil {
		return err
	}

	return err
}

//runDatabaseMigrations runs database migrations
func runDatabaseMigrations() error {

	fmt.Println("..running database migrations")
	//Open db connection makes singleton
	_, err := postgresql.GetDB()

	if err != nil {
		return err
	}
	// Migrating models to the database
	err = migration.MigrateDatabase()

	if err != nil {
		return err
	}

	fmt.Println("..database migrations completed successfully")

	return nil
}

func startWebServer() error {

	log.Log("..registering http routes")

	//initialize order service
	ordersService := services.NewOrderService(
		repository.NewOrderRepository(),
		repository.NewOrderItemRepository(),
		repository.NewDeliveryRepository(),
		repository.NewCustomerRepository(),
		repository.NewCustomerCompanyRepository())

	//initialize order controller
	ordersController := controllers.NewOrderController(ordersService)
	//intialize router
	router := api.NewRouter(ordersController)
	//register routes
	routes := router.RegisterRoutes()
	log.Log("..http routes registered successfully")
	err := routes.Start(":5000")

	return err
}
