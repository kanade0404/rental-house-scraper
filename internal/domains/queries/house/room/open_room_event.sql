-- name: CreateOpenRoomEvent :one
INSERT INTO open_room_event (room_id, recorded_at)
SELECT $1, $2
WHERE NOT EXISTS (
    SELECT room_id
    FROM open_room_event
    WHERE room_id = $1 AND recorded_at = $2
)
    RETURNING room_id, recorded_at;
