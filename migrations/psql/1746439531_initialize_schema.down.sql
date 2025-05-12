-- DROP TRIGGERS
DO $$
    DECLARE
        t_name TEXT;
    BEGIN
        -- === USERS ===
        t_name := 'trg_user_activate_insert';
        IF EXISTS (
            SELECT 1 FROM pg_trigger trg
                              JOIN pg_class tbl ON trg.tgrelid = tbl.oid
            WHERE trg.tgname = t_name AND tbl.relname = 'users'
        ) THEN
            EXECUTE format('DROP TRIGGER %I ON users', t_name);
        END IF;
    END $$;


DROP FUNCTION IF EXISTS trg_user_activate_insert();

-- DROP TABLES IN REVERSE ORDER
DROP TABLE IF EXISTS user_login_logs;
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS user_roles;
DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS user_settings;
DROP TABLE IF EXISTS user_profiles;
DROP TABLE IF EXISTS files;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS schema_migrations;


-- DROP ENUM TYPES
DROP TYPE IF EXISTS users_status_enum CASCADE;
DROP TYPE IF EXISTS user_settings_theme_enum CASCADE;
DROP TYPE IF EXISTS user_files_storage_type_enum CASCADE;
