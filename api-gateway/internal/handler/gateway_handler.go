package handler

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HandleProductProxy(c *gin.Context) {
	proxyRequest(c, os.Getenv("INVENTORY_GRPC_URL"))
}

func HandleOrderProxy(c *gin.Context) {
	proxyRequest(c, os.Getenv("ORDER_GRPC_URL"))
}

func proxyRequest(c *gin.Context, target string) {
	req, err := http.NewRequest(c.Request.Method, "http://"+target+c.Request.RequestURI, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "proxy error"})
		return
	}

	req.Header = c.Request.Header

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "service unreachable"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
