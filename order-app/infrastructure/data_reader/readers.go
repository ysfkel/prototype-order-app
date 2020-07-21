package data_reader

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ysfkel/order-app/domain/models"
)

const (
	base                  = "test_data"
	order_file            = base + "/postgres-orders.csv"
	oder_item_file        = base + "/postgres-order_items.csv"
	deliveries_file       = base + "/postgres-deliveries.csv"
	customers_file        = base + "/mongo-customers.csv"
	customer_company_file = base + "/mongo-customer_companies.csv"
)

func ReadOrders() ([]*models.Order, error) {

	file, err := openFile(order_file)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := getScanner(file)

	orders := []*models.Order{}

	scanner.Scan() //scans colums
	for scanner.Scan() {
		order := &models.Order{}

		fields := strings.Split(scanner.Text(), ",")

		id, err := strconv.ParseUint(fields[0], 10, 64)

		if err != nil {
			return nil, errors.New("error failed to parse oder id")
		}

		order.ID = id
		createdAt, err := time.Parse(time.RFC3339, fields[1])

		if err != nil {
			return nil, errors.New("error failed to parse created_at")
		}
		order.CreatedAt = createdAt

		order.Name = fields[2]

		order.CustomerID = fields[3]

		orders = append(orders, order)

	}

	return orders, nil
}

func ReadOrderItems() ([]*models.OrderItem, error) {

	file, err := openFile(oder_item_file)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := getScanner(file)

	orderItems := []*models.OrderItem{}

	scanner.Scan() //scans colums
	for scanner.Scan() {

		orderItem := &models.OrderItem{}

		fields := strings.Split(scanner.Text(), ",")

		id, err := strconv.ParseUint(fields[0], 10, 64)

		if err != nil {
			return nil, errors.New("error failed to parse  id")
		}

		orderItem.ID = id

		orderId, err := strconv.ParseUint(fields[1], 10, 64)

		if err != nil {
			return nil, errors.New("error failed to parse oder_id")
		}

		orderItem.OrderID = orderId
		price, err := strconv.ParseFloat(fields[2], 64)

		if err != nil {
			fmt.Println(err)
			return nil, errors.New("error failed to parse price")
		}

		orderItem.Price = price

		quantity, err := strconv.ParseUint(fields[3], 10, 64)

		if err != nil {
			return nil, errors.New("error failed to parse quantity")
		}

		orderItem.Quantity = quantity

		orderItem.Product = fields[4]

		orderItems = append(orderItems, orderItem)

	}

	return orderItems, nil
}

func ReadDeliveries() ([]*models.Delivery, error) {

	file, err := openFile(deliveries_file)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := getScanner(file)

	deliveries := []*models.Delivery{}

	scanner.Scan() //scans colums
	for scanner.Scan() {

		delivery := &models.Delivery{}

		fields := strings.Split(scanner.Text(), ",")

		id, err := strconv.ParseUint(fields[0], 10, 64)

		if err != nil {
			return nil, errors.New("error failed to parse  id")
		}

		delivery.ID = id

		orderItemId, err := strconv.ParseUint(fields[1], 10, 64)

		if err != nil {
			return nil, errors.New("error failed to parse oder_item_id")
		}

		delivery.OrderItemID = orderItemId
		quantity, err := strconv.ParseUint(fields[2], 10, 64)

		if err != nil {
			fmt.Println(err)
			return nil, errors.New("error failed to parse delivery_quantity")
		}

		delivery.DeliveredQuantity = quantity

		deliveries = append(deliveries, delivery)

	}

	return deliveries, nil
}

func ReadCustomers() ([]*models.Customer, error) {

	file, err := openFile(customers_file)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := getScanner(file)

	customers := []*models.Customer{}

	scanner.Scan() //scans colums
	for scanner.Scan() {

		customer := &models.Customer{}

		fields := strings.Split(scanner.Text(), ", ")
		customer.UserID = fields[0]

		customer.Login = fields[1]

		customer.Password = fields[2]

		customer.Name = fields[3]

		companyId, err := strconv.ParseUint(fields[4], 10, 64)

		if err != nil {
			return nil, errors.New("error failed to parse company id")
		}

		customer.CompanyID = companyId

		customer.CreditCards = fields[5]

		customers = append(customers, customer)

	}

	return customers, nil
}

func ReadCustomerCompany() ([]*models.CustomerCompany, error) {

	file, err := openFile(customer_company_file)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := getScanner(file)

	companies := []*models.CustomerCompany{}

	scanner.Scan() //scans colums
	for scanner.Scan() {

		company := &models.CustomerCompany{}

		fields := strings.Split(scanner.Text(), ",")

		id, err := strconv.ParseUint(fields[0], 10, 64)

		if err != nil {
			return nil, errors.New("error failed to parse company id")
		}

		company.CompanyID = id

		company.Name = fields[1]

		companies = append(companies, company)

	}

	return companies, nil
}
