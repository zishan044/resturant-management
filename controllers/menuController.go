package controller

import (
	"context"
	"golang-resturant-management/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		result, err := menuCollection.Find(ctx, bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var menus []bson.M
		err = result.All(ctx, menus)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, menus)
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var menu models.Menu
		meniId := c.Param("menu_id")

		err := menuCollection.FindOne(ctx, bson.M{"menu_id": meniId}).Decode(&menu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the menu item"})
			return
		}
		c.JSON(http.StatusOK, menu)

	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		var menu models.Menu

		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if vallidationErr := validate.Struct(menu); vallidationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": vallidationErr.Error()})
			return
		}

		menu.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		menu.ID = primitive.NewObjectID()
		menu.Menu_id = menu.ID.Hex()

		result, err := menuCollection.InsertOne(ctx, menu)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while inserting menu item in database"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, result)
		defer cancel()

	}
}

func inTimeSapn(start, end, check time.Time) bool {
	return start.After(check) && end.After(start)
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var menu models.Menu

		if err := c.BindJSON(&menu); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		menuId := c.Param("menu_id")
		filter := bson.M{"menu_id": menuId}

		var updateObj primitive.D

		if menu.Start_date != nil && menu.End_date != nil && !inTimeSapn(*menu.Start_date, *menu.End_date, time.Now()) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "please set the dates  correctly"})
			return
		}

		updateObj = append(updateObj, bson.E{"start_date", menu.Start_date})
		updateObj = append(updateObj, bson.E{"end_date", menu.End_date})

		if menu.Name != "" {
			updateObj = append(updateObj, bson.E{"name", menu.Name})
		}

		if menu.Category != "" {
			updateObj = append(updateObj, bson.E{"category", menu.Category})
		}

		menu.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", menu.Updated_at})

		upsert := true

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := menuCollection.UpdateOne(ctx, filter, bson.D{{"$set", updateObj}}, &opt)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, result)

	}
}
