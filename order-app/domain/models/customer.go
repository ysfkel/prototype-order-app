package models

import "github.com/Kamva/mgm"

type Customer struct {
	mgm.DefaultModel `bson:",inline"`
	UserID           string `json:"user_id" json:"user_id,omitempty" csv:"user_id"`
	Name             string `json:"name" csv:"name"`
	Login            string `json:"login" csv:"login"`
	Password         string `json:"password" csv:"password"`
	CompanyID        uint64 `json:"company_id" csv:"company_id"`
	CreditCards      string `json:"credit_cards" csv:"credit_cards"`
}

type ICustomerRepository interface {
	Create(m *Customer) (*Customer, error)
	Get(id string) (*Customer, error)
	List() ([]*Customer, error)
	FindByUserID(userID string) (*Customer, error)
}
