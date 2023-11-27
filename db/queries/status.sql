-- name: InsertStatus :one
INSERT INTO status (
    id_status,
    nama_status
) VALUES (
    $1,$2
) RETURNING *;


-- name: UpdateStatus :exec
UPDATE status
SET nama_status = $1
WHERE id_status = $2;