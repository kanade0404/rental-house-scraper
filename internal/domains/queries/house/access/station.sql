-- name: CreateStation :one
INSERT INTO station (name)
VALUES ($1)
    RETURNING name;
