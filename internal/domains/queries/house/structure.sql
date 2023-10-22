-- name: CreateStructure :one
INSERT INTO structure (name)
VALUES ($1)
    RETURNING name;
