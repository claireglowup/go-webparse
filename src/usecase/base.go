package usecase

import (
	"context"
	"testfastprint/src/repository"
)

type Usecase interface {
	InsertProductTx(ctx context.Context, arg []repository.InsertProductParamsTx) error
	GetAllProduct(ctx context.Context) ([]repository.GetAllProductRow, error)
	GetProduct(ctx context.Context, idProduct string) (repository.GetProductRow, error)
	GetProdukForUpdate(ctx context.Context, idProduct string) (repository.GetProductForUpdateRow, error)
	UpdateProductTx(ctx context.Context, arg repository.UpdateProductParamsTx) error
	DeleteProduk(ctx context.Context, idProduk string) error
}

type usecase struct {
	store repository.Store
}

func NewUsecase(store repository.Store) Usecase {
	return &usecase{
		store: store,
	}
}
