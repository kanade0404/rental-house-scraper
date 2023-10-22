-- name: CreateOwner :one
INSERT INTO owner (name, tel, address_id)
SELECT $1, $2, $3
WHERE NOT EXISTS (
    SELECT name, tel, address_id
    FROM owner
    WHERE name = $1 AND tel = $2 AND address_id = $3
)
    RETURNING name, tel, address_id;
