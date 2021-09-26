package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/healthz", healthz, gin.Logger(), gin.Recovery())
	if err := r.Run(":80"); err != nil {
		panic(err)
	}
}

func healthz(c *gin.Context) {
	headers := make(map[string]string)

	err := c.BindHeader(&headers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	for k, v := range c.Request.Header {
		c.Writer.Header().Set(k, v[0])
	}

	c.Writer.Header().Set("version", os.Getenv("VERSION"))

	c.JSON(http.StatusOK, nil)
}
