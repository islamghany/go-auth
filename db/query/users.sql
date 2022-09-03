-- name: CreateUser :one
INSERT INTO users(name, email, hashed_password, activated)
VALUES ($1,$2,$3,$4)
RETURNING id, created_at, version;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;