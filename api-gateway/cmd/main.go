package main

import (
	"log"
	"net/http"
	"os"

	"api-gateway/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginprometheus "github.com/zsais/go-gin-prometheus"

	"google.golang.org/grpc"
	orderpb "assignment1/order-service/proto"
	inventorypb "assignment1/inventory-service/proto"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env файл не найден, продолжаем...")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	orderConn, err := grpc.Dial(os.Getenv("ORDER_GRPC_URL"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось подключиться к order-service: %v", err)
	}
	defer orderConn.Close()
	orderClient := orderpb.NewOrderServiceClient(orderConn)

	inventoryConn, err := grpc.Dial(os.Getenv("INVENTORY_GRPC_URL"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось подключиться к inventory-service: %v", err)
	}
	defer inventoryConn.Close()
	inventoryClient := inventorypb.NewInventoryServiceClient(inventoryConn)

	r := gin.Default()
	prometheus := ginprometheus.NewPrometheus("api_gateway")
	prometheus.Use(r)

	h := handler.NewHandler(orderClient, inventoryClient)

	r.POST("/order", h.CreateOrder)
	r.GET("/order/:id", h.GetOrderByID)
	r.PATCH("/order/:id/status", h.UpdateOrderStatus)
	r.GET("/order/user/:userID", h.ListUserOrders)

	r.POST("/inventory/products", h.CreateProduct)
	r.GET("/inventory/products/:id", h.GetProductByID)
	r.PATCH("/inventory/products/:id", h.UpdateProduct)
	r.DELETE("/inventory/products/:id", h.DeleteProduct)
	r.GET("/inventory/products", h.ListProducts)

	log.Println("API Gateway запущен на порту:", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Ошибка запуска API Gateway:", err)
	}
}


// package main

// import (
// 	"log"
// 	"os"

// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"

// 	ginprometheus "github.com/zsais/go-gin-prometheus"
// 	"api-gateway/internal/middleware"
// 	"api-gateway/internal/routes"
// )

// func main() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Println(" .env файл не найден, продолжаем...")
// 	}

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8082"
// 	}
// 	router := gin.Default()

// 	router.Use(middleware.RequestLogger())

// 	prometheus := ginprometheus.NewPrometheus("api_gateway")
// 	prometheus.Use(router)

// 	routes.SetupRoutes(router)

// 	log.Println(" API Gateway работает на порту:", port)
// 	if err := router.Run(":" + port); err != nil {
// 		log.Fatal(" Ошибка запуска сервера:", err)
// 	}
// }
