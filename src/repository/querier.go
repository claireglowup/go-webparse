// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package repository

import (
	"context"
)

type Querier interface {
	DeleteProduk(ctx context.Context, idProduk string) error
	GetAllProduct(ctx context.Context) ([]GetAllProductRow, error)
	GetProduct(ctx context.Context, idProduk string) (GetProductRow, error)
	GetProductForUpdate(ctx context.Context, idProduk string) (GetProductForUpdateRow, error)
	InsertKategori(ctx context.Context, arg InsertKategoriParams) (Kategori, error)
	InsertProduct(ctx context.Context, arg InsertProductParams) (Produk, error)
	InsertStatus(ctx context.Context, arg InsertStatusParams) (Status, error)
	UpdateKategori(ctx context.Context, arg UpdateKategoriParams) error
	UpdateProduct(ctx context.Context, arg UpdateProductParams) error
	UpdateStatus(ctx context.Context, arg UpdateStatusParams) error
}

var _ Querier = (*Queries)(nil)