package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SpikeController(c *gin.Context) {
	c.String(http.StatusOK, "spike controller ok")
}

func main() {
	g := gin.Default()

	g.GET("spike", SpikeController)

	err := g.Run()

	if err != nil {
		fmt.Println("gin run err", err)
	}
}
