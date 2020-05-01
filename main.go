package main

import (
	"GolangNorhtwindRestApi/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

var databaseConnection *sql.DB

type Producto struct {
	ID            int    `json:"id"`
	Producto_code string `json:"product_code"`
	Description   string `json:"description"`
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {

	databaseConnection = database.InitDB()
	defer databaseConnection.Close()

	r := chi.NewRouter()
	r.Get("/products", AllProductos)
	r.Post("/create/productos", CreateProducto)
	r.Put("/update/productos/{id}", UpdateProducto)
	r.Delete("/delete/productos/{id}", DeleteProducto)
	http.ListenAndServe(":3000", r)
}

func CreateProducto(w http.ResponseWriter, r *http.Request) {
	var product Producto
	json.NewDecoder(r.Body).Decode(&product)

	query, err := databaseConnection.Prepare("Insert products SET product_code=?, description=?")
	catch(err)

	_, er := query.Exec(&product.Producto_code, &product.Description)
	catch(er)

	defer query.Close()
	responsewithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

func UpdateProducto(w http.ResponseWriter, r *http.Request) {
	var product Producto
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&product)

	query, err := databaseConnection.Prepare("Update products set product_code=?, description=? where id=?")
	catch(err)

	_, er := query.Exec(&product.Producto_code, &product.Description, id)
	catch(er)

	defer query.Close()
	responsewithJSON(w, http.StatusCreated, map[string]string{"message": "delete successfully"})
}

func DeleteProducto(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query, err := databaseConnection.Prepare("delete from products where id=?")
	catch(err)

	_, er := query.Exec(id)
	catch(er)

	defer query.Close()
	responsewithJSON(w, http.StatusCreated, map[string]string{"message": "update successfully"})
}

func AllProductos(w http.ResponseWriter, r *http.Request) {
	const sql = `SELECT id,product_code, COALESCE(description, '') FROM products;`
	results, err := databaseConnection.Query(sql)
	catch(err)
	var products []*Producto

	for results.Next() {
		product := &Producto{}
		err = results.Scan(&product.ID, &product.Producto_code, &product.Description)

		catch(err)
		products = append(products, product)
	}
	responsewithJSON(w, http.StatusOK, products)
}
func responsewithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
