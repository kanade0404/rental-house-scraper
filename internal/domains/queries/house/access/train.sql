-- name: CreateTrain :one
INSERT INTO train (name)
VALUES ($1)
    RETURNING name;
