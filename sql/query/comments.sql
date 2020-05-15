-- name: AddComment :one
INSERT INTO Comments (
    Bug_ID, Author_ID, Comment_Text, Created_At
) VALUES (
    $1, $2, $3, NOW()
)
RETURNING *;

-- name: EditComment :one
UPDATE Comments
    SET Comment_Text = $2,
        Edited_At = NOW()
    WHERE ID = $1
    RETURNING *;

-- name: RedactComment :exec
UPDATE Comments
SET Redacted = TRUE;