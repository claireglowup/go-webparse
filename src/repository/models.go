// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package repository

import ()

type Kategori struct {
	IDKategori   string `json:"id_kategori"`
	NamaKategori string `json:"nama_kategori"`
}

type Produk struct {
	IDProduk   string `json:"id_produk"`
	NamaProduk string `json:"nama_produk"`
	Harga      string `json:"harga"`
	KategoriID string `json:"kategori_id"`
	StatusID   string `json:"status_id"`
}

type Status struct {
	IDStatus   string `json:"id_status"`
	NamaStatus string `json:"nama_status"`
}
