// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createRule = `-- name: CreateRule :exec
INSERT INTO rules (
    name, description, rule, ast
) VALUES (
    $1, $2, $3, $4
)
`

type CreateRuleParams struct {
	Name        string
	Description pgtype.Text
	Rule        string
	Ast         []byte
}

func (q *Queries) CreateRule(ctx context.Context, arg CreateRuleParams) error {
	_, err := q.db.Exec(ctx, createRule,
		arg.Name,
		arg.Description,
		arg.Rule,
		arg.Ast,
	)
	return err
}

const getRule = `-- name: GetRule :one
SELECT id, name, description, rule, ast, created_at, updated_at FROM rules
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRule(ctx context.Context, id int32) (Rule, error) {
	row := q.db.QueryRow(ctx, getRule, id)
	var i Rule
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Rule,
		&i.Ast,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRules = `-- name: GetRules :many
SELECT id, name, description, rule, ast, created_at, updated_at FROM rules
`

func (q *Queries) GetRules(ctx context.Context) ([]Rule, error) {
	rows, err := q.db.Query(ctx, getRules)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Rule
	for rows.Next() {
		var i Rule
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Rule,
			&i.Ast,
			&i.CreatedAt,
			&i.UpdatedAt,
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
