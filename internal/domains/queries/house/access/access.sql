-- name: CreateAccess :one
INSERT INTO access (station, movement_method, time_in_minutes, house_id)
SELECT $1, $2, $3, $4
WHERE NOT EXISTS (
    SELECT 1
    FROM access
    WHERE station = $1 AND movement_method = $2 AND time_in_minutes = $3 AND house_id = $4
)
RETURNING station, movement_method, time_in_minutes, house_id;

