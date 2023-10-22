-- name: CreateFloorPlan :one
INSERT INTO floor_plan (name, floor_plan_type)
SELECT $1, $2
WHERE NOT EXISTS (
    SELECT name, floor_plan_type
    FROM floor_plan
    WHERE name = $1 AND floor_plan_type = $2
)
    RETURNING name, floor_plan_type;
