-- ENUM TYPES
CREATE TYPE user_settings_theme_enum AS ENUM ('LIGHT', 'DARK', 'AUTO');
CREATE TYPE users_status_enum AS ENUM ('PENDING', 'ACTIVE', 'INACTIVE', 'SUSPENDED', 'DELETED');
CREATE TYPE user_files_storage_type_enum AS ENUM ('LOCAL', 'FTP', 'S3');

-- USERS TABLE
CREATE TABLE users (
                       id TEXT PRIMARY KEY DEFAULT idkit_ksuid_generate_text(),
                       email TEXT NOT NULL UNIQUE,
                       username TEXT NOT NULL UNIQUE,
                       password TEXT,
                       status users_status_enum NOT NULL,
                       created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                       created_by_id TEXT REFERENCES users(id) ON DELETE SET NULL,
                       updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                       updated_by_id TEXT REFERENCES users(id) ON DELETE SET NULL,
                       deleted_at TIMESTAMPTZ,
                       deleted_by_id TEXT REFERENCES users(id) ON DELETE SET NULL
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_status ON users(status);
CREATE INDEX idx_users_created_at ON users(created_at);
CREATE INDEX idx_users_created_by_id ON users(created_by_id);
CREATE INDEX idx_users_updated_by_id ON users(updated_by_id);
CREATE INDEX idx_users_deleted_by_id ON users(deleted_by_id);

-- FILES TABLE
CREATE TABLE files (
                       id TEXT PRIMARY KEY DEFAULT idkit_ksuid_generate_text(),
                       user_id TEXT REFERENCES users(id) ON DELETE CASCADE,
                       storage_type user_files_storage_type_enum NOT NULL,
                       bucket TEXT,
                       object_key TEXT NOT NULL,
                       content_type TEXT NOT NULL,
                       size BIGINT NOT NULL,
                       category TEXT,
                       is_public BOOLEAN NOT NULL DEFAULT FALSE,
                       uploaded_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                       uploaded_by_id TEXT REFERENCES users(id) ON DELETE SET NULL,
                       updated_at TIMESTAMPTZ NOT NULL  DEFAULT now(),
                       updated_by_id TEXT REFERENCES users(id) ON DELETE SET NULL,
                       deleted_at TIMESTAMPTZ,
                       deleted_by_id TEXT REFERENCES users(id) ON DELETE SET NULL
);

CREATE INDEX idx_files_user_id ON files(user_id);
CREATE INDEX idx_files_storage_key ON files(storage_type, object_key);
CREATE INDEX idx_files_public_cat ON files(is_public, category);
CREATE INDEX idx_files_size ON files(size);
CREATE INDEX idx_files_uploaded_at ON files(uploaded_at);
CREATE INDEX idx_files_created_by_id ON files(uploaded_by_id);
CREATE INDEX idx_files_updated_by_id ON files(updated_by_id);
CREATE INDEX idx_files_deleted_by_id ON files(deleted_by_id);

-- USER_PROFILES TABLE
CREATE TABLE user_profiles (
                               user_id TEXT PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
                               first_name TEXT,
                               last_name TEXT,
                               avatar_file_id TEXT REFERENCES files(id) ON DELETE SET NULL,
                               updated_by_id TEXT REFERENCES users(id) ON DELETE SET NULL,
                               updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_user_profiles_name ON user_profiles(first_name, last_name);
CREATE INDEX idx_user_profiles_updated_by_id ON user_profiles(updated_by_id);

-- USER SETTINGS
CREATE TABLE user_settings (
                               user_id TEXT PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
                               language TEXT NOT NULL,
                               theme user_settings_theme_enum NOT NULL DEFAULT 'LIGHT',
                               timezone TEXT,
                               updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_user_settings_language ON user_settings(language);
CREATE INDEX idx_user_settings_theme ON user_settings(theme);

-- ROLES TABLE
CREATE TABLE roles (
                       id TEXT PRIMARY KEY DEFAULT idkit_ksuid_generate_text(),
                       name TEXT NOT NULL UNIQUE,
                       is_system BOOLEAN NOT NULL DEFAULT FALSE,
                       is_default BOOLEAN NOT NULL DEFAULT FALSE,
                       created_by_id TEXT REFERENCES users(id) ON DELETE SET NULL,
                       created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                       updated_by_id TEXT REFERENCES users(id) ON DELETE SET NULL,
                       updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                       deleted_by_id TEXT REFERENCES users(id) ON DELETE SET NULL,
                       deleted_at TIMESTAMPTZ
);

COMMENT ON COLUMN roles.is_default IS
    'Indicates the default role for new user registrations.';

CREATE UNIQUE INDEX idx_roles_name ON roles(name);
CREATE INDEX idx_roles_is_default ON roles(is_default);
CREATE INDEX idx_roles_is_system ON roles(is_system);
CREATE INDEX idx_roles_created_at ON roles(created_at);
CREATE INDEX idx_roles_created_by_id ON roles(created_by_id);
CREATE INDEX idx_roles_updated_by_id ON roles(updated_by_id);
CREATE INDEX idx_roles_deleted_by_id ON roles(deleted_by_id);

-- PERMISSIONS TABLE
CREATE TABLE permissions (
                             id TEXT PRIMARY KEY DEFAULT idkit_ksuid_generate_text(),
                             name TEXT,
                             description TEXT,
                             scope TEXT NOT NULL,
                             action TEXT NOT NULL
);

CREATE UNIQUE INDEX idx_permissions_scope_action ON permissions(scope, action);
CREATE INDEX idx_permissions_scope ON permissions(scope);

-- ROLE_PERMISSIONS TABLE
CREATE TABLE role_permissions (
                                  role_id TEXT NOT NULL,
                                  permission_id TEXT NOT NULL,
                                  PRIMARY KEY (role_id, permission_id),
                                  FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
                                  FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE CASCADE
);

CREATE INDEX idx_role_permissions_permission_id ON role_permissions(permission_id);

-- USER_ROLES TABLE
CREATE TABLE user_roles (
                            user_id TEXT NOT NULL,
                            role_id TEXT NOT NULL,
                            granted_by_id TEXT NOT NULL REFERENCES users(id) ON DELETE SET NULL,
                            granted_at TIMESTAMPTZ NOT NULL DEFAULT now(),
                            PRIMARY KEY (user_id, role_id),
                            FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
                            FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);

CREATE INDEX idx_user_roles_user_id ON user_roles(user_id);
CREATE INDEX idx_user_roles_role_id ON user_roles(role_id);
CREATE INDEX idx_user_roles_granted_by_id ON user_roles(granted_by_id);

-- SESSIONS TABLE
CREATE TABLE sessions (
                          id TEXT PRIMARY KEY DEFAULT idkit_ksuid_generate_text(),
                          user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                          refresh_token TEXT NOT NULL UNIQUE DEFAULT idkit_ksuid_generate_text(),
                          device_id TEXT NOT NULL,
                          device_name TEXT NOT NULL,
                          country TEXT,
                          city TEXT,
                          user_agent TEXT,
                          ip_address TEXT,
                          is_active BOOLEAN NOT NULL DEFAULT TRUE,
                          session_expires_at TIMESTAMPTZ NOT NULL,
                          created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_sessions_user_id ON sessions(user_id);
CREATE UNIQUE INDEX idx_sessions_refresh_token ON sessions(refresh_token);
CREATE INDEX idx_sessions_expires_at ON sessions(session_expires_at);
CREATE INDEX idx_sessions_device_id ON sessions(device_id);
CREATE INDEX idx_sessions_created_at ON sessions(created_at);

-- USER LOGIN LOGS
CREATE TABLE user_login_logs (
                                 id BIGSERIAL PRIMARY KEY,
                                 user_id TEXT REFERENCES users(id) ON DELETE SET NULL,
                                 email TEXT,
                                 ip_address TEXT,
                                 country TEXT,
                                 city TEXT,
                                 user_agent TEXT,
                                 is_success BOOLEAN NOT NULL,
                                 reason TEXT NOT NULL,
                                 created_at TIMESTAMPTZ
);

CREATE INDEX idx_user_login_logs_user_id ON user_login_logs(user_id);
CREATE INDEX idx_user_login_logs_ip ON user_login_logs(ip_address);
CREATE INDEX idx_user_login_logs_created_at ON user_login_logs(created_at);
CREATE INDEX idx_user_login_logs_success ON user_login_logs(is_success);

-- TRIGGER create user profile
CREATE OR REPLACE FUNCTION trg_user_activate_insert()
    RETURNS TRIGGER AS $$
BEGIN
    IF NEW.status = 'ACTIVE' THEN
        INSERT INTO user_profiles (user_id) VALUES (NEW.id);
        INSERT INTO user_settings (user_id, language, theme, timezone)
        VALUES (NEW.id, 'en', 'LIGHT', 'UTC');
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_user_activate_insert
    AFTER INSERT ON users
    FOR EACH ROW
EXECUTE FUNCTION trg_user_activate_insert();
