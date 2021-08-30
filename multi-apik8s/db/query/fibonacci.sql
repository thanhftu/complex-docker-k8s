-- name: CreateFibsTable :exec
CREATE TABLE fibs (
    id bigserial PRIMARY KEY,
    index bigint UNIQUE NOT NULL,
    value  bigint NOT NULL,
    create_at timestamptz NOT NULL DEFAULT (now())
);
-- name: CreateFibonacci :one
INSERT INTO fibs (
  index, value
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetFibonacciByIndex :one
SELECT * FROM fibs
WHERE index = $1 LIMIT 1;

-- name: GetFibonacciByID :one
SELECT * FROM fibs
WHERE id = $1 LIMIT 1;

-- name: GetLatestCreatedFibonacci :one
SELECT * FROM fibs
ORDER BY id DESC LIMIT 1;

-- name: ListFibonaccis :many
SELECT * FROM fibs
ORDER BY id DESC
LIMIT $1;

-- name: DeleteFibonacci :exec
DELETE FROM fibs WHERE id = $1;