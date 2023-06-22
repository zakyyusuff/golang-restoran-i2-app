package controller

import (
	"context"
	"fmt"
	"golang-restaurant-management/database"
	"golang-restaurant-management/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

var menuCol *mongo.Collection = database.OpenCollection(database.Client, "menu")
var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing order items"})
		}
		var allOrders []bson.M
		if err = result.All(ctx, &allOrders); err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, allOrders)
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		orderId := c.Param("order_id")
		var order models.Order

		err := orderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while fetching the orders"})
		}
		c.JSON(http.StatusOK, order)
	}
}

// func CreateOrder() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var table models.Table
// 		var order models.Order

// 		if err := c.BindJSON(&order); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		validationErr := validate.Struct(order)

// 		if validationErr != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
// 			return
// 		}

// 		if order.Table_id != nil {
// 			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
// 			defer cancel()
// 			if err != nil {
// 				msg := fmt.Sprintf("message:Table was not found")
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 				return
// 			}
// 		}

// 		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
// 		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

// 		order.ID = primitive.NewObjectID()
// 		order.Order_id = order.ID.Hex()

// 		result, insertErr := orderCollection.InsertOne(ctx, order)

// 		if insertErr != nil {
// 			msg := fmt.Sprintf("order item was not created")
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 			return
// 		}

// 		defer cancel()
// 		c.JSON(http.StatusOK, result)
// 	}
// }

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var table models.Table
		var order models.Order

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(order)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		tableCollection := database.OpenCollection(database.Client, "table")

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second) // Tambahkan ini

		if order.Table_id != nil {
			err := tableCollection.FindOne(ctx, bson.M{"table_id": *order.Table_id}).Decode(&table)
			if err != nil {
				msg := fmt.Sprintf("message:Table was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				cancel() // Batalkan konteks jika terjadi kesalahan
				return
			}
		}

		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		order.ID = primitive.NewObjectID()
		order.Order_id = order.ID.Hex()

		result, insertErr := orderCollection.InsertOne(ctx, order)

		if insertErr != nil {
			msg := fmt.Sprintf("order item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			cancel() // Batalkan konteks jika terjadi kesalahan
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var table models.Table
		var order models.Order

		var updateObj primitive.D

		orderId := c.Param("order_id")
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if order.Table_id != nil {
			err := menuCollection.FindOne(ctx, bson.M{"tabled_id": order.Table_id}).Decode(&table)
			defer cancel()
			if err != nil {
				msg := fmt.Sprintf("message:Menu was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
			updateObj = append(updateObj, bson.E{"menu", order.Table_id})
		}

		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", order.Updated_at})

		upsert := true

		filter := bson.M{"order_id": orderId}
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := orderCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opt,
		)

		if err != nil {
			msg := fmt.Sprintf("order item update failed")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}

// //////////////////////// delete order
func DeleteOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID := c.Param("order_id")

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		// Create a filter to match the order_id
		filter := bson.M{"order_id": orderID}

		// Delete the order
		result, err := orderCollection.DeleteOne(ctx, filter)
		if err != nil {
			log.Println("Error deleting order:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "terjadi kesalahan saat menghapus pesanan"})
			return
		}

		// Check the result to determine if the order was deleted
		if result.DeletedCount == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "pesanan tidak ditemukan"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "pesanan berhasil dihapus"})
	}
}

func OrderItemOrderCreator(order models.Order) string {

	order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.ID = primitive.NewObjectID()
	order.Order_id = order.ID.Hex()

	orderCollection.InsertOne(ctx, order)
	defer cancel()

	return order.Order_id
}

/////////////////////////////////////////////////////////////////// coba 1

// package controller

// import (
// 	"context"
// 	"fmt"
// 	"golang-restaurant-management/database"
// 	"golang-restaurant-management/models"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"gopkg.in/go-playground/validator.v9"
// )

// var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// // var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")
// var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")
// var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")

// var validate *validator.Validate

// func init() {
// 	validate = validator.New()
// }

// func GetOrders() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

// 		result, err := orderCollection.Find(context.TODO(), bson.M{})
// 		defer cancel()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing order items"})
// 			return
// 		}

// 		var allOrders []models.Order
// 		if err = result.All(ctx, &allOrders); err != nil {
// 			log.Fatal(err)
// 			return
// 		}

// 		c.JSON(http.StatusOK, allOrders)
// 	}
// }

// func GetOrder() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
// 		orderID := c.Param("order_id")
// 		var order models.Order

// 		err := orderCollection.FindOne(ctx, bson.M{"order_id": orderID}).Decode(&order)
// 		defer cancel()
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while fetching the order"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, order)
// 	}
// }

// func CreateOrder() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var table models.Table
// 		var order models.Order

// 		if err := c.BindJSON(&order); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		validationErr := validate.Struct(order)
// 		if validationErr != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
// 			return
// 		}

// 		if order.TableID != nil {
// 			err := tableCollection.FindOne(ctx, bson.M{"table_id": *order.TableID}).Decode(&table)
// 			if err != nil {
// 				msg := fmt.Sprintf("Table was not found")
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 				return
// 			}
// 		}

// 		order.CreatedAt = time.Now()
// 		order.UpdatedAt = time.Now()
// 		order.ID = primitive.NewObjectID()
// 		order.OrderID = order.ID.Hex()

// 		result, insertErr := orderCollection.InsertOne(ctx, order)
// 		if insertErr != nil {
// 			msg := fmt.Sprintf("Order item was not created")
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 			return
// 		}

// 		c.JSON(http.StatusOK, result)
// 	}
// }

// func UpdateOrder() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var table models.Table
// 		var order models.Order

// 		orderID := c.Param("order_id")
// 		if err := c.BindJSON(&order); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if order.TableID != nil {
// 			err := tableCollection.FindOne(ctx, bson.M{"table_id": *order.TableID}).Decode(&table)
// 			if err != nil {
// 				msg := fmt.Sprintf("Table was not found")
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 				return
// 			}
// 		}

// 		order.UpdatedAt = time.Now()

// 		filter := bson.M{"order_id": orderID}
// 		update := bson.M{"$set": bson.M{"updated_at": order.UpdatedAt}}
// 		opts := options.Update().SetUpsert(true)

// 		result, err := orderCollection.UpdateOne(ctx, filter, update, opts)
// 		if err != nil {
// 			msg := fmt.Sprintf("Order item update failed")
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
// 			return
// 		}

// 		c.JSON(http.StatusOK, result)
// 	}
// }

// func DeleteOrder() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		orderID := c.Param("order_id")

// 		filter := bson.M{"order_id": orderID}
// 		result, err := orderCollection.DeleteOne(ctx, filter)
// 		if err != nil {
// 			log.Println("Error deleting order:", err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while deleting the order"})
// 			return
// 		}

// 		if result.DeletedCount == 0 {
// 			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
// 	}
// }

// func OrderItemOrderCreator(order models.Order) string {
// 	order.CreatedAt = time.Now()
// 	order.UpdatedAt = time.Now()
// 	order.ID = primitive.NewObjectID()
// 	order.OrderID = order.ID.Hex()

// 	orderCollection.InsertOne(ctx, order)
// 	defer cancel()

// 	return order.OrderID
// }
