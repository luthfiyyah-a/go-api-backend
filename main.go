package main

import (
    // "database/sql"
    // "encoding/json"
    // "log"
    // "net/http"
	"restapi-minio.com/controllers/productcontroller"
	"restapi-minio.com/models"

	"github.com/gin-gonic/gin"
    //"github.com/gorilla/mux"
    //"github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.GET("/api/product/:name", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	r.Run()
}
