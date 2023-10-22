-- name: CreateHouse :one
INSERT INTO house (name, construction_year, structure, address_id, owner_id)
VALUES ($1, $2, $3, $4, $5)
    RETURNING id, name, construction_year, structure, address_id, owner_id;
