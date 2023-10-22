-- name: CreateFloorPlanType :one
INSERT INTO floor_plan_type (name)
SELECT $1
WHERE NOT EXISTS (
    SELECT name
    FROM floor_plan_type
    WHERE name = $1
)
    RETURNING name;
