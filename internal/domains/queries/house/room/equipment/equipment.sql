-- name: CreateEquipment :one
INSERT INTO equipment (name)
SELECT $1
WHERE NOT EXISTS (
    SELECT name
    FROM equipment
    WHERE name = $1
)
    RETURNING name;

