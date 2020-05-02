package main

import (
	"net/http"

	"github.com/GolangNorhtwindRestApi/database"
	"github.com/GolangNorhtwindRestApi/product"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	databaseConnection := database.InitDB()
	defer databaseConnection.Close()

	var productRespository = product.NewRepository(databaseConnection)
	var productService product.Service
	productService = product.NewService(productRespository)

	r := chi.NewRouter()
	r.Mount("/productos", product.MakeHttpHandler(productService))
	http.ListenAndServe(":3000", r)
}
