-- name: CreateComponent :one
INSERT INTO Components (
    Name, Component_Description, Product_ID
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: EditComponent :exec
UPDATE Components
    SET Name = $2,
        Component_Description = $3
    WHERE ID = $1;

-- name: DeleteComponent :exec
DELETE FROM Components
    WHERE ID=$1;

-- name: CountComponentBugs :many
SELECT COUNT(*) FROM Bugs
    WHERE Component_ID=$1;