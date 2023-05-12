package main

import (

    // "database/sql"
    // "encoding/json"
    // "log"
    // "net/http"
	"github.com/luthfiyyah-a/go-api-backend/controllers/productcontroller"
	"github.com/luthfiyyah-a/go-api-backend/models"

	"github.com/gin-gonic/gin"
    //"github.com/gorilla/mux"
    //"github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)

	r.Run()
}
