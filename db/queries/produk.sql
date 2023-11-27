-- name: InsertProduct :one
INSERT INTO produk (
    id_produk,
    nama_produk,
    harga,
    kategori_id,
    status_id
) VALUES (
    $1,$2,$3,$4,$5
) RETURNING *;



-- name: GetAllProduct :many
SELECT produk.id_produk,produk.nama_produk,produk.harga,kategori.nama_kategori,status.nama_status
FROM produk
JOIN status ON produk.status_id = status.id_status
JOIN kategori ON produk.kategori_id = kategori.id_kategori
WHERE status.nama_status = 'bisa dijual';



-- name: GetProduct :one
SELECT produk.id_produk,produk.nama_produk,produk.harga,kategori.nama_kategori,status.nama_status
FROM produk
JOIN status ON produk.status_id = status.id_status
JOIN kategori ON produk.kategori_id = kategori.id_kategori
WHERE produk.id_produk = $1;


-- name: UpdateProduct :exec
UPDATE produk
SET nama_produk = $1, 
    harga = $2
WHERE id_produk = $3;

-- name: GetProductForUpdate :one
SELECT status.id_status,kategori.id_kategori
FROM produk
JOIN status ON produk.status_id = status.id_status
JOIN kategori ON produk.kategori_id = kategori.id_kategori
WHERE produk.id_produk = $1;


-- name: DeleteProduk :exec
DELETE FROM produk
WHERE id_produk = $1;