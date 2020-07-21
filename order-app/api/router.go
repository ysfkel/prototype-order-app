package api

import (
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	controllers "github.com/ysfkel/order-app/controllers"
)

type Router struct {
	HTTP             *echo.Echo
	ordersController *controllers.OrderControlller
}

func NewRouter(ordersController *controllers.OrderControlller) *Router {

	router := &Router{
		HTTP:             echo.New(),
		ordersController: ordersController,
	}

	router.registerMiddlewares()

	return router
}

func (r *Router) RegisterRoutes() *echo.Echo {

	v1API := r.HTTP.Group("/v1/api")
	ordersGroup := v1API.Group("/orders")
	ordersGroup.GET("", r.ordersController.List)
	return r.HTTP
}

//registerMiddlewares registeres global echo middlewares
func (r *Router) registerMiddlewares() {
	// Log HTTP requests
	r.HTTP.Use(middleware.Logger())
	// Recover from panics
	r.HTTP.Use(middleware.Recover())
	r.HTTP.Use(middleware.Secure())

	// Enable CORS
	r.HTTP.Use(middleware.CORSWithConfig(middleware.CORSConfig{

		AllowOrigins: []string{"http://client:6000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAuthorization},
	},
	))
}
