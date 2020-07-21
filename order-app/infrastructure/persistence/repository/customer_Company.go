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

func NewCustomerCompanyRepository() *CustomerCompanyRepository {
	return &CustomerCompanyRepository{}
}

type CustomerCompanyRepository struct {
}

func (c *CustomerCompanyRepository) Create(company *models.CustomerCompany) (*models.CustomerCompany, error) {

	err := mgm.Coll(company).Create(company)

	if err != nil {
		return nil, err
	}

	return company, nil
}

func (c *CustomerCompanyRepository) Get(id uint64) (*models.CustomerCompany, error) {

	company := &models.CustomerCompany{}

	collection := mgm.Coll(company)

	err := collection.FindByID(id, company)

	if err != nil {
		return nil, err
	}

	return company, nil

}

func (c *CustomerCompanyRepository) FindByCompanyID(userID uint64) (*models.CustomerCompany, error) {

	collection := mgm.Coll(&models.CustomerCompany{})

	company := &models.CustomerCompany{}

	result := collection.FindOne(mgm.Ctx(), bson.M{"companyid": bson.M{operator.Eq: userID}})

	err := result.Decode(company)

	if err != nil {
		return nil, err
	}

	return company, nil

}

func (c *CustomerCompanyRepository) List() ([]*models.CustomerCompany, error) {

	collection := mgm.Coll(&models.CustomerCompany{})

	companies := []*models.CustomerCompany{}

	err := collection.SimpleFind(&companies, bson.M{})

	if err != nil {
		return nil, err
	}

	return companies, nil

}
