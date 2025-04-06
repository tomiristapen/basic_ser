package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	ginprometheus "github.com/zsais/go-gin-prometheus"
	"api-gateway/internal/middleware"
	"api-gateway/internal/routes"
)

func main() {
	// Загружаем .env файл
	if err := godotenv.Load(); err != nil {
		log.Println(" .env файл не найден, продолжаем...")
	}

	// Получаем порт из .env или по умолчанию
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	// Инициализация Gin
	router := gin.Default()

	// Кастомный логгер
	router.Use(middleware.RequestLogger())

	// Prometheus middleware
	prometheus := ginprometheus.NewPrometheus("api_gateway")
	prometheus.Use(router)

	// роуты
	routes.SetupRoutes(router)

	log.Println(" API Gateway работает на порту:", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(" Ошибка запуска сервера:", err)
	}
}
