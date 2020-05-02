package product

import "database/sql"

type Respository interface {
	GetProductById(id int) (*Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(databaseConnection *sql.DB) Respository {
	return &repository{db: databaseConnection}
}

func (repo *repository) GetProductById(id int) (*Product, error) {
	const sql = `SELECT id, product_code, product_name,
				 coalesce(description, ''), standard_cost, list_price, 
				 category FROM products where id=?`

	row := repo.db.QueryRow(sql, id)
	product := &Product{}

	err := row.Scan(&product.Id, &product.ProductCode, &product.ProductName,
		&product.Description, &product.StandardCost, &product.ListPrice, &product.Category)

	if err != nil {
		panic(err)
	}
	return product, err
}
