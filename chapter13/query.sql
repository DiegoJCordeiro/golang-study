-- name: QueryAllCategories :many
SELECT * FROM Categories;

-- name: QueryOneCategory :one
SELECT * FROM Categories WHERE Name = ? LIMIT 1;

-- name: InsertOneCategory :exec
INSERT INTO Categories(Id, Description, Name) Values (?, ?, ?)