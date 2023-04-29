// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: qInstances.sql

package db

import (
	"context"

	"github.com/jackc/pgtype"
)

const countInstances = `-- name: CountInstances :one
SELECT COUNT(*)
FROM filter_instances
WHERE (user_id = $1 AND template_name = $2)
`

type CountInstancesParams struct {
	UserID       string
	TemplateName string
}

func (q *Queries) CountInstances(ctx context.Context, arg CountInstancesParams) (int64, error) {
	row := q.db.QueryRow(ctx, countInstances, arg.UserID, arg.TemplateName)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createInstance = `-- name: CreateInstance :exec
INSERT INTO filter_instances (list_id, user_id, template_name, params, test_mode)
VALUES ((SELECT id FROM filter_lists WHERE user_id = $1), $1, $2, $3, $4)
`

type CreateInstanceParams struct {
	UserID       string
	TemplateName string
	Params       pgtype.JSONB
	TestMode     bool
}

func (q *Queries) CreateInstance(ctx context.Context, arg CreateInstanceParams) error {
	_, err := q.db.Exec(ctx, createInstance,
		arg.UserID,
		arg.TemplateName,
		arg.Params,
		arg.TestMode,
	)
	return err
}

const deleteInstance = `-- name: DeleteInstance :exec
DELETE
FROM filter_instances
WHERE (user_id = $1 AND template_name = $2)
`

type DeleteInstanceParams struct {
	UserID       string
	TemplateName string
}

func (q *Queries) DeleteInstance(ctx context.Context, arg DeleteInstanceParams) error {
	_, err := q.db.Exec(ctx, deleteInstance, arg.UserID, arg.TemplateName)
	return err
}

const getInstance = `-- name: GetInstance :one
SELECT params, test_mode
FROM filter_instances
WHERE (user_id = $1 AND template_name = $2)
`

type GetInstanceParams struct {
	UserID       string
	TemplateName string
}

type GetInstanceRow struct {
	Params   pgtype.JSONB
	TestMode bool
}

func (q *Queries) GetInstance(ctx context.Context, arg GetInstanceParams) (GetInstanceRow, error) {
	row := q.db.QueryRow(ctx, getInstance, arg.UserID, arg.TemplateName)
	var i GetInstanceRow
	err := row.Scan(&i.Params, &i.TestMode)
	return i, err
}

const getInstancesForList = `-- name: GetInstancesForList :many
SELECT template_name, params, test_mode
FROM filter_instances
WHERE list_id = $1
ORDER BY template_name ASC
`

type GetInstancesForListRow struct {
	TemplateName string
	Params       pgtype.JSONB
	TestMode     bool
}

func (q *Queries) GetInstancesForList(ctx context.Context, listID int32) ([]GetInstancesForListRow, error) {
	rows, err := q.db.Query(ctx, getInstancesForList, listID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetInstancesForListRow
	for rows.Next() {
		var i GetInstancesForListRow
		if err := rows.Scan(&i.TemplateName, &i.Params, &i.TestMode); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getInstancesForUser = `-- name: GetInstancesForUser :many
SELECT template_name, params, test_mode
FROM filter_instances
WHERE user_id = $1
`

type GetInstancesForUserRow struct {
	TemplateName string
	Params       pgtype.JSONB
	TestMode     bool
}

func (q *Queries) GetInstancesForUser(ctx context.Context, userID string) ([]GetInstancesForUserRow, error) {
	rows, err := q.db.Query(ctx, getInstancesForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetInstancesForUserRow
	for rows.Next() {
		var i GetInstancesForUserRow
		if err := rows.Scan(&i.TemplateName, &i.Params, &i.TestMode); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateInstance = `-- name: UpdateInstance :exec
UPDATE filter_instances
SET params     = $3,
    test_mode  = $4,
    updated_at = NOW()
WHERE (user_id = $1 AND template_name = $2)
`

type UpdateInstanceParams struct {
	UserID       string
	TemplateName string
	Params       pgtype.JSONB
	TestMode     bool
}

func (q *Queries) UpdateInstance(ctx context.Context, arg UpdateInstanceParams) error {
	_, err := q.db.Exec(ctx, updateInstance,
		arg.UserID,
		arg.TemplateName,
		arg.Params,
		arg.TestMode,
	)
	return err
}
