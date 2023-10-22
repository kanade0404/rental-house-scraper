-- name: CreateRoomMetadata :one
INSERT INTO room_metadata (room_id, direction, area_square_meter, floor, floor_plan)
SELECT $1, $2, $3, $4, $5
WHERE NOT EXISTS (
    SELECT room_id
    FROM room_metadata
    WHERE room_id = $1
)
    RETURNING room_id, direction, area_square_meter, floor, floor_plan;
