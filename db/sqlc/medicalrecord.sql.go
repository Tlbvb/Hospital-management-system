// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: medicalrecord.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createMedRecord = `-- name: CreateMedRecord :one
INSERT INTO "MedicalRecord" (
appointment_id, diagnosis, treatment,notes,patient_id,doctor_id
) VALUES (
  $1, $2,$3,$4,$5,$6
)
RETURNING id, appointment_id, patient_id, doctor_id, diagnosis, treatment, notes, created_at
`

type CreateMedRecordParams struct {
	AppointmentID pgtype.Int8 `json:"appointment_id"`
	Diagnosis     string      `json:"diagnosis"`
	Treatment     string      `json:"treatment"`
	Notes         string      `json:"notes"`
	PatientID     pgtype.Int8 `json:"patient_id"`
	DoctorID      pgtype.Int8 `json:"doctor_id"`
}

func (q *Queries) CreateMedRecord(ctx context.Context, arg CreateMedRecordParams) (MedicalRecord, error) {
	row := q.db.QueryRow(ctx, createMedRecord,
		arg.AppointmentID,
		arg.Diagnosis,
		arg.Treatment,
		arg.Notes,
		arg.PatientID,
		arg.DoctorID,
	)
	var i MedicalRecord
	err := row.Scan(
		&i.ID,
		&i.AppointmentID,
		&i.PatientID,
		&i.DoctorID,
		&i.Diagnosis,
		&i.Treatment,
		&i.Notes,
		&i.CreatedAt,
	)
	return i, err
}

const getMedicalRecord = `-- name: GetMedicalRecord :one
SELECT id, appointment_id, patient_id, doctor_id, diagnosis, treatment, notes, created_at FROM "MedicalRecord"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMedicalRecord(ctx context.Context, id int64) (MedicalRecord, error) {
	row := q.db.QueryRow(ctx, getMedicalRecord, id)
	var i MedicalRecord
	err := row.Scan(
		&i.ID,
		&i.AppointmentID,
		&i.PatientID,
		&i.DoctorID,
		&i.Diagnosis,
		&i.Treatment,
		&i.Notes,
		&i.CreatedAt,
	)
	return i, err
}

const getMedicalRecordAppointment = `-- name: GetMedicalRecordAppointment :one
SELECT id, appointment_id, patient_id, doctor_id, diagnosis, treatment, notes, created_at FROM "MedicalRecord"
WHERE appointment_id = $1 LIMIT 1
`

func (q *Queries) GetMedicalRecordAppointment(ctx context.Context, appointmentID pgtype.Int8) (MedicalRecord, error) {
	row := q.db.QueryRow(ctx, getMedicalRecordAppointment, appointmentID)
	var i MedicalRecord
	err := row.Scan(
		&i.ID,
		&i.AppointmentID,
		&i.PatientID,
		&i.DoctorID,
		&i.Diagnosis,
		&i.Treatment,
		&i.Notes,
		&i.CreatedAt,
	)
	return i, err
}

const getMedicalRecordsDoctor = `-- name: GetMedicalRecordsDoctor :many
SELECT id, appointment_id, patient_id, doctor_id, diagnosis, treatment, notes, created_at FROM "MedicalRecord"
WHERE doctor_id = $1
`

func (q *Queries) GetMedicalRecordsDoctor(ctx context.Context, doctorID pgtype.Int8) ([]MedicalRecord, error) {
	rows, err := q.db.Query(ctx, getMedicalRecordsDoctor, doctorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []MedicalRecord{}
	for rows.Next() {
		var i MedicalRecord
		if err := rows.Scan(
			&i.ID,
			&i.AppointmentID,
			&i.PatientID,
			&i.DoctorID,
			&i.Diagnosis,
			&i.Treatment,
			&i.Notes,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMedicalRecordsPatient = `-- name: GetMedicalRecordsPatient :many
SELECT id, appointment_id, patient_id, doctor_id, diagnosis, treatment, notes, created_at FROM "MedicalRecord"
WHERE patient_id = $1
`

func (q *Queries) GetMedicalRecordsPatient(ctx context.Context, patientID pgtype.Int8) ([]MedicalRecord, error) {
	rows, err := q.db.Query(ctx, getMedicalRecordsPatient, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []MedicalRecord{}
	for rows.Next() {
		var i MedicalRecord
		if err := rows.Scan(
			&i.ID,
			&i.AppointmentID,
			&i.PatientID,
			&i.DoctorID,
			&i.Diagnosis,
			&i.Treatment,
			&i.Notes,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMedicalRecord = `-- name: UpdateMedicalRecord :exec
UPDATE "MedicalRecord"
  set diagnosis = $2,treatment=$3, notes=$4
WHERE id = $1
`

type UpdateMedicalRecordParams struct {
	ID        int64  `json:"id"`
	Diagnosis string `json:"diagnosis"`
	Treatment string `json:"treatment"`
	Notes     string `json:"notes"`
}

func (q *Queries) UpdateMedicalRecord(ctx context.Context, arg UpdateMedicalRecordParams) error {
	_, err := q.db.Exec(ctx, updateMedicalRecord,
		arg.ID,
		arg.Diagnosis,
		arg.Treatment,
		arg.Notes,
	)
	return err
}
