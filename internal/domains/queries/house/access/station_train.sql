-- name: CreateStationTrain :one
INSERT INTO station_train (station, train)
SELECT $1, $2
WHERE NOT EXISTS (
    SELECT 1
    FROM station_train
    WHERE station = $1 AND train = $2
)
RETURNING station, train;
