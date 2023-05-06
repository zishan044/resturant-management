package routes

import (
	controller "golang-resturant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("orderItems/", controller.GetOrderItems())
	incomingRoutes.GET("orderItems/:orderItem_id", controller.GetOrderItem())
	incomingRoutes.POST("orderItems", controller.CreateOrderItem())
	incomingRoutes.PATCH("orderItems/:orderItem_id", controller.UpdateOrderItem())
}
