package repository

import (
	"context"

	"github.com/google/uuid"
)

type InsertProductParamsTx struct {
	IdProduk   string `json:"id_produk"`
	NamaProduk string `json:"nama_produk"`
	Kategori   string `json:"kategori"`
	Harga      string `json:"harga"`
	Status     string `json:"status"`
}

type UpdateProductParamsTx struct {
	IdProduk   string `json:"id_produk"`
	NamaProduk string `json:"nama_produk"`
	Kategori   string `json:"kategori"`
	Harga      string `json:"harga"`
	Status     string `json:"status"`
	IdStatus   string `json:"id_status"`
	IdKategori string `json:"id_kategori"`
}

func (s *sqlStore) UpdateProductTx(ctx context.Context, arg UpdateProductParamsTx) error {

	var err error

	err = s.ExecTx(ctx, func(q *Queries) error {

		err = q.UpdateProduct(ctx, UpdateProductParams{
			NamaProduk: arg.NamaProduk,
			Harga:      arg.Harga,
			IDProduk:   arg.IdProduk,
		})
		if err != nil {
			return err
		}

		err = q.UpdateKategori(ctx, UpdateKategoriParams{
			IDKategori:   arg.IdKategori,
			NamaKategori: arg.Kategori,
		})
		if err != nil {
			return err
		}

		err = q.UpdateStatus(ctx, UpdateStatusParams{
			IDStatus:   arg.IdStatus,
			NamaStatus: arg.Status,
		})
		if err != nil {
			return err
		}

		return err

	})

	if err != nil {
		return err
	}

	return nil

}

func (s *sqlStore) InsertProductManyTx(ctx context.Context, arg []InsertProductParamsTx) error {

	err := s.ExecTx(ctx, func(q *Queries) error {

		var (
			err                      error
			uuidKategori, uuidStatus string
			argInsertProduct         InsertProductParams
			argInsertKategori        InsertKategoriParams
			argInsertStatus          InsertStatusParams
		)

		for _, each := range arg {

			uuidKategori = uuid.New().String()
			uuidStatus = uuid.New().String()

			argInsertProduct = InsertProductParams{
				IDProduk:   each.IdProduk,
				NamaProduk: each.NamaProduk,
				Harga:      each.Harga,
				KategoriID: uuidKategori,
				StatusID:   uuidStatus,
			}

			argInsertKategori = InsertKategoriParams{
				IDKategori:   uuidKategori,
				NamaKategori: each.Kategori,
			}

			argInsertStatus = InsertStatusParams{
				IDStatus:   uuidStatus,
				NamaStatus: each.Status,
			}

			_, err = q.InsertKategori(ctx, argInsertKategori)
			if err != nil {
				return err
			}

			_, err = q.InsertStatus(ctx, argInsertStatus)
			if err != nil {
				return err
			}

			_, err = q.InsertProduct(ctx, argInsertProduct)
			if err != nil {
				return err
			}

		}

		return err

	})

	if err != nil {
		return err
	}

	return err

}

func (s *sqlStore) InsertOneProductTx(ctx context.Context, arg InsertProductParamsTx) error {

	var err error

	uuidKategori := uuid.New().String()
	uuidStatus := uuid.New().String()

	err = s.ExecTx(ctx, func(q *Queries) error {

		argInsertProduct := InsertProductParams{
			IDProduk:   arg.IdProduk,
			NamaProduk: arg.NamaProduk,
			Harga:      arg.Harga,
			KategoriID: uuidKategori,
			StatusID:   uuidStatus,
		}

		argInsertKategori := InsertKategoriParams{
			IDKategori:   uuidKategori,
			NamaKategori: arg.Kategori,
		}

		argInsertStatus := InsertStatusParams{
			IDStatus:   uuidStatus,
			NamaStatus: arg.Status,
		}

		_, err = q.InsertKategori(ctx, argInsertKategori)
		if err != nil {
			return err
		}

		_, err = q.InsertStatus(ctx, argInsertStatus)
		if err != nil {
			return err
		}

		_, err = q.InsertProduct(ctx, argInsertProduct)
		if err != nil {
			return err
		}

		return err

	})

	if err != nil {
		return err
	}

	return err

}
