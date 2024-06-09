CREATE TABLE "Patient" (
  "id" bigserial PRIMARY KEY,
  "fullname" varchar,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

-- name: GetPatientById :one
SELECT * FROM "Patient"
WHERE id = $1 LIMIT 1;

-- name: GetPatientByName :one
SELECT * FROM "Patient"
WHERE fullname = $1 LIMIT 1;

-- name: CreatePatient :one
INSERT INTO "Patient" (
fullname,address
) VALUES (
  $1, $2
)
RETURNING *;