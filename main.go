package main

import (
	"golang-resturant-management/middlewares"
	"golang-resturant-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	port := "8000"

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middlewares.Authentication())

	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderItemRoutes(router)
	routes.OrderRoutes(router)
	routes.TableRoutes(router)

	router.Run(":" + port)

}
