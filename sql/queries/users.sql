-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, hashed_password, email, is_chirpy_red)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3
)
RETURNING *;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: UpdateUser :one
UPDATE users
SET email = $2,
    hashed_password = $3,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpgradeUser :one
UPDATE users
set is_chirpy_red = TRUE, updated_at = NOW()
WHERE id = $1
RETURNING *;