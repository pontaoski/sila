// Code generated by sqlc. DO NOT EDIT.
// source: dependencies.sql

package db

import (
	"context"
)

const addDependency = `-- name: AddDependency :exec
INSERT INTO Bug_Dependencies (
    Required_By, Depends_On
) VALUES (
    $1, $2
)
`

type AddDependencyParams struct {
	RequiredBy int64 `json:"required_by"`
	DependsOn  int64 `json:"depends_on"`
}

func (q *Queries) AddDependency(ctx context.Context, arg AddDependencyParams) error {
	_, err := q.exec(ctx, q.addDependencyStmt, addDependency, arg.RequiredBy, arg.DependsOn)
	return err
}

const clearDependenciesForBug = `-- name: ClearDependenciesForBug :exec
DELETE FROM Bug_Dependencies
    WHERE Required_By=$1
`

func (q *Queries) ClearDependenciesForBug(ctx context.Context, requiredBy int64) error {
	_, err := q.exec(ctx, q.clearDependenciesForBugStmt, clearDependenciesForBug, requiredBy)
	return err
}

const clearDependenciesOnBug = `-- name: ClearDependenciesOnBug :exec
DELETE FROM Bug_Dependencies
    WHERE Depends_On=$1
`

func (q *Queries) ClearDependenciesOnBug(ctx context.Context, dependsOn int64) error {
	_, err := q.exec(ctx, q.clearDependenciesOnBugStmt, clearDependenciesOnBug, dependsOn)
	return err
}

const removeDependency = `-- name: RemoveDependency :exec
DELETE FROM Bug_Dependencies
    WHERE Required_By=$1
    AND Depends_On=$2
`

type RemoveDependencyParams struct {
	RequiredBy int64 `json:"required_by"`
	DependsOn  int64 `json:"depends_on"`
}

func (q *Queries) RemoveDependency(ctx context.Context, arg RemoveDependencyParams) error {
	_, err := q.exec(ctx, q.removeDependencyStmt, removeDependency, arg.RequiredBy, arg.DependsOn)
	return err
}
