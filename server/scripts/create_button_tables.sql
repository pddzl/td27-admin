-- Create button permission tables

CREATE TABLE IF NOT EXISTS sys_management_button (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    button_code VARCHAR(100) NOT NULL UNIQUE,
    button_name VARCHAR(100) NOT NULL,
    description VARCHAR(200),
    page_path VARCHAR(200) NOT NULL
);

CREATE INDEX idx_sys_management_button_page_path ON sys_management_button(page_path);
CREATE INDEX idx_sys_management_button_deleted_at ON sys_management_button(deleted_at);

-- Add menu for Button Management
INSERT INTO sys_management_menu (menu_name, icon, path, component, parent_id, sort, status, title) 
SELECT '按钮权限', 'mouse', '/sysManagement/button', 'sysManagement/button/index.vue', 1, 7, true, '按钮权限'
WHERE NOT EXISTS (SELECT 1 FROM sys_management_menu WHERE path = '/sysManagement/button');
