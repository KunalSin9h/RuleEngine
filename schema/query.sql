-- name: GetRule :one
SELECT * FROM rules
WHERE id = $1 LIMIT 1;

-- name: GetRules :many
SELECT * FROM rules;
