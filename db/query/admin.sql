-- name: GetAdminById :one
SELECT * FROM "Admin"
WHERE id = $1 LIMIT 1;

-- name: GetAdminByName :one
SELECT * FROM "Admin"
WHERE fullname = $1 LIMIT 1;

-- name: CreateAdmin :one
INSERT INTO "Admin" (
fullname
) VALUES (
  $1
)
RETURNING *;