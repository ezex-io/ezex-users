-- name: GetRoleByName :one
SELECT * FROM roles
WHERE name = $1;

-- name: CreateRole :exec
INSERT INTO roles (id, name, is_system, is_default)
VALUES (gen_random_uuid(), $1, $2, $3)
ON CONFLICT (name) DO NOTHING;

-- name: CreateRoles :copyfrom
INSERT INTO roles (id, name, is_system, is_default)
VALUES ($1, $2, $3, $4);

-- name: GrantRole :exec
INSERT INTO user_roles (user_id, role_id, granted_by_id, granted_at)
SELECT
    $1,
    $2,
    $3,
    now()
FROM users u, roles r
WHERE u.email = $4
ON CONFLICT DO NOTHING;
