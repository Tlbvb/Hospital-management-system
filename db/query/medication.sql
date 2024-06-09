
-- name: GetMedication :one
SELECT * FROM "Medication"
WHERE id = $1 LIMIT 1;

-- name: GetMedicationByName :one
SELECT * FROM "Medication"
WHERE medicationame = $1 LIMIT 1;

-- name: CreateMedication :one
INSERT INTO "Medication" (
medicationame, description,sideeffects
) VALUES (
  $1, $2,$3
)
RETURNING *;