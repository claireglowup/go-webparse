-- name: InsertKategori :one
INSERT INTO kategori (
    id_kategori,
    nama_kategori
) VALUES (
    $1,$2
) RETURNING *;


-- name: UpdateKategori :exec
UPDATE kategori
SET nama_kategori = $1
WHERE id_kategori = $2;