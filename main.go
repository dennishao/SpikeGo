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
	Sku   int
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

func getProducts(c *gin.Context) {
	var products []Product
	DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func addProduct(c *gin.Context) {
	product := Product{Name: "lemon", Price: 8, Sku: 20000}
	res := DB.Create(&product)

	if res.RowsAffected == 1 {
		c.JSON(http.StatusOK, gin.H{"code": 2000, "message": "新增成功 "})
	}
}

func buyProduct(c *gin.Context) {
	name := c.Param("name")
	var product Product
	res := DB.Find(&product, "name = ?", name)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 2000, "message": "不存在此商品"})
		return
	}

	product.Sku = product.Sku - 1
	res1 := DB.Save(&product)
	if res1.RowsAffected == 1 {
		c.JSON(http.StatusOK, gin.H{"code": 2000, "message": "抢购成功"})
		return
	}
}

func main() {
	g := gin.Default()

	g.GET("/products", getProducts)

	g.POST("/products", addProduct)

	g.GET("/kill/:name", buyProduct)

	err := g.Run()

	if err != nil {
		fmt.Println("gin run err", err)
	}

}
