package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var products []request
var lastID int

type request struct {
	ID       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Tipo     string  `json:"tipo"`
	Cantidad int     `json:"cantidad"`
	Precio   float64 `json:"precio"`
}

func Guardar() gin.HandlerFunc {
	return func(ctxt *gin.Context) {
		var req request
		if err := ctxt.ShouldBindJSON(&req); err != nil {
			ctxt.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		lastID++
		req.ID = lastID
		products = append(products, req)
		ctxt.JSON(http.StatusOK, req)
	}
}

func main() {
	router := gin.Default()
	pr := router.Group("/productos")
	pr.POST("/", Guardar())

	router.Run()
}
