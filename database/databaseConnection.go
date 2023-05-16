package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
	// MongoDb := "mongodb://localhost:27017"
	// MongoDb := "mongodb+srv://zakymuhammadyusuf:zaky123@zaky.oy6yt60.mongodb.net/?retryWrites=true&w=majority"
	// MongoDb := "mongodb+srv://zakymuhammadyusuf:zaky123@zaky.oy6yt60.mongodb.net/test"
	// MongoDb := "mongodb+srv://zakymuhammadyusuf:zaky123@zaky.oy6yt60.mongodb.net/restaurant"
	os.Setenv("restaurant", "mongodb+srv://zakymuhammadyusuf:zaky123@zaky.oy6yt60.mongodb.net/")
	MongoDb := os.Getenv("restaurant")
	fmt.Print(MongoDb)

	client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)

	return collection
}

////////////////////////////////////////////////////////////////////////////// mysql

// package database

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"time"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func DBinstance() *sql.DB {
// 	DBURL := ":3306/restoran"

// 	db, err := sql.Open("mysql", DBURL)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Atur timeout koneksi
// 	db.SetConnMaxLifetime(5 * time.Minute)
// 	db.SetMaxOpenConns(10)
// 	db.SetMaxIdleConns(5)

// 	// Tes koneksi
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to MySQL")
// 	return db
// }

// var DB *sql.DB = DBinstance()
