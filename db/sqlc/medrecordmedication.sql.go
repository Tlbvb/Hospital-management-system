// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: medrecordmedication.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createMedRecMedic = `-- name: CreateMedRecMedic :one
INSERT INTO "MedicalRecordMedication" (
medical_record_id,medication_id
) VALUES (
  $1, $2
)
RETURNING id, medical_record_id, medication_id, created_at
`

type CreateMedRecMedicParams struct {
	MedicalRecordID pgtype.Int8 `json:"medical_record_id"`
	MedicationID    pgtype.Int8 `json:"medication_id"`
}

func (q *Queries) CreateMedRecMedic(ctx context.Context, arg CreateMedRecMedicParams) (MedicalRecordMedication, error) {
	row := q.db.QueryRow(ctx, createMedRecMedic, arg.MedicalRecordID, arg.MedicationID)
	var i MedicalRecordMedication
	err := row.Scan(
		&i.ID,
		&i.MedicalRecordID,
		&i.MedicationID,
		&i.CreatedAt,
	)
	return i, err
}

const getMedicationByRecord = `-- name: GetMedicationByRecord :many
SELECT id, medical_record_id, medication_id, created_at FROM "MedicalRecordMedication"
WHERE medical_record_id = $1
`

func (q *Queries) GetMedicationByRecord(ctx context.Context, medicalRecordID pgtype.Int8) ([]MedicalRecordMedication, error) {
	rows, err := q.db.Query(ctx, getMedicationByRecord, medicalRecordID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []MedicalRecordMedication{}
	for rows.Next() {
		var i MedicalRecordMedication
		if err := rows.Scan(
			&i.ID,
			&i.MedicalRecordID,
			&i.MedicationID,
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
