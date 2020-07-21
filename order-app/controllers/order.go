package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ysfkel/order-app/services"
)

type OrderControlller struct {
	orderService *services.OrderService
}

func NewOrderController(orderService *services.OrderService) *OrderControlller {

	return &OrderControlller{
		orderService: orderService,
	}
}

type Response struct {
	Data  interface{}
	Error string
}

func (o *OrderControlller) List(c echo.Context) error {

	var startDate time.Time

	var endDate time.Time

	var offSet int

	var pageCount int

	search := strings.TrimSpace(c.QueryParam("search"))

	startDateStr := strings.TrimSpace(c.QueryParam("start_date"))

	endDateStr := strings.TrimSpace(c.QueryParam("end_date"))

	offSetStr := strings.TrimSpace(c.QueryParam("off_set"))

	pageCountStr := strings.TrimSpace(c.QueryParam("page_count"))

	var orders *services.ResultDTO

	var err error

	//vaidate inputs
	if startDateStr != "" && endDateStr == "" {
		return c.JSON(http.StatusBadRequest, &Response{
			Data:  nil,
			Error: "end_date cannot be empty",
		})
	}

	if startDateStr == "" && endDateStr != "" {
		return c.JSON(http.StatusBadRequest, &Response{
			Data:  nil,
			Error: "start_date cannot be empty",
		})

	}

	if startDateStr != "" {

		startDate, err = time.Parse(time.RFC3339, startDateStr)

		if err != nil {
			return c.JSON(http.StatusBadRequest, &Response{
				Data:  nil,
				Error: "invalid start_date",
			})
		}
	}

	if endDateStr != "" {

		endDate, err = time.Parse(time.RFC3339, endDateStr)

		if err != nil {
			return c.JSON(http.StatusBadRequest, &Response{
				Data:  nil,
				Error: "invalid end_date",
			})
		}
	}

	if offSetStr != "" {

		offSet, err = strconv.Atoi(offSetStr)

		if err != nil {
			return c.JSON(http.StatusBadRequest, &Response{
				Data:  nil,
				Error: "invalid pager off_set",
			})
		}
	}

	if pageCountStr != "" {

		pageCount, err = strconv.Atoi(pageCountStr)

		if err != nil {
			return c.JSON(http.StatusBadRequest, &Response{
				Data:  nil,
				Error: "invalid pager page_count",
			})
		}
	}

	//START SEARCH
	if search == "" && startDateStr == "" && endDateStr == "" {
		orders, err = o.orderService.List(offSet, pageCount)
	}

	if search != "" && startDateStr == "" && endDateStr == "" {
		orders, err = o.orderService.ListBySearchParam(search, offSet, pageCount)
	}

	//search by ListByDateRange if start_date and end_date are not empty
	if search == "" && startDateStr != "" && endDateStr != "" {
		orders, err = o.orderService.ListByDateRange(startDate, endDate, offSet, pageCount)
	}

	//search by ListBySearchParamAndDate if all parameters are present
	if search != "" && startDateStr != "" && endDateStr != "" {
		orders, err = o.orderService.ListBySearchParamAndDate(search, startDate, endDate, offSet, pageCount)
	}

	//HANDLE RESPONSE
	if err != nil {

		return c.JSON(http.StatusInternalServerError, &Response{
			Error: err.Error(),
		})

	}

	return c.JSON(http.StatusOK, &Response{
		Data: orders,
	})
}
