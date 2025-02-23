// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: kategori.sql

package repository

import (
	"context"
)

const insertKategori = `-- name: InsertKategori :one
INSERT INTO kategori (
    id_kategori,
    nama_kategori
) VALUES (
    $1,$2
) RETURNING id_kategori, nama_kategori
`

type InsertKategoriParams struct {
	IDKategori   string `json:"id_kategori"`
	NamaKategori string `json:"nama_kategori"`
}

func (q *Queries) InsertKategori(ctx context.Context, arg InsertKategoriParams) (Kategori, error) {
	row := q.db.QueryRowContext(ctx, insertKategori, arg.IDKategori, arg.NamaKategori)
	var i Kategori
	err := row.Scan(&i.IDKategori, &i.NamaKategori)
	return i, err
}

const updateKategori = `-- name: UpdateKategori :exec
UPDATE kategori
SET nama_kategori = $1
WHERE id_kategori = $2
`

type UpdateKategoriParams struct {
	NamaKategori string `json:"nama_kategori"`
	IDKategori   string `json:"id_kategori"`
}

func (q *Queries) UpdateKategori(ctx context.Context, arg UpdateKategoriParams) error {
	_, err := q.db.ExecContext(ctx, updateKategori, arg.NamaKategori, arg.IDKategori)
	return err
}
