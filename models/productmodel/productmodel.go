package productmodel

import (
	"go-crud/config"
	"go-crud/entities"
)

func GetAll() []entities.Product {
	rows, err := config.DB.Query(`
	SELECT 
		products.id,
		products.name,
		categories.name as category_name,
		products.stock,
		products.description,
		products.created_at,
		products.updated_at
	FROM products JOIN categories ON products.category_id = categories.id
	`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Category.Name,
			&product.Stock,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt); err != nil {
			panic(err)
		}
		products = append(products, product)
	}
	return products
}

func Create(product entities.Product) bool {
	result, err := config.DB.Exec(`
	INSERT INTO products (name, category_id, stock, description,  created_at, updated_at)
	VALUE (?,?,?,?,?,?)`,
		product.Name, product.Category.Id, product.Stock, product.Description, product.CreatedAt, product.UpdatedAt,
	)
	if err != nil {
		panic(err)
	}
	LastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return LastInsertId > 0
}

func GetDetail(id int) entities.Product {
	row := config.DB.QueryRow(`
	SELECT 
		products.id,
		products.name,
		products.category_id,
		products.stock,
		products.description,
		products.created_at,
		products.updated_at
	FROM products WHERE id = ?
	`, id)

	var product entities.Product
	if err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Category.Id,
		&product.Stock,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt); err != nil {
		panic(err)
	}

	return product
}

func Detail(id int) entities.Product {
	row := config.DB.QueryRow(`
	SELECT 
		products.id,
		products.name,
		categories.name as category_name,
		products.stock,
		products.description,
		products.created_at,
		products.updated_at
	FROM products JOIN categories ON products.category_id = categories.id
	`)

	var product entities.Product
	if err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Category.Name,
		&product.Stock,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt); err != nil {
		panic(err)
	}

	return product
}

func Update(id int, product entities.Product) bool {
	query, err := config.DB.Exec(`UPDATE products SET name=?, category_id=?, stock=?, description=?, updated_at=? WHERE id=?`, product.Name, product.Category.Id, product.Stock, product.Description, product.CreatedAt, id)
	if err != nil {
		panic(err)
	}
	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}
	return result > 0
}
func Delete(id int) error {
	_, err := config.DB.Exec(`DELETE FROM products WHERE id=?`, id)
	return err
}
