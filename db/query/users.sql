-- name: CreateUser :one
INSERT INTO users(name, email, hashed_password, activated)
VALUES ($1,$2,$3,$4)
RETURNING id, created_at, version;

-- name: GetUserEmail :one
SELECT * FROM users
WHERE email = $1; 

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserFromToken :one
select users.* from tokens join users on users.id = tokens.user_id
where tokens.hashed_token = $1 and tokens.scope = $2 and tokens.expiry > $3;
