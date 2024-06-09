// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: patient.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPatient = `-- name: CreatePatient :one
INSERT INTO "Patient" (
fullname,address
) VALUES (
  $1, $2
)
RETURNING id, fullname, address, created_at
`

type CreatePatientParams struct {
	Fullname pgtype.Text `json:"fullname"`
	Address  pgtype.Text `json:"address"`
}

func (q *Queries) CreatePatient(ctx context.Context, arg CreatePatientParams) (Patient, error) {
	row := q.db.QueryRow(ctx, createPatient, arg.Fullname, arg.Address)
	var i Patient
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}

const getPatientById = `-- name: GetPatientById :one
SELECT id, fullname, address, created_at FROM "Patient"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPatientById(ctx context.Context, id int64) (Patient, error) {
	row := q.db.QueryRow(ctx, getPatientById, id)
	var i Patient
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}

const getPatientByName = `-- name: GetPatientByName :one
SELECT id, fullname, address, created_at FROM "Patient"
WHERE fullname = $1 LIMIT 1
`

func (q *Queries) GetPatientByName(ctx context.Context, fullname pgtype.Text) (Patient, error) {
	row := q.db.QueryRow(ctx, getPatientByName, fullname)
	var i Patient
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.Address,
		&i.CreatedAt,
	)
	return i, err
}
