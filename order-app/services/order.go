package services

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/ysfkel/order-app/domain/models"
)

type OrderDTO struct {
	ID              uint64    `json:"id"`
	OrderName       string    `json:"order_name"`
	CompanyName     string    `json:"company_name"`
	CustomerName    string    `json:"customer_name"`
	CreatedAt       time.Time `json:"created_at"`
	DeliveredAmount float64   `json:"delivered_amount"`
	TotalAmount     float64   `json:"total_amount"`
}

type ResultDTO struct {
	Orders           []*OrderDTO
	GrandTotalAmount float64
}

//OrderService OrderService type
type OrderService struct {
	orderRepository     models.IOrderRepository
	orderItemRepository models.IOrderItemRepository
	deliveryRepository  models.IDeliveryRepository
	customerRepository  models.ICustomerRepository
	companyRepository   models.ICustomerCompanyRepository
}

//NewOrderService creates a new OrderService
func NewOrderService(
	orderRepository models.IOrderRepository,
	orderItemRepository models.IOrderItemRepository,
	deliveryRepository models.IDeliveryRepository,
	customerRepository models.ICustomerRepository,
	companyRepository models.ICustomerCompanyRepository,
) *OrderService {
	return &OrderService{
		orderRepository:     orderRepository,
		orderItemRepository: orderItemRepository,
		deliveryRepository:  deliveryRepository,
		customerRepository:  customerRepository,
		companyRepository:   companyRepository,
	}
}

const errorMessage = "error fetching orders"

func (s *OrderService) List(offSet int, pageCount int) (*ResultDTO, error) {

	orders, err := s.orderRepository.List(offSet, pageCount)

	if err != nil {
		return nil, errors.New(errorMessage)
	}

	if len(orders) == 0 {
		return &ResultDTO{
			Orders: []*OrderDTO{},
		}, nil
	}

	ordersDTO, err := s.createOrderDTO(orders)

	if err != nil {
		return &ResultDTO{
			Orders: []*OrderDTO{},
		}, nil
	}

	return ordersDTO, err
}

func (s *OrderService) ListBySearchParam(search string, offSet int, pageCount int) (*ResultDTO, error) {

	var orders []*models.Order
	var err error

	search = strings.TrimSpace(search)

	if search == "" {
		return nil, errors.New("error:: query parameter cannot be empty")
	}

	orders, err = s.orderRepository.FindLike(search, offSet, pageCount)

	if err != nil {
		return nil, errors.New(errorMessage)
	}

	if len(orders) == 0 {

		orders, err = s.orderRepository.ListByProductName(search, offSet, pageCount)

		if err != nil {
			return nil, errors.New(errorMessage)
		}
	}

	if len(orders) == 0 {
		return &ResultDTO{
			Orders: []*OrderDTO{},
		}, nil
	}

	ordersDTO, err := s.createOrderDTO(orders)

	if err != nil {
		return &ResultDTO{
			Orders: []*OrderDTO{},
		}, nil
	}

	return ordersDTO, err
}

func (s *OrderService) ListByDateRange(startDate time.Time, endDate time.Time, offSet int, pageCount int) (*ResultDTO, error) {

	orders, err := s.orderRepository.ListByDateRange(startDate, endDate, offSet, pageCount)

	if err != nil {
		return nil, errors.New(errorMessage)
	}

	if len(orders) == 0 {
		return &ResultDTO{
			Orders: []*OrderDTO{},
		}, nil
	}

	ordersDTO, err := s.createOrderDTO(orders)

	if err != nil {
		return &ResultDTO{
			Orders: []*OrderDTO{},
		}, nil
	}

	return ordersDTO, err
}

func (s *OrderService) ListBySearchParamAndDate(search string, startDate time.Time, endDate time.Time, offSet int, pageCount int) (*ResultDTO, error) {

	var orders []*models.Order
	var err error

	search = strings.TrimSpace(search)

	if search == "" {
		return nil, errors.New("error:: query parameter cannot be empty")
	}

	orders, err = s.orderRepository.FindLikeAndDate(search, startDate, endDate, offSet, pageCount)

	if err != nil {
		return nil, errors.New(errorMessage)
	}

	if len(orders) == 0 {

		orders, err = s.orderRepository.ListByProductNameAndDate(search, startDate, endDate, offSet, pageCount)

		if err != nil {
			return nil, errors.New(errorMessage)
		}
	}

	if len(orders) == 0 {
		return &ResultDTO{
			Orders: []*OrderDTO{},
		}, nil
	}

	ordersDTO, err := s.createOrderDTO(orders)

	if err != nil {
		return &ResultDTO{
			Orders: []*OrderDTO{},
		}, nil
	}

	return ordersDTO, err
}

func (s *OrderService) createOrderDTO(orders []*models.Order) (*ResultDTO, error) {

	ordersDTO := []*OrderDTO{}

	grantTotal := 0.

	for _, order := range orders {

		orderDTO := &OrderDTO{
			ID:        order.ID,
			OrderName: order.Name,
			CreatedAt: order.CreatedAt,
		}

		//calculate total
		orderItems, err := s.orderItemRepository.FindByOrderID(order.ID)

		if err != nil && !gorm.IsRecordNotFoundError(err) {
			log.Println("error fetching orderItems ", err)
		}

		/**
		if orderItems is not empty
		 - calculate total and delivered amounts
		*/
		if len(orderItems) > 0 {
			totalAmount, deliveryTotal := calculateAmounts(orderItems, s.deliveryRepository.FindByOrderItemID)
			orderDTO.TotalAmount = totalAmount
			orderDTO.DeliveredAmount = deliveryTotal
			grantTotal += totalAmount
		}

		//Get customer
		customer, err := s.customerRepository.FindByUserID(order.CustomerID)

		//if customer is empty log
		//we will not return because if customer is empty
		//the ui will show remaining info and show customer name as empty
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			log.Println("error fetching customer ", err)
		}

		if customer != nil {
			orderDTO.CustomerName = customer.Name

			company, err := s.companyRepository.FindByCompanyID(customer.CompanyID)

			if err != nil && !gorm.IsRecordNotFoundError(err) {
				log.Println("error fetching company ", err)
			}

			if company != nil {
				orderDTO.CompanyName = company.Name
			}
		}
		ordersDTO = append(ordersDTO, orderDTO)
	}

	result := &ResultDTO{
		Orders:           ordersDTO,
		GrandTotalAmount: grantTotal,
	}

	return result, nil

}

func calculateAmounts(orderItems []*models.OrderItem, getDelivery func(orderItemID uint64) (*models.Delivery, error)) (totalAmount float64, deliveredAmount float64) {

	for _, item := range orderItems {
		totalAmount += item.Price

		delivery, err := getDelivery(item.ID)

		if err != nil && !gorm.IsRecordNotFoundError(err) {
			log.Println("error fetching delivery ", err)
		}

		if delivery != nil {
			deliveredAmount += float64(delivery.DeliveredQuantity) * item.Price
		}

	}
	return totalAmount, deliveredAmount
}
