package repository

import (
	mgm "github.com/Kamva/mgm"
	"github.com/Kamva/mgm/operator"

	"github.com/ysfkel/order-app/domain/models"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	_ = mgm.SetDefaultConfig(nil, "storedb", options.Client().ApplyURI("mongodb://mongo:27017"))
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}

type CustomerRepository struct {
}

func (c *CustomerRepository) Create(customer *models.Customer) (*models.Customer, error) {

	err := mgm.Coll(customer).Create(customer)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (c *CustomerRepository) Get(id string) (*models.Customer, error) {

	customer := &models.Customer{}

	customerCollection := mgm.Coll(customer)

	err := customerCollection.FindByID(id, customer)

	if err != nil {
		return nil, err
	}

	return customer, nil

}

func (c *CustomerRepository) FindByUserID(userID string) (*models.Customer, error) {

	customerCollection := mgm.Coll(&models.Customer{})

	customer := &models.Customer{}

	result := customerCollection.FindOne(mgm.Ctx(), bson.M{"userid": bson.M{operator.Eq: userID}})

	err := result.Decode(customer)

	if err != nil {
		return nil, err
	}

	return customer, nil

}

func (c *CustomerRepository) List() ([]*models.Customer, error) {

	customerCollection := mgm.Coll(&models.Customer{})

	customers := []*models.Customer{}

	err := customerCollection.SimpleFind(&customers, bson.M{})

	if err != nil {
		return nil, err
	}

	return customers, nil

}
