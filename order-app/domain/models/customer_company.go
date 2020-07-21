package models

import "github.com/Kamva/mgm"

type CustomerCompany struct {
	mgm.DefaultModel `bson:",inline"`
	CompanyID        uint64 `json:"company_id" json:"company_id,omitempty" csv:"company_id"`
	Name             string `json:"company_name" csv:"company_name"`
}

type ICustomerCompanyRepository interface {
	Create(m *CustomerCompany) (*CustomerCompany, error)
	Get(id uint64) (*CustomerCompany, error)
	List() ([]*CustomerCompany, error)
	FindByCompanyID(id uint64) (*CustomerCompany, error)
}
