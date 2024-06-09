-- name: GetMedicationByRecord :many
SELECT * FROM "MedicalRecordMedication"
WHERE medical_record_id = $1;

-- name: CreateMedRecMedic :one
INSERT INTO "MedicalRecordMedication" (
medical_record_id,medication_id
) VALUES (
  $1, $2
)
RETURNING *;