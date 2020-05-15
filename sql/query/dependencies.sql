-- name: AddDependency :exec
INSERT INTO Bug_Dependencies (
    Required_By, Depends_On
) VALUES (
    $1, $2
);

-- name: RemoveDependency :exec
DELETE FROM Bug_Dependencies
    WHERE Required_By=$1
    AND Depends_On=$2;

-- name: ClearDependenciesForBug :exec
DELETE FROM Bug_Dependencies
    WHERE Required_By=$1;

-- name: ClearDependenciesOnBug :exec
DELETE FROM Bug_Dependencies
    WHERE Depends_On=$1;