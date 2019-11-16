package product

import "context"

type Service interface {
	Fetch(ctx context.Context) ([]*Product, error)
}
