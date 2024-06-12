
-- name: GetMedicalRecord :one
SELECT * FROM "MedicalRecord"
WHERE id = $1 LIMIT 1;

-- name: GetMedicalRecordAppointment :one
SELECT * FROM "MedicalRecord"
WHERE appointment_id = $1 LIMIT 1;

-- name: GetMedicalRecordsPatient :many
SELECT * FROM "MedicalRecord"
WHERE patient_id = $1;


-- name: GetMedicalRecordsDoctor :many
SELECT * FROM "MedicalRecord"
WHERE doctor_id = $1;

-- name: CreateMedRecord :one
INSERT INTO "MedicalRecord" (
appointment_id, diagnosis, treatment,notes,patient_id,doctor_id
) VALUES (
  $1, $2,$3,$4,$5,$6
)
RETURNING *;

-- name: UpdateMedicalRecord :exec
UPDATE "MedicalRecord"
  set diagnosis = $2,treatment=$3, notes=$4
WHERE id = $1;


