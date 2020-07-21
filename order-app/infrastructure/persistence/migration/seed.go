package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/ysfkel/order-app/domain/models"
)

//Seeder seeder type
type Seeder struct {
	orderRepository     models.IOrderRepository
	orderItemRepository models.IOrderItemRepository
	deliveryRepository  models.IDeliveryRepository
	customerRepository  models.ICustomerRepository
	companyRepository   models.ICustomerCompanyRepository
}

//NewSeeder creates a new Seeder
func NewSeeder(
	orderRepository models.IOrderRepository,
	orderItemRepository models.IOrderItemRepository,
	deliveryRepository models.IDeliveryRepository,
	customerRepository models.ICustomerRepository,
	companyRepository models.ICustomerCompanyRepository,
) *Seeder {
	return &Seeder{
		orderRepository:     orderRepository,
		orderItemRepository: orderItemRepository,
		deliveryRepository:  deliveryRepository,
		customerRepository:  customerRepository,
		companyRepository:   companyRepository,
	}
}

//SeedOrders adds orders to db
func (s *Seeder) SeedOrders(orders []*models.Order) error {
	for _, order := range orders {

		if _, err := s.orderRepository.Get(order.ID); gorm.IsRecordNotFoundError(err) {
			_, err := s.orderRepository.Create(order)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}

	return nil
}

//SeedOrderItems adds orderitems to db
func (s *Seeder) SeedOrderItems(orderItems []*models.OrderItem) error {
	for _, orderItem := range orderItems {

		if _, err := s.orderItemRepository.Get(orderItem.ID); gorm.IsRecordNotFoundError(err) {
			_, err := s.orderItemRepository.Create(orderItem)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}

	return nil
}

//SeedDeliveries adds deliveries to db
func (s *Seeder) SeedDeliveries(deliveries []*models.Delivery) error {
	for _, delivery := range deliveries {

		if _, err := s.deliveryRepository.Get(delivery.ID); gorm.IsRecordNotFoundError(err) {
			_, err := s.deliveryRepository.Create(delivery)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}

	return nil
}

//SeedCustomers adds customers to db
func (s *Seeder) SeedCustomers(customers []*models.Customer) error {
	for _, customer := range customers {

		if cst, err := s.customerRepository.FindByUserID(customer.UserID); cst == nil {
			_, err := s.customerRepository.Create(customer)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}

	return nil
}

//SeedCusyomers adds customers to db
func (s *Seeder) SeedCustomersCompany(companies []*models.CustomerCompany) error {
	for _, company := range companies {

		if cst, err := s.companyRepository.FindByCompanyID(company.CompanyID); cst == nil {
			_, err := s.companyRepository.Create(company)
			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}

	return nil
}

//Seed runs all seed methods
func (s *Seeder) Seed(
	orders []*models.Order,
	orderItems []*models.OrderItem,
	deliveries []*models.Delivery,
	customers []*models.Customer,
	companies []*models.CustomerCompany) error {

	err := s.SeedOrders(orders)

	if err != nil {
		return err
	}

	err = s.SeedOrderItems(orderItems)

	if err != nil {
		return err
	}

	err = s.SeedDeliveries(deliveries)

	if err != nil {
		return err
	}

	err = s.SeedCustomers(customers)

	if err != nil {
		return err
	}

	err = s.SeedCustomersCompany(companies)

	if err != nil {
		return err
	}

	return nil
}



