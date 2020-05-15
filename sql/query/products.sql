-- name: CreateProduct :one
INSERT INTO Products (
    Name, Product_Description, Active
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: EditProduct :exec
UPDATE Products
    SET Name = $2,
        Product_Description = $3
    WHERE ID = $1;

-- name: DeleteProduct :exec
DELETE FROM Products
    WHERE ID=$1;