-- name: CreateCity :one
INSERT INTO city (name, prefecture)
SELECT $1, $2
WHERE NOT EXISTS (
    SELECT name, prefecture
    FROM city
    WHERE name = $1 AND prefecture = $2
)
    RETURNING name, prefecture;
