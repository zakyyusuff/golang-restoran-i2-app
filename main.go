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
// 	os.Setenv("restaurant", "mongodb+srv://zakymuhammadyusuf:zaky123@zaky.oy6yt60.mongodb.net/")
// 	// url := os.Getenv("restaurant")
// 	// clientOptions := options.Client().ApplyURI(url)
// 	// client, err := mongo.Connect(context.TODO(), clientOptions)
// 	// port := os.Getenv(clientOptions)

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

// //////////////////////////////////////////////////////////////////////////////////////////	1 sudah bisa
// package main

// import (
// 	"fmt"
// 	"os"

// 	middleware "golang-restaurant-management/middleware"
// 	routes "golang-restaurant-management/routes"

// 	"github.com/gin-gonic/gin"
// )

// // var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

// func main() {
// 	fmt.Println("Server is running...")
// 	os.Setenv("restaurant", "mongodb+srv://zakymuhammadyusuf:zaky123@zaky.oy6yt60.mongodb.net/")
// 	// url := os.Getenv("restaurant")
// 	// clientOptions := options.Client().ApplyURI(url)
// 	// client, err := mongo.Connect(context.TODO(), clientOptions)
// 	// port := os.Getenv(clientOptions)

// 	port := os.Getenv("PORT")
// 	// ip := os.Getenv("INTERNALHOST")

// 	// if port == "" {
// 	// 	port = "3000"
// 	// }

// 	router := gin.Default()
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
// 	// router.Run(ip + ":" + port)
// }

////////////////////////////////////////////////////////////////////////////////////////////		2 uji coba

// package main

// import (
// 	"fmt"
// 	"os"

// 	middleware "golang-restaurant-management/middleware"
// 	routes "golang-restaurant-management/routes"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	fmt.Println("Server is running...")
// 	os.Setenv("restaurant", "mongodb+srv://zakymuhammadyusuf:zaky123@zaky.oy6yt60.mongodb.net/")

// 	port := os.Getenv("PORT")
// 	// ip := os.Getenv("INTERNALHOST")

// 	if port == "" {
// 		port = "3000"
// 	}
// 	router := gin.Default()
// 	router.Use(gin.Logger())

// 	// Set the HTML templates directory
// 	// router.SetHTMLTemplate(template.Must(template.ParseGlob("templates/*")))
// 	// router.SetHTMLTemplate(template.Must(template.ParseGlob("templates/*.html")))
// 	router.LoadHTMLGlob("templates/*.html")
// 	// router.LoadHTMLGlob("login.html")

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

// //////////////////////////////////////////////////////////////////////////////////////////	uji coba 3
// package main

// import (
// 	"fmt"
// 	"os"

// 	middleware "golang-restaurant-management/middleware"
// 	routes "golang-restaurant-management/routes"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	fmt.Println("Server is running...")
// 	os.Setenv("restaurant", "mongodb+srv://zakymuhammadyusuf:zaky123@zaky.oy6yt60.mongodb.net/")
// 	port := os.Getenv("PORT")

// 	router := gin.Default()

// 	// Menambahkan pengaturan CORS
// 	router.Use(func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
// 		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(200)
// 		} else {
// 			c.Next()
// 		}
// 	})

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

//////////////////////////////////////////////////////////	script perbaikan local

package main

import (
	"fmt"
	"os"

	middleware "golang-restaurant-management/middleware"
	routes "golang-restaurant-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Server is running...")
	os.Setenv("restaurant", "mongodb+srv://zakymuhammadyusuf:zaky123@zaky.oy6yt60.mongodb.net/")
	port := "8000"

	router := gin.Default()

	// Menambahkan pengaturan CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

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
