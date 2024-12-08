package repository

import (
	"context"
	models "controlEntity/pkg/model"
	"database/sql"
	"errors"
)

type ProductRepositoryDB struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepositoryDB {
	return &ProductRepositoryDB{db: db}
}

func (r *ProductRepositoryDB) CreateProduct(ctx context.Context, product models.Product) (models.Product, error) {
	query := "INSERT INTO products (name, description, price, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id"
	var id int
	err := r.db.QueryRow(query, product.Name, product.Description, product.Price, product.UserID).Scan(&id)
	return models.Product{}, err
}

func (r *ProductRepositoryDB) GetProductID(ctx context.Context, id int) (models.Product, error) {
	query := "SELECT id, name, description, price, user_id, created_at, updated_at FROM products WHERE id = $1"
	var product models.Product
	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.UserID, &product.CreatedAt, &product.UpdatedAt)
	if err == sql.ErrNoRows {
		return models.Product{}, nil
	}
	return product, err
}

func (r *ProductRepositoryDB) UpdateProduct(ctx context.Context, id int, product models.Product) error {
	query := "UPDATE products SET name = $1, description = $2, price = $3, updated_at = NOW() WHERE id = $4"
	_, err := r.db.Exec(query, product.Name, product.Description, product.Price, id)
	return err
}

func (r *ProductRepositoryDB) DeleteProduct(ctx context.Context, id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return errors.New("product not found")
	}
	return nil
}

func (r *ProductRepositoryDB) GetListProduct(ctx context.Context) ([]models.Product, error) {
	query := "SELECT id, name, description, price, user_id, created_at, updated_at FROM products"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.UserID, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepositoryDB) GetProductUserID(ctx context.Context, userID int) ([]models.Product, error) {
	query := "SELECT id, name, description, price, user_id, created_at, updated_at FROM products WHERE user_id = $1"
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.UserID, &product.CreatedAt, &product.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
