-- name: InsertRolePermission :exec
INSERT INTO role_permissions (role_id, permission_id)
VALUES ($1, $2)
    ON CONFLICT DO NOTHING;
