-- name: CreateStory :one
INSERT INTO story (house_id, num)
SELECT $1, $2
FROM story
WHERE house_id = $1 AND num = $2
    RETURNING house_id, num;

-- name: CreateBasementStory :one
INSERT INTO basement_story (house_id, num)
SELECT $1, $2
FROM basement_story
WHERE house_id = $1 AND num = $2
    RETURNING house_id, num;
