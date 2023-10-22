-- name: CreateAddress :one
INSERT INTO address (name, city)
SELECT $1, $2
WHERE NOT EXISTS (
    SELECT name, city
    FROM address
    WHERE name = $1 AND city = $2
)
    RETURNING name, city;
