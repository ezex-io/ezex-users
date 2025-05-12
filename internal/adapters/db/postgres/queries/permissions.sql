-- name: UpsertPermission :one
INSERT INTO permissions (name, description, scope, action)
VALUES ($1, $2, $3, $4)
ON CONFLICT (scope, action) DO UPDATE
    SET name = EXCLUDED.name,
        description = EXCLUDED.description
RETURNING *;

-- name: GetPermissionByScopeAction :one
SELECT * FROM permissions
WHERE scope = $1 AND action = $2;
