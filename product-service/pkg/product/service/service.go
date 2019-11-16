package service

import (
	"context"
	"sync"

	"github.com/klferreira/hash-challenge/product-service/pkg/product"
	"github.com/klferreira/hash-challenge/product-service/pkg/util/authctx"
	pb "github.com/klferreira/hash-challenge/product-service/proto"
	"github.com/pkg/errors"
)

type productService struct {
	repo     product.Repository
	discount pb.DiscountServiceClient
}

func NewProductService(repo product.Repository, discount pb.DiscountServiceClient) product.Service {
	return &productService{repo, discount}
}

func (s *productService) getDiscount(ctx context.Context, p *product.Product, userID int64, done chan error) {
	req := &pb.GetDiscountRequest{ProductID: p.ID, UserID: userID}
	res, err := s.discount.GetDiscount(ctx, req)
	if err == nil {
		p.Discount = &product.Discount{
			Percent:      res.Discount.GetPercentual(),
			ValueInCents: res.Discount.GetPriceInCents(),
		}
	} else {
		done <- err
	}

	done <- nil
}

func (s *productService) fillDiscounts(ctx context.Context, products []*product.Product, userID int64) ([]*product.Product, error) {

	var wg = sync.WaitGroup{}

	for _, p := range products {
		wg.Add(1)
		var done = make(chan error)
		go s.getDiscount(ctx, p, userID, done)
		<-done
		wg.Done()
	}

	wg.Wait()

	return products, nil
}

func (s *productService) Fetch(ctx context.Context) ([]*product.Product, error) {
	ps, err := s.repo.Fetch(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Could not fetch products")
	}

	if userID, ok := authctx.FromContext(ctx); !ok {
		s.fillDiscounts(ctx, ps, userID)
	}

	return ps, nil
}
