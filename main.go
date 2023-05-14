package main

import (
	"os"

	"golang-restaurant-management/database"

	middleware "golang-restaurant-management/middleware"
	routes "golang-restaurant-management/routes"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	os.Setenv("restaurant", "mongodb+srv://zakymuhammadyusuf:zaky123@zaky.oy6yt60.mongodb.net/")
	// url := os.Getenv("restaurant")
	// clientOptions := options.Client().ApplyURI(url)
	// client, err := mongo.Connect(context.TODO(), clientOptions)
	// port := os.Getenv(clientOptions)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}

// //////////////////////////////////////////////////////////////////////////////////////////
// func main() {
// 	port := os.Getenv("PORT")

// 	if port == "" {
// 		port = "8000"
// 	}

// 	router := gin.New()
// 	router.Use(gin.Logger())
// 	routes.UserRoutes(router)

// 	router.Use(middleware.Authentication())

// 	// API-2
// 	router.GET("/api-1", func(c *gin.Context) {

// 		c.JSON(200, gin.H{"success": "Access granted for api-1"})

// 	})

// 	// API-1
// 	router.GET("/api-2", func(c *gin.Context) {
// 		c.JSON(200, gin.H{"success": "Access granted for api-2"})
// 	})

// 	router.Run(":" + port)
// }

////////////////////////////////////////////////////////////////////////////////////////////

// package main

// import (
// 	"os"

// 	"golang-restaurant-management/database"

// 	middleware "golang-restaurant-management/middleware"
// 	routes "golang-restaurant-management/routes"

// 	"github.com/gin-gonic/gin"

// 	"go.mongodb.org/mongo-driver/mongo"
// )

// var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

// func main() {
// 	port := os.Getenv("PORT")

// 	if port == "" {
// 		port = "8000"
// 	}

// 	router := gin.New()
// 	router.Use(gin.Logger())
// 	routes.UserRoutes(router)
// 	router.Use(middleware.Authentication())

// 	routes.FoodRoutes(router)
// 	routes.MenuRoutes(router)
// 	routes.TableRoutes(router)
// 	routes.OrderRoutes(router)
// 	routes.OrderItemRoutes(router)
// 	routes.InvoiceRoutes(router)

// 	router.Run(":" + port)
// }
