package repository

import (
	"context"
	"database/sql"

	"github.com/klferreira/hash-challenge/product-service/pkg/product"
)

type mysqlProductRepository struct {
	conn *sql.DB
}

func NewMySQLProductRepository(Conn *sql.DB) product.Repository {
	return &mysqlProductRepository{Conn}
}

func (m *mysqlProductRepository) Fetch(ctx context.Context) ([]*product.Product, error) {
	query := `
		SELECT id, price_in_cents, title, description FROM products
	`

	rows, err := m.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*product.Product
	for rows.Next() {
		p := &product.Product{}
		if err := rows.Scan(&p.ID, &p.PriceInCents, &p.Title, &p.Description); err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}
