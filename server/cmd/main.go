package main

import (
	"context"
	"os"

	"github.com/alexadastra/hisho/server/service"
	"github.com/alexadastra/hisho/server/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.Use(cors.Default())

	db := storage.DBinstance()

	if err := db.Ping(context.Background(), nil); err != nil {
		panic(err)
	}

	coll := storage.OpenCollection(db, "orders")
	s := service.New(coll)

	// these are the endpoints
	//C
	router.POST("/order/create", s.AddOrder)
	//R
	router.GET("/waiter/:waiter", s.GetOrdersByWaiter)
	router.GET("/orders", s.GetOrders)
	router.GET("/order/:id/", s.GetOrderById)
	//U
	router.PUT("/waiter/update/:id", s.UpdateWaiter)
	router.PUT("/order/update/:id", s.UpdateOrder)
	//D
	router.DELETE("/order/delete/:id", s.DeleteOrder)

	//this runs the server and allows it to listen to requests.
	router.Run(":" + port)
}
