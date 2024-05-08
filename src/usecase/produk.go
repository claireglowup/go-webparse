package usecase

import (
	"context"
	"testfastprint/src/repository"
)

func (u *usecase) InsertProductManyTx(ctx context.Context, arg []repository.InsertProductParamsTx) error {

	err := u.store.InsertProductManyTx(ctx, arg)
	if err != nil {
		return err
	}

	return nil

}

func (u *usecase) InsertOneProductTx(ctx context.Context, arg repository.InsertProductParamsTx) error {

	err := u.store.InsertOneProductTx(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}

func (u *usecase) GetAllProduct(ctx context.Context) ([]repository.GetAllProductRow, error) {

	result, err := u.store.GetAllProduct(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *usecase) GetProduct(ctx context.Context, idProduct string) (repository.GetProductRow, error) {

	result, err := u.store.GetProduct(ctx, idProduct)
	if err != nil {
		return result, err
	}

	return result, nil

}

func (u *usecase) UpdateProductTx(ctx context.Context, arg repository.UpdateProductParamsTx) error {
	err := u.store.UpdateProductTx(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}

func (u *usecase) GetProdukForUpdate(ctx context.Context, idProduct string) (repository.GetProductForUpdateRow, error) {

	produk, err := u.store.GetProductForUpdate(ctx, idProduct)
	if err != nil {
		return produk, err
	}

	return produk, nil

}

func (u *usecase) DeleteProduk(ctx context.Context, idProduk string) error {

	err := u.store.DeleteProduk(ctx, idProduk)
	if err != nil {
		return err
	}

	return nil

}
