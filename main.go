package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Product struct {
	Id    int
	Name  string
	Price int
}

var DB *gorm.DB

func init() {
	//连接数据库
	dsn := "root:root@tcp(127.0.0.1:3306)/ming?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("connect db err", err)
	}

	//插入模型
	err1 := db.AutoMigrate(&Product{})
	if err1 != nil {
		fmt.Println("AutoMigrate err", err1)
	}

	DB = db
}

func SpikeController(c *gin.Context) {
	c.String(http.StatusOK, "spike controller ok")
}

func getProducts(c *gin.Context) {
	var products []Product
	DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func main() {
	g := gin.Default()

	g.GET("spike", SpikeController)

	g.GET("/products", getProducts)

	err := g.Run()

	if err != nil {
		fmt.Println("gin run err", err)
	}

}
