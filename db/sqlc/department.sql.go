// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: department.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createDepartment = `-- name: CreateDepartment :one
INSERT INTO "Department" (
name,description,head_id
) VALUES (
  $1,$2,$3
)
RETURNING id, name, description, head_id, created_at
`

type CreateDepartmentParams struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	HeadID      pgtype.Int8 `json:"head_id"`
}

func (q *Queries) CreateDepartment(ctx context.Context, arg CreateDepartmentParams) (Department, error) {
	row := q.db.QueryRow(ctx, createDepartment, arg.Name, arg.Description, arg.HeadID)
	var i Department
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.HeadID,
		&i.CreatedAt,
	)
	return i, err
}

const getDepartmentByHead = `-- name: GetDepartmentByHead :one
SELECT id, name, description, head_id, created_at FROM "Department"
WHERE head_id = $1 LIMIT 1
`

func (q *Queries) GetDepartmentByHead(ctx context.Context, headID pgtype.Int8) (Department, error) {
	row := q.db.QueryRow(ctx, getDepartmentByHead, headID)
	var i Department
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.HeadID,
		&i.CreatedAt,
	)
	return i, err
}

const getDepartmentById = `-- name: GetDepartmentById :one
SELECT id, name, description, head_id, created_at FROM "Department"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetDepartmentById(ctx context.Context, id int64) (Department, error) {
	row := q.db.QueryRow(ctx, getDepartmentById, id)
	var i Department
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.HeadID,
		&i.CreatedAt,
	)
	return i, err
}

const getDepartmentByName = `-- name: GetDepartmentByName :one
SELECT id, name, description, head_id, created_at FROM "Department"
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetDepartmentByName(ctx context.Context, name string) (Department, error) {
	row := q.db.QueryRow(ctx, getDepartmentByName, name)
	var i Department
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.HeadID,
		&i.CreatedAt,
	)
	return i, err
}

const getDepartmentHeadId = `-- name: GetDepartmentHeadId :one
SELECT head_id FROM "Department"
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetDepartmentHeadId(ctx context.Context, name string) (pgtype.Int8, error) {
	row := q.db.QueryRow(ctx, getDepartmentHeadId, name)
	var head_id pgtype.Int8
	err := row.Scan(&head_id)
	return head_id, err
}
