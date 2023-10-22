-- name: CreateMovementMethod :one
INSERT INTO movement_method (name)
VALUES ($1)
    RETURNING name;
