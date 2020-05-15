-- name: CreateUser :one
INSERT INTO Users (
    Email, FullName
) VALUES (
    $1, $2
)
RETURNING *;

-- name: ListUsers :many
SELECT * FROM Users
    ORDER BY ID
    LIMIT $1 OFFSET $2;