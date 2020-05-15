-- name: CreateBug :one
INSERT INTO Bugs (
    Title, Bug_Description, Component_ID, Author_ID, Created_At
) VALUES (
    $1, $2, $3, $4, NOW()
)
RETURNING *;

-- name: DeleteBug :exec
DELETE FROM Bugs
    WHERE ID = $1;

-- name: WhatRequires :many
SELECT *
FROM Bugs
INNER JOIN Bug_Dependencies
    ON Bugs.ID = Bug_Dependencies.Depends_On
    WHERE Bugs.ID = $1;

-- name: Requires :many
SELECT *
FROM Bugs
INNER JOIN Bug_Dependencies
    ON Bugs.ID = Bug_Dependencies.Required_By
    WHERE Bugs.ID = $1;

-- name: CountBugs :one
SELECT count(*) FROM Bugs;

-- name: ListBugs :many
SELECT * FROM Bugs
    ORDER BY ID
    LIMIT $1 OFFSET $2;

-- name: EditDescription :exec
UPDATE Bugs
    SET Bug_Description = $2,
        Edited_At = NOW()
    WHERE ID = $1;

-- name: EditSeverity :exec
UPDATE Bugs
    SET Bug_Severity = $2,
        Edited_At = NOW()
    WHERE ID = $1;

-- name: EditStatus :exec
UPDATE Bugs
    SET Bug_Status = $2,
        Edited_At = NOW()
    WHERE ID = $1;

-- name: RelocateBug :exec
UPDATE Bugs
    SET Component_ID = $2,
        Edited_At = NOW()
    WHERE ID = $1;