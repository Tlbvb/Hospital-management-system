-- name: GetDoctorById :one
SELECT * FROM "Doctor"
WHERE id = $1 LIMIT 1;

-- name: GetDoctorByName :one
SELECT * FROM "Doctor"
WHERE fullname = $1 LIMIT 1;


-- name: GetDoctorsOfDepartment :many
SELECT * FROM "Doctor"
WHERE department_id = $1;

-- name: CreateDoctor :one
INSERT INTO "Doctor" (
fullname,specialty,department_id
) VALUES (
  $1, $2,$3
)
RETURNING *;