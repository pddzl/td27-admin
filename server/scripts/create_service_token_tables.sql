-- Create service token tables
-- Run this SQL in your PostgreSQL database

-- Service Token table
CREATE TABLE IF NOT EXISTS sys_tool_service_token (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    name VARCHAR(100) NOT NULL,
    token_hash VARCHAR(255) UNIQUE NOT NULL,
    status BOOLEAN DEFAULT TRUE,
    expires_at BIGINT
);

CREATE INDEX IF NOT EXISTS idx_sys_tool_service_token_deleted_at ON sys_tool_service_token(deleted_at);
CREATE INDEX IF NOT EXISTS idx_sys_tool_service_token_token_hash ON sys_tool_service_token(token_hash);

-- Token Permission association table
CREATE TABLE IF NOT EXISTS sys_tool_token_permission (
    token_id BIGINT NOT NULL,
    permission_id BIGINT NOT NULL,
    PRIMARY KEY (token_id, permission_id),
    CONSTRAINT fk_token_permission_token FOREIGN KEY (token_id) REFERENCES sys_tool_service_token(id) ON DELETE CASCADE,
    CONSTRAINT fk_token_permission_perm FOREIGN KEY (permission_id) REFERENCES sys_management_permission(id) ON DELETE CASCADE
);

-- Add menu for Service Token (under System Tools)
INSERT INTO sys_management_menu (menu_name, icon, path, component, parent_id, sort, hidden, status, title) 
SELECT '服务令牌', 'key', '/systool/serviceToken', 'sysTool/service_token/index.vue', 20, 3, false, true, '服务令牌'
WHERE NOT EXISTS (SELECT 1 FROM sys_management_menu WHERE path = '/systool/serviceToken');
