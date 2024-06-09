// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: appointment.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAppointment = `-- name: CreateAppointment :one
INSERT INTO "Appointment" (
time, patient_id,doctor_id, visitreason,status
) VALUES (
  $1, $2,$3,$4,&5
)
RETURNING id, time, patient_id, doctor_id, visitreason, status, created_at
`

type CreateAppointmentParams struct {
	Time        time.Time   `json:"time"`
	PatientID   pgtype.Int8 `json:"patient_id"`
	DoctorID    pgtype.Int8 `json:"doctor_id"`
	Visitreason string      `json:"visitreason"`
}

func (q *Queries) CreateAppointment(ctx context.Context, arg CreateAppointmentParams) (Appointment, error) {
	row := q.db.QueryRow(ctx, createAppointment,
		arg.Time,
		arg.PatientID,
		arg.DoctorID,
		arg.Visitreason,
	)
	var i Appointment
	err := row.Scan(
		&i.ID,
		&i.Time,
		&i.PatientID,
		&i.DoctorID,
		&i.Visitreason,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getAppointment = `-- name: GetAppointment :one
SELECT id, time, patient_id, doctor_id, visitreason, status, created_at FROM "Appointment"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAppointment(ctx context.Context, id int64) (Appointment, error) {
	row := q.db.QueryRow(ctx, getAppointment, id)
	var i Appointment
	err := row.Scan(
		&i.ID,
		&i.Time,
		&i.PatientID,
		&i.DoctorID,
		&i.Visitreason,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const listAppointmentsDoctor = `-- name: ListAppointmentsDoctor :many
SELECT id, time, patient_id, doctor_id, visitreason, status, created_at FROM "Appointment"
WHERE doctor_id = $1 order by time
`

func (q *Queries) ListAppointmentsDoctor(ctx context.Context, doctorID pgtype.Int8) ([]Appointment, error) {
	rows, err := q.db.Query(ctx, listAppointmentsDoctor, doctorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Appointment{}
	for rows.Next() {
		var i Appointment
		if err := rows.Scan(
			&i.ID,
			&i.Time,
			&i.PatientID,
			&i.DoctorID,
			&i.Visitreason,
			&i.Status,
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

const listAppointmentsDoctorStatus = `-- name: ListAppointmentsDoctorStatus :many
SELECT id, time, patient_id, doctor_id, visitreason, status, created_at FROM "Appointment"
WHERE doctor_id = $1 AND status=$2 order by time
`

type ListAppointmentsDoctorStatusParams struct {
	DoctorID pgtype.Int8 `json:"doctor_id"`
	Status   string      `json:"status"`
}

func (q *Queries) ListAppointmentsDoctorStatus(ctx context.Context, arg ListAppointmentsDoctorStatusParams) ([]Appointment, error) {
	rows, err := q.db.Query(ctx, listAppointmentsDoctorStatus, arg.DoctorID, arg.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Appointment{}
	for rows.Next() {
		var i Appointment
		if err := rows.Scan(
			&i.ID,
			&i.Time,
			&i.PatientID,
			&i.DoctorID,
			&i.Visitreason,
			&i.Status,
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

const listAppointmentsPatient = `-- name: ListAppointmentsPatient :many
SELECT id, time, patient_id, doctor_id, visitreason, status, created_at FROM "Appointment"
WHERE patient_id = $1 ORDER BY time
`

func (q *Queries) ListAppointmentsPatient(ctx context.Context, patientID pgtype.Int8) ([]Appointment, error) {
	rows, err := q.db.Query(ctx, listAppointmentsPatient, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Appointment{}
	for rows.Next() {
		var i Appointment
		if err := rows.Scan(
			&i.ID,
			&i.Time,
			&i.PatientID,
			&i.DoctorID,
			&i.Visitreason,
			&i.Status,
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

const listAppointmentsPatientDoctorStatus = `-- name: ListAppointmentsPatientDoctorStatus :many
SELECT id, time, patient_id, doctor_id, visitreason, status, created_at FROM "Appointment"
WHERE doctor_id = $1 AND status=$2 AND patient_id=$3 order by time
`

type ListAppointmentsPatientDoctorStatusParams struct {
	DoctorID  pgtype.Int8 `json:"doctor_id"`
	Status    string      `json:"status"`
	PatientID pgtype.Int8 `json:"patient_id"`
}

func (q *Queries) ListAppointmentsPatientDoctorStatus(ctx context.Context, arg ListAppointmentsPatientDoctorStatusParams) ([]Appointment, error) {
	rows, err := q.db.Query(ctx, listAppointmentsPatientDoctorStatus, arg.DoctorID, arg.Status, arg.PatientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Appointment{}
	for rows.Next() {
		var i Appointment
		if err := rows.Scan(
			&i.ID,
			&i.Time,
			&i.PatientID,
			&i.DoctorID,
			&i.Visitreason,
			&i.Status,
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

const listAppointmentsPatientStatus = `-- name: ListAppointmentsPatientStatus :many
SELECT id, time, patient_id, doctor_id, visitreason, status, created_at FROM "Appointment"
WHERE patient_id = $1 AND status=$2 order by time
`

type ListAppointmentsPatientStatusParams struct {
	PatientID pgtype.Int8 `json:"patient_id"`
	Status    string      `json:"status"`
}

func (q *Queries) ListAppointmentsPatientStatus(ctx context.Context, arg ListAppointmentsPatientStatusParams) ([]Appointment, error) {
	rows, err := q.db.Query(ctx, listAppointmentsPatientStatus, arg.PatientID, arg.Status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Appointment{}
	for rows.Next() {
		var i Appointment
		if err := rows.Scan(
			&i.ID,
			&i.Time,
			&i.PatientID,
			&i.DoctorID,
			&i.Visitreason,
			&i.Status,
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

const updateAppointmentStatus = `-- name: UpdateAppointmentStatus :exec
UPDATE "Appointment"
  set status = $2
WHERE id = $1
`

type UpdateAppointmentStatusParams struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

func (q *Queries) UpdateAppointmentStatus(ctx context.Context, arg UpdateAppointmentStatusParams) error {
	_, err := q.db.Exec(ctx, updateAppointmentStatus, arg.ID, arg.Status)
	return err
}
