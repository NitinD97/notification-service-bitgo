-- name: CreateUser :one
INSERT INTO users (name, email)
VALUES ($1, $2)
RETURNING *;

-- name: GetUserById :one
SELECT *
from users
where id = $1;

-- name: TruncateUsers :exec
TRUNCATE TABLE users CASCADE;