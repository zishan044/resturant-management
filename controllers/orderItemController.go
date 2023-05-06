package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrderItems() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func UpdateOrderItem() gin.HandlerFunc {

}

func itemsByOrder(id string) (orderItems []primitive.M, err error) {

}
