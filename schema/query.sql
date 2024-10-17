-- name: GetRule :one
SELECT * FROM rules
WHERE id = $1 LIMIT 1;

-- name: GetRules :many
SELECT * FROM rules;

-- name: CreateRule :exec
INSERT INTO rules (
    name, description, rule, ast
) VALUES (
    $1, $2, $3, $4
);
