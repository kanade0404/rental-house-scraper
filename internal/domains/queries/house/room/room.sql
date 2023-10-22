-- name: CreateRoom :one
INSERT INTO room (name, house_id)
SELECT $1, $2
WHERE NOT EXISTS (
    SELECT name
    FROM room
    WHERE name = $1 AND house_id = $2
)
    RETURNING id, name, house_id;
