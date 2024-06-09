-- name: GetDepartmentById :one
SELECT * FROM "Department"
WHERE id = $1 LIMIT 1;

-- name: GetDepartmentByName :one
SELECT * FROM "Department"
WHERE name = $1 LIMIT 1;

-- name: GetDepartmentHeadId :one
SELECT head_id FROM "Department"
WHERE name = $1 LIMIT 1;

-- name: GetDepartmentByHead :one
SELECT * FROM "Department"
WHERE head_id = $1 LIMIT 1;

-- name: CreateDepartment :one
INSERT INTO "Department" (
name,description,head_id
) VALUES (
  $1,$2,$3
)
RETURNING *;