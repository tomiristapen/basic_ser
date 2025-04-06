package handler

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HandleProductProxy(c *gin.Context) {
	proxyRequest(c, os.Getenv("INVENTORY_SERVICE_URL"))
}

func HandleOrderProxy(c *gin.Context) {
	proxyRequest(c, os.Getenv("ORDER_SERVICE_URL"))
}

func proxyRequest(c *gin.Context, target string) {
	// Создаем новый HTTP-запрос
	req, err := http.NewRequest(c.Request.Method, target+c.Request.RequestURI, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "proxy error"})
		return
	}

	// Копируем заголовки
	req.Header = c.Request.Header

	// Отправляем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "service unreachable"})
		return
	}
	defer resp.Body.Close()

	// Копируем ответ
	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
