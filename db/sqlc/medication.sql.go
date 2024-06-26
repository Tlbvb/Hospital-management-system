// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: medication.sql

package db

import (
	"context"
)

const createMedication = `-- name: CreateMedication :one
INSERT INTO "Medication" (
medicationame, description,sideeffects
) VALUES (
  $1, $2,$3
)
RETURNING id, medicationame, description, sideeffects, created_at
`

type CreateMedicationParams struct {
	Medicationame string `json:"medicationame"`
	Description   string `json:"description"`
	Sideeffects   string `json:"sideeffects"`
}

func (q *Queries) CreateMedication(ctx context.Context, arg CreateMedicationParams) (Medication, error) {
	row := q.db.QueryRow(ctx, createMedication, arg.Medicationame, arg.Description, arg.Sideeffects)
	var i Medication
	err := row.Scan(
		&i.ID,
		&i.Medicationame,
		&i.Description,
		&i.Sideeffects,
		&i.CreatedAt,
	)
	return i, err
}

const getMedication = `-- name: GetMedication :one
SELECT id, medicationame, description, sideeffects, created_at FROM "Medication"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetMedication(ctx context.Context, id int64) (Medication, error) {
	row := q.db.QueryRow(ctx, getMedication, id)
	var i Medication
	err := row.Scan(
		&i.ID,
		&i.Medicationame,
		&i.Description,
		&i.Sideeffects,
		&i.CreatedAt,
	)
	return i, err
}

const getMedicationByName = `-- name: GetMedicationByName :one
SELECT id, medicationame, description, sideeffects, created_at FROM "Medication"
WHERE medicationame = $1 LIMIT 1
`

func (q *Queries) GetMedicationByName(ctx context.Context, medicationame string) (Medication, error) {
	row := q.db.QueryRow(ctx, getMedicationByName, medicationame)
	var i Medication
	err := row.Scan(
		&i.ID,
		&i.Medicationame,
		&i.Description,
		&i.Sideeffects,
		&i.CreatedAt,
	)
	return i, err
}
