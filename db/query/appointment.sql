-- name: GetAppointment :one
SELECT * FROM "Appointment"
WHERE id = $1 LIMIT 1;

-- name: ListAppointmentsPatient :many
SELECT * FROM "Appointment"
WHERE patient_id = $1 ORDER BY start_time;

-- name: ListAppointmentsDoctor :many
SELECT * FROM "Appointment"
WHERE doctor_id = $1 order by start_time; 

-- name: ListAppointmentsDoctorStatus :many
SELECT * FROM "Appointment"
WHERE doctor_id = $1 AND status=$2 order by start_time; 

-- name: ListAppointmentsPatientStatus :many
SELECT * FROM "Appointment"
WHERE patient_id = $1 AND status=$2 order by start_time; 

-- name: ListAppointmentsPatientDoctorStatus :many
SELECT * FROM "Appointment"
WHERE doctor_id = $1 AND status=$2 AND patient_id=$3 order by start_time; 

-- name: CreateAppointment :one
INSERT INTO "Appointment" (
start_time,end_time, patient_id,doctor_id, visitreason,status
) VALUES (
  $1, $2,$3,$4,$5,$6
)
RETURNING *;

-- name: UpdateAppointmentStatus :exec
UPDATE "Appointment"
  set status = $2
WHERE id = $1;