// Code generated by sqlc. DO NOT EDIT.
// source: fibonacci.sql

package db

import (
	"context"
)

const createFibonacci = `-- name: CreateFibonacci :one
INSERT INTO fibs (
  index, value
) VALUES (
  $1, $2
)
RETURNING id, index, value, create_at
`

type CreateFibonacciParams struct {
	Index int64 `json:"index"`
	Value int64 `json:"value"`
}

func (q *Queries) CreateFibonacci(ctx context.Context, arg CreateFibonacciParams) (Fib, error) {
	row := q.db.QueryRowContext(ctx, createFibonacci, arg.Index, arg.Value)
	var i Fib
	err := row.Scan(
		&i.ID,
		&i.Index,
		&i.Value,
		&i.CreateAt,
	)
	return i, err
}

const deleteFibonacci = `-- name: DeleteFibonacci :exec
DELETE FROM fibs WHERE id = $1
`

func (q *Queries) DeleteFibonacci(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteFibonacci, id)
	return err
}

const getFibonacciByID = `-- name: GetFibonacciByID :one
SELECT id, index, value, create_at FROM fibs
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetFibonacciByID(ctx context.Context, id int64) (Fib, error) {
	row := q.db.QueryRowContext(ctx, getFibonacciByID, id)
	var i Fib
	err := row.Scan(
		&i.ID,
		&i.Index,
		&i.Value,
		&i.CreateAt,
	)
	return i, err
}

const getFibonacciByIndex = `-- name: GetFibonacciByIndex :one
SELECT id, index, value, create_at FROM fibs
WHERE index = $1 LIMIT 1
`

func (q *Queries) GetFibonacciByIndex(ctx context.Context, index int64) (Fib, error) {
	row := q.db.QueryRowContext(ctx, getFibonacciByIndex, index)
	var i Fib
	err := row.Scan(
		&i.ID,
		&i.Index,
		&i.Value,
		&i.CreateAt,
	)
	return i, err
}

const getLatestCreatedFibonacci = `-- name: GetLatestCreatedFibonacci :one
SELECT id, index, value, create_at FROM fibs
ORDER BY id DESC LIMIT 1
`

func (q *Queries) GetLatestCreatedFibonacci(ctx context.Context) (Fib, error) {
	row := q.db.QueryRowContext(ctx, getLatestCreatedFibonacci)
	var i Fib
	err := row.Scan(
		&i.ID,
		&i.Index,
		&i.Value,
		&i.CreateAt,
	)
	return i, err
}

const listFibonaccis = `-- name: ListFibonaccis :many
SELECT id, index, value, create_at FROM fibs
ORDER BY id
LIMIT $1
`

func (q *Queries) ListFibonaccis(ctx context.Context, limit int32) ([]Fib, error) {
	rows, err := q.db.QueryContext(ctx, listFibonaccis, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Fib{}
	for rows.Next() {
		var i Fib
		if err := rows.Scan(
			&i.ID,
			&i.Index,
			&i.Value,
			&i.CreateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}