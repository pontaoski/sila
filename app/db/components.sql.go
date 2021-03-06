// Code generated by sqlc. DO NOT EDIT.
// source: components.sql

package db

import (
	"context"
	"database/sql"
)

const countComponentBugs = `-- name: CountComponentBugs :many
SELECT COUNT(*) FROM Bugs
    WHERE Component_ID=$1
`

func (q *Queries) CountComponentBugs(ctx context.Context, componentID int64) ([]int64, error) {
	rows, err := q.query(ctx, q.countComponentBugsStmt, countComponentBugs, componentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int64
	for rows.Next() {
		var count int64
		if err := rows.Scan(&count); err != nil {
			return nil, err
		}
		items = append(items, count)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createComponent = `-- name: CreateComponent :one
INSERT INTO Components (
    Name, Component_Description, Product_ID
) VALUES (
    $1, $2, $3
)
RETURNING id, name, component_description, product_id
`

type CreateComponentParams struct {
	Name                 sql.NullString `json:"name"`
	ComponentDescription sql.NullString `json:"component_description"`
	ProductID            int64          `json:"product_id"`
}

func (q *Queries) CreateComponent(ctx context.Context, arg CreateComponentParams) (Component, error) {
	row := q.queryRow(ctx, q.createComponentStmt, createComponent, arg.Name, arg.ComponentDescription, arg.ProductID)
	var i Component
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ComponentDescription,
		&i.ProductID,
	)
	return i, err
}

const deleteComponent = `-- name: DeleteComponent :exec
DELETE FROM Components
    WHERE ID=$1
`

func (q *Queries) DeleteComponent(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteComponentStmt, deleteComponent, id)
	return err
}

const editComponent = `-- name: EditComponent :exec
UPDATE Components
    SET Name = $2,
        Component_Description = $3
    WHERE ID = $1
`

type EditComponentParams struct {
	ID                   int64          `json:"id"`
	Name                 sql.NullString `json:"name"`
	ComponentDescription sql.NullString `json:"component_description"`
}

func (q *Queries) EditComponent(ctx context.Context, arg EditComponentParams) error {
	_, err := q.exec(ctx, q.editComponentStmt, editComponent, arg.ID, arg.Name, arg.ComponentDescription)
	return err
}

const listComponents = `-- name: ListComponents :many
SELECT id, name, component_description, product_id FROM Components
    ORDER BY ID
    LIMIT $1 OFFSET $2
`

type ListComponentsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListComponents(ctx context.Context, arg ListComponentsParams) ([]Component, error) {
	rows, err := q.query(ctx, q.listComponentsStmt, listComponents, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Component
	for rows.Next() {
		var i Component
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ComponentDescription,
			&i.ProductID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
