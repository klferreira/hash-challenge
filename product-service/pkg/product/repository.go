package product

import "context"

type Repository interface {
	Fetch(ctx context.Context) ([]*Product, error)
}
