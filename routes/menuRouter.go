package routes

import (
	controller "golang-resturant-management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("menus/", controller.GetMenus()) //get all menus
	incomingRoutes.GET("menus/:menu_id", controller.GetMenu())
	incomingRoutes.POST("menus", controller.CreateMenu())
	incomingRoutes.PATCH("menus/:menu_id", controller.UpdateMenu())
}
