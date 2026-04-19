--
-- PostgreSQL database dump
--

\restrict 4WWg2v0b89ELCZoC7rixwJMnlXF8OjechMDjAb7DVqne6yGHnhAy8NuEabt9kPy

-- Dumped from database version 14.22 (Debian 14.22-1.pgdg13+1)
-- Dumped by pg_dump version 18.3 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

-- *not* creating schema, since initdb creates it


ALTER SCHEMA public OWNER TO postgres;

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: sys_management_api; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_management_api (
    id bigint NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    path character varying(200) NOT NULL,
    method character varying(10) NOT NULL,
    description character varying(100),
    group_en character varying(50),
    group_cn character varying(50)
);


ALTER TABLE public.sys_management_api OWNER TO postgres;

--
-- Name: COLUMN sys_management_api.path; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_api.path IS 'API路径';


--
-- Name: COLUMN sys_management_api.method; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_api.method IS 'HTTP方法';


--
-- Name: COLUMN sys_management_api.description; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_api.description IS 'API描述';


--
-- Name: COLUMN sys_management_api.group_en; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_api.group_en IS 'API分组(英文)';


--
-- Name: COLUMN sys_management_api.group_cn; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_api.group_cn IS 'API分组(中文)';


--
-- Name: sys_management_api_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_management_api_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sys_management_api_id_seq OWNER TO postgres;

--
-- Name: sys_management_api_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_management_api_id_seq OWNED BY public.sys_management_api.id;


--
-- Name: sys_management_dept; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_management_dept (
    id bigint NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    dept_name character varying(100) NOT NULL,
    parent_id bigint,
    path character varying(500),
    sort bigint DEFAULT 0,
    status boolean DEFAULT true,
    level bigint
);


ALTER TABLE public.sys_management_dept OWNER TO postgres;

--
-- Name: COLUMN sys_management_dept.dept_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_dept.dept_name IS '部门名称';


--
-- Name: COLUMN sys_management_dept.parent_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_dept.parent_id IS '父部门ID';


--
-- Name: COLUMN sys_management_dept.path; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_dept.path IS '部门路径(materialized path),如:/1/2/3/';


--
-- Name: COLUMN sys_management_dept.level; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_dept.level IS 'depth level';


--
-- Name: sys_management_dept_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_management_dept_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sys_management_dept_id_seq OWNER TO postgres;

--
-- Name: sys_management_dept_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_management_dept_id_seq OWNED BY public.sys_management_dept.id;


--
-- Name: sys_management_dict; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_management_dict (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    cn_name text,
    en_name text
);


ALTER TABLE public.sys_management_dict OWNER TO postgres;

--
-- Name: sys_management_dict_detail; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_management_dict_detail (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    label text,
    value text,
    sort bigint,
    dict_id bigint,
    parent_id bigint,
    description text
);


ALTER TABLE public.sys_management_dict_detail OWNER TO postgres;

--
-- Name: sys_management_dict_detail_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_management_dict_detail_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sys_management_dict_detail_id_seq OWNER TO postgres;

--
-- Name: sys_management_dict_detail_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_management_dict_detail_id_seq OWNED BY public.sys_management_dict_detail.id;


--
-- Name: sys_management_dict_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_management_dict_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sys_management_dict_id_seq OWNER TO postgres;

--
-- Name: sys_management_dict_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_management_dict_id_seq OWNED BY public.sys_management_dict.id;


--
-- Name: sys_management_menu; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_management_menu (
    id bigint NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    menu_name character varying(100) NOT NULL,
    icon character varying(100),
    path character varying(200),
    component character varying(200),
    redirect character varying(200),
    parent_id bigint,
    sort bigint DEFAULT 0,
    hidden boolean DEFAULT false,
    keep_alive boolean DEFAULT false,
    affix boolean DEFAULT false,
    always_show boolean DEFAULT false,
    title text
);


ALTER TABLE public.sys_management_menu OWNER TO postgres;

--
-- Name: COLUMN sys_management_menu.menu_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_menu.menu_name IS '菜单名称';


--
-- Name: COLUMN sys_management_menu.icon; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_menu.icon IS '图标';


--
-- Name: COLUMN sys_management_menu.path; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_menu.path IS '路由路径';


--
-- Name: COLUMN sys_management_menu.component; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_menu.component IS '前端组件';


--
-- Name: COLUMN sys_management_menu.redirect; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_menu.redirect IS '重定向';


--
-- Name: COLUMN sys_management_menu.parent_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_menu.parent_id IS '父菜单ID';


--
-- Name: COLUMN sys_management_menu.sort; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_menu.sort IS '排序';


--
-- Name: COLUMN sys_management_menu.hidden; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_menu.hidden IS '是否隐藏';


--
-- Name: COLUMN sys_management_menu.keep_alive; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_menu.keep_alive IS '缓存';


--
-- Name: COLUMN sys_management_menu.affix; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_menu.affix IS '是否固定';


--
-- Name: COLUMN sys_management_menu.always_show; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_menu.always_show IS '一直显示根路由';


--
-- Name: COLUMN sys_management_menu.title; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_menu.title IS '菜单名';


--
-- Name: sys_management_menu_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_management_menu_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sys_management_menu_id_seq OWNER TO postgres;

--
-- Name: sys_management_menu_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_management_menu_id_seq OWNED BY public.sys_management_menu.id;


--
-- Name: sys_management_permission; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_management_permission (
    id bigint NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    name character varying(100) NOT NULL,
    resource character varying(200) NOT NULL,
    action character varying(20) DEFAULT 'all'::character varying,
    domain_id bigint,
    domain character varying(20) NOT NULL,
    CONSTRAINT chk_sys_management_permission_domain CHECK (((domain)::text = ANY (ARRAY[('menu'::character varying)::text, ('api'::character varying)::text, ('button'::character varying)::text, ('data'::character varying)::text])))
);


ALTER TABLE public.sys_management_permission OWNER TO postgres;

--
-- Name: COLUMN sys_management_permission.name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_permission.name IS '权限名称';


--
-- Name: COLUMN sys_management_permission.resource; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_permission.resource IS '资源标识，如:/api/user';


--
-- Name: COLUMN sys_management_permission.action; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_permission.action IS '操作:all|view|create|update|delete';


--
-- Name: COLUMN sys_management_permission.domain_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_permission.domain_id IS '关联领域表ID(menu/api/button)';


--
-- Name: COLUMN sys_management_permission.domain; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_permission.domain IS '领域,(''menu'',''api'',''button'',''data'')';


--
-- Name: sys_management_permission_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_management_permission_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sys_management_permission_id_seq OWNER TO postgres;

--
-- Name: sys_management_permission_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_management_permission_id_seq OWNED BY public.sys_management_permission.id;


--
-- Name: sys_management_role; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_management_role (
    id bigint NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    role_name character varying(191),
    parent_id bigint,
    permission_hash text
);


ALTER TABLE public.sys_management_role OWNER TO postgres;

--
-- Name: COLUMN sys_management_role.parent_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_role.parent_id IS '父角色ID';


--
-- Name: COLUMN sys_management_role.permission_hash; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_role.permission_hash IS '权限哈希，用于缓存失效判断';


--
-- Name: sys_management_role_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_management_role_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sys_management_role_id_seq OWNER TO postgres;

--
-- Name: sys_management_role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_management_role_id_seq OWNED BY public.sys_management_role.id;


--
-- Name: sys_management_role_permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_management_role_permissions (
    role_id bigint NOT NULL,
    permission_id bigint NOT NULL,
    data_scope character varying(20) DEFAULT 'all'::character varying,
    custom_sql character varying(500)
);


ALTER TABLE public.sys_management_role_permissions OWNER TO postgres;

--
-- Name: COLUMN sys_management_role_permissions.data_scope; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_role_permissions.data_scope IS '数据权限范围:all全部|dept部门|self本人|custom自定义';


--
-- Name: COLUMN sys_management_role_permissions.custom_sql; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_role_permissions.custom_sql IS '自定义数据权限SQL条件';


--
-- Name: sys_management_user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_management_user (
    id bigint NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    username character varying(191),
    password text NOT NULL,
    phone text,
    email text,
    active boolean,
    dept_id bigint
);


ALTER TABLE public.sys_management_user OWNER TO postgres;

--
-- Name: COLUMN sys_management_user.username; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_user.username IS '用户名';


--
-- Name: COLUMN sys_management_user.password; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_user.password IS '密码';


--
-- Name: COLUMN sys_management_user.phone; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_user.phone IS '手机号';


--
-- Name: COLUMN sys_management_user.email; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_user.email IS '邮箱';


--
-- Name: COLUMN sys_management_user.dept_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_management_user.dept_id IS '部门ID';


--
-- Name: sys_management_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_management_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sys_management_user_id_seq OWNER TO postgres;

--
-- Name: sys_management_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_management_user_id_seq OWNED BY public.sys_management_user.id;


--
-- Name: sys_management_user_roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_management_user_roles (
    user_id bigint NOT NULL,
    role_id bigint NOT NULL
);


ALTER TABLE public.sys_management_user_roles OWNER TO postgres;

--
-- Name: sys_monitor_operation_log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_monitor_operation_log (
    id bigint NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    ip text,
    method text,
    path text,
    status bigint,
    user_agent text,
    req_param text,
    resp_data text,
    resp_time bigint,
    user_id bigint,
    user_name text
);


ALTER TABLE public.sys_monitor_operation_log OWNER TO postgres;

--
-- Name: COLUMN sys_monitor_operation_log.ip; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_monitor_operation_log.ip IS '请求ip';


--
-- Name: COLUMN sys_monitor_operation_log.method; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_monitor_operation_log.method IS '请求方法';


--
-- Name: COLUMN sys_monitor_operation_log.path; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_monitor_operation_log.path IS '请求路径';


--
-- Name: COLUMN sys_monitor_operation_log.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_monitor_operation_log.status IS '请求状态';


--
-- Name: COLUMN sys_monitor_operation_log.req_param; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_monitor_operation_log.req_param IS '请求Body';


--
-- Name: COLUMN sys_monitor_operation_log.resp_data; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_monitor_operation_log.resp_data IS '响应数据';


--
-- Name: COLUMN sys_monitor_operation_log.user_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_monitor_operation_log.user_id IS '用户id';


--
-- Name: COLUMN sys_monitor_operation_log.user_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_monitor_operation_log.user_name IS '用户名称';


--
-- Name: sys_monitor_operation_log_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_monitor_operation_log_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sys_monitor_operation_log_id_seq OWNER TO postgres;

--
-- Name: sys_monitor_operation_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_monitor_operation_log_id_seq OWNED BY public.sys_monitor_operation_log.id;


--
-- Name: sys_tool_cache; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_tool_cache (
    id bigint NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    key character varying(255),
    value text,
    expires_at timestamp without time zone NOT NULL,
    username text
);


ALTER TABLE public.sys_tool_cache OWNER TO postgres;

--
-- Name: COLUMN sys_tool_cache.username; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tool_cache.username IS '用户名';


--
-- Name: sys_tool_cache_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_tool_cache_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sys_tool_cache_id_seq OWNER TO postgres;

--
-- Name: sys_tool_cache_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_tool_cache_id_seq OWNED BY public.sys_tool_cache.id;


--
-- Name: sys_tool_cron; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_tool_cron (
    id bigint NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    name text,
    method text NOT NULL,
    expression text NOT NULL,
    strategy character varying(20) DEFAULT 'always'::character varying,
    open boolean,
    "extraParams" json,
    "entryId" bigint,
    comment text
);


ALTER TABLE public.sys_tool_cron OWNER TO postgres;

--
-- Name: COLUMN sys_tool_cron.name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tool_cron.name IS '任务名称';


--
-- Name: COLUMN sys_tool_cron.method; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tool_cron.method IS '任务方法';


--
-- Name: COLUMN sys_tool_cron.expression; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tool_cron.expression IS '表达式';


--
-- Name: COLUMN sys_tool_cron.strategy; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tool_cron.strategy IS '执行策略';


--
-- Name: COLUMN sys_tool_cron.open; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tool_cron.open IS '活跃状态';


--
-- Name: COLUMN sys_tool_cron."extraParams"; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tool_cron."extraParams" IS '额外参数';


--
-- Name: COLUMN sys_tool_cron."entryId"; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tool_cron."entryId" IS 'cron ID';


--
-- Name: COLUMN sys_tool_cron.comment; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tool_cron.comment IS '备注';


--
-- Name: sys_tool_cron_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_tool_cron_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sys_tool_cron_id_seq OWNER TO postgres;

--
-- Name: sys_tool_cron_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_tool_cron_id_seq OWNED BY public.sys_tool_cron.id;


--
-- Name: sys_tool_file; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_tool_file (
    id bigint NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    file_name text,
    full_path text,
    mime text
);


ALTER TABLE public.sys_tool_file OWNER TO postgres;

--
-- Name: COLUMN sys_tool_file.file_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tool_file.file_name IS '文件名';


--
-- Name: COLUMN sys_tool_file.full_path; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tool_file.full_path IS '文件完整路径';


--
-- Name: COLUMN sys_tool_file.mime; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tool_file.mime IS '文件类型';


--
-- Name: sys_tool_file_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_tool_file_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sys_tool_file_id_seq OWNER TO postgres;

--
-- Name: sys_tool_file_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_tool_file_id_seq OWNED BY public.sys_tool_file.id;


--
-- Name: sys_tool_service_token; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_tool_service_token (
    id bigint NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp with time zone,
    name character varying(100) NOT NULL,
    token_hash character varying(255) NOT NULL,
    status boolean DEFAULT true,
    expires_at bigint
);


ALTER TABLE public.sys_tool_service_token OWNER TO postgres;

--
-- Name: COLUMN sys_tool_service_token.name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tool_service_token.name IS '令牌名称/描述';


--
-- Name: COLUMN sys_tool_service_token.token_hash; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tool_service_token.token_hash IS '令牌哈希值';


--
-- Name: COLUMN sys_tool_service_token.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tool_service_token.status IS '是否启用';


--
-- Name: COLUMN sys_tool_service_token.expires_at; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_tool_service_token.expires_at IS '过期时间戳(秒)';


--
-- Name: sys_tool_service_token_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_tool_service_token_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.sys_tool_service_token_id_seq OWNER TO postgres;

--
-- Name: sys_tool_service_token_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_tool_service_token_id_seq OWNED BY public.sys_tool_service_token.id;


--
-- Name: sys_tool_token_permission; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_tool_token_permission (
    token_id bigint NOT NULL,
    permission_id bigint NOT NULL
);


ALTER TABLE public.sys_tool_token_permission OWNER TO postgres;

--
-- Name: sys_tool_token_permissions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_tool_token_permissions (
    token_id bigint NOT NULL,
    permission_id bigint NOT NULL
);


ALTER TABLE public.sys_tool_token_permissions OWNER TO postgres;

--
-- Name: sys_management_api id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_api ALTER COLUMN id SET DEFAULT nextval('public.sys_management_api_id_seq'::regclass);


--
-- Name: sys_management_dept id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_dept ALTER COLUMN id SET DEFAULT nextval('public.sys_management_dept_id_seq'::regclass);


--
-- Name: sys_management_dict id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_dict ALTER COLUMN id SET DEFAULT nextval('public.sys_management_dict_id_seq'::regclass);


--
-- Name: sys_management_dict_detail id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_dict_detail ALTER COLUMN id SET DEFAULT nextval('public.sys_management_dict_detail_id_seq'::regclass);


--
-- Name: sys_management_menu id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_menu ALTER COLUMN id SET DEFAULT nextval('public.sys_management_menu_id_seq'::regclass);


--
-- Name: sys_management_permission id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_permission ALTER COLUMN id SET DEFAULT nextval('public.sys_management_permission_id_seq'::regclass);


--
-- Name: sys_management_role id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_role ALTER COLUMN id SET DEFAULT nextval('public.sys_management_role_id_seq'::regclass);


--
-- Name: sys_management_user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_user ALTER COLUMN id SET DEFAULT nextval('public.sys_management_user_id_seq'::regclass);


--
-- Name: sys_monitor_operation_log id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_monitor_operation_log ALTER COLUMN id SET DEFAULT nextval('public.sys_monitor_operation_log_id_seq'::regclass);


--
-- Name: sys_tool_cache id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_tool_cache ALTER COLUMN id SET DEFAULT nextval('public.sys_tool_cache_id_seq'::regclass);


--
-- Name: sys_tool_cron id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_tool_cron ALTER COLUMN id SET DEFAULT nextval('public.sys_tool_cron_id_seq'::regclass);


--
-- Name: sys_tool_file id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_tool_file ALTER COLUMN id SET DEFAULT nextval('public.sys_tool_file_id_seq'::regclass);


--
-- Name: sys_tool_service_token id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_tool_service_token ALTER COLUMN id SET DEFAULT nextval('public.sys_tool_service_token_id_seq'::regclass);


--
-- Data for Name: sys_management_api; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_management_api (id, created_at, updated_at, deleted_at, path, method, description, group_en, group_cn) FROM stdin;
2	2026-04-18 10:50:40	2026-04-18 10:50:41	\N	/dashboard/recent-operations	GET	获取最近操作记录	dashboard	仪表盘
3	2026-04-18 10:52:10	\N	\N	/dashboard/system-info	GET	获取系统信息	dashboard	仪表盘
216	\N	2026-04-18 15:20:30.850806	\N	/menu/create	POST	添加菜单	menu	菜单
218	\N	2026-04-18 15:20:35.118278	\N	/menu/delete	POST	删除菜单	menu	菜单
220	\N	2026-04-18 15:20:47.008698	\N	/api/list	POST	获取API列表	api	API
221	\N	2026-04-18 15:20:49.114311	\N	/api/create	POST	添加API	api	API
222	\N	2026-04-18 15:20:51.188525	\N	/api/update	POST	编辑API	api	API
223	\N	2026-04-18 15:20:53.68847	\N	/api/delete	POST	删除API	api	API
244	2026-04-18 15:18:48.424482	2026-04-18 15:20:55.955959	\N	/api/elTree	POST	elTree	api	API
247	2026-04-18 16:12:45.484907	2026-04-18 16:12:45.484907	\N	/role_permission/rebuild	POST	重置角色权限	role_permission	角色权限
1	2026-04-18 10:46:01	2026-04-18 10:45:58	\N	/dashboard/statistics	GET	获取统计信息	dashboard	仪表盘
204	\N	2026-04-18 23:35:44.125117	\N	/user/list	POST	获取所有用户	user	用户
249	2026-04-18 23:58:06.921498	2026-04-18 23:59:17.637958	\N	/serviceToken/list	POST	获取服务令牌列表	serviceToken	服务令牌
203	\N	\N	\N	/user/getUserInfo	GET	获取用户信息	user	用户
207	\N	\N	\N	/user/update	POST	编辑用户	user	用户
213	\N	\N	\N	/role/update	POST	编辑角色	role	角色
224	\N	\N	\N	/dict/list	POST	获取字典列表	dict	字典
230	\N	\N	\N	/dict_detail/update	POST	编辑字典详情	dict_detail	字典详情
227	\N	\N	\N	/dict/delete	POST	删除字典	dict	字典
202	\N	\N	\N	/logout	POST	登出	base	基础
208	\N	\N	\N	/user/modifyPasswd	POST	修改用户密码	user	用户
233	\N	\N	\N	/file/upload	POST	上传文件	file	文件
232	\N	\N	\N	/file/list	POST	获取文件列表	file	文件
209	\N	\N	\N	/user/switchActive	POST	切换用户状态	user	用户
200	\N	\N	\N	/captcha	POST	获取验证码	base	基础
234	\N	\N	\N	/file/download	GET	下载文件	file	文件
235	\N	\N	\N	/file/delete	GET	删除文件	file	文件
205	\N	\N	\N	/user/delete	POST	删除用户	user	用户
212	\N	\N	\N	/role/delete	POST	删除角色	role	角色
211	\N	\N	\N	/role/create	POST	添加角色	role	角色
236	\N	\N	\N	/cron/list	POST	获取定时任务	cron	定时任务
225	\N	\N	\N	/dict/create	POST	添加字典	dict	字典
226	\N	\N	\N	/dict/update	POST	编辑字典	dict	字典
237	\N	\N	\N	/cron/create	POST	添加定时任务	cron	定时任务
206	\N	\N	\N	/user/create	POST	添加用户	user	用户
210	\N	\N	\N	/role/list	POST	获取所有角色	role	角色
239	\N	\N	\N	/cron/delete	POST	删除定时任务	cron	定时任务
231	\N	\N	\N	/dict_detail/delete	POST	删除字典详情	dict_detail	字典详情
238	\N	\N	\N	/cron/update	POST	编辑定时任务	cron	定时任务
214	\N	\N	\N	/role/updateRoleMenu	POST	编辑角色菜单	role	角色
215	\N	2026-04-18 15:20:27.928536	\N	/menu/list	GET	获取所有菜单	menu	菜单
217	\N	2026-04-18 15:20:33.136977	\N	/menu/update	POST	编辑菜单	menu	菜单
246	2026-04-18 16:09:49.149787	2026-04-18 16:09:49.149787	\N	/menu/elTree	POST	菜单elTree	menu	菜单
5	2026-04-18 11:06:42.623232	2026-04-18 23:08:50.056164	\N	/dept/list	POST	部门列表	dept	部门
229	\N	\N	\N	/dict_detail/create	POST	添加字典详情	dict_detail	字典详情
201	\N	\N	\N	/login	POST	登录	base	基础
228	\N	\N	\N	/dict_detail/list	POST	获取字典详情	dict_detail	字典详情
240	\N	\N	\N	/opl/list	POST	获取操作日志	opl	操作日志
248	2026-04-18 23:39:44.605729	2026-04-18 23:54:57.496227	\N	/dept/getElTreeDepts	POST	获取部门树	dept	部门
241	\N	2026-04-19 11:19:33.706838	\N	/opl/delete	POST	删除操作日志	opl	操作日志
250	2026-04-19 11:54:43.853749	2026-04-19 11:57:59.236109	\N	/opl/deleteByIds	POST	批量删除操作日志	opl	操作日志
\.


--
-- Data for Name: sys_management_dept; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_management_dept (id, created_at, updated_at, deleted_at, dept_name, parent_id, path, sort, status, level) FROM stdin;
2	2026-04-07 17:00:32.880775	2026-04-07 17:00:32.88671	\N	研发部门	1	0/1	0	t	2
1	2026-04-08 16:25:48	2026-04-08 16:25:51	\N	总公司	0	0	1	t	1
\.


--
-- Data for Name: sys_management_dict; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_management_dict (id, created_at, updated_at, deleted_at, cn_name, en_name) FROM stdin;
\.


--
-- Data for Name: sys_management_dict_detail; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_management_dict_detail (id, created_at, updated_at, deleted_at, label, value, sort, dict_id, parent_id, description) FROM stdin;
\.


--
-- Data for Name: sys_management_menu; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_management_menu (id, created_at, updated_at, deleted_at, menu_name, icon, path, component, redirect, parent_id, sort, hidden, keep_alive, affix, always_show, title) FROM stdin;
6	2026-04-05 21:30:24	2026-04-05 21:31:02	\N	字典管理	\N	/sysManagement/dict	sysManagement/dict/index.vue	\N	1	5	f	f	f	f	字典管理
102	2026-04-05 21:30:47	2026-04-05 21:31:15	\N	cenu1-1	\N	/cenu/cenu1/cenu1-1	cenu/cenu1/cenu1-1/index.vue	\N	101	1	f	f	f	f	cenu1-1
20	2026-04-05 21:30:28	2026-04-05 21:31:06	\N	系统工具	config	/systool	Layout	/systool/cron	0	4	f	f	f	f	系统工具
7	2026-04-05 21:30:26	2026-04-05 21:31:04	\N	部门管理	\N	/sysManagement/dept	sysManagement/dept/index.vue	\N	1	6	f	f	f	f	部门管理
3	2026-04-05 21:30:19	2026-04-05 21:30:55	\N	角色管理	\N	/sysManagement/role	sysManagement/role/index.vue	\N	1	2	f	f	f	f	角色管理
4	2026-04-05 21:30:21	2026-04-05 21:30:58	\N	菜单管理	\N	/sysManagement/menu	sysManagement/menu/index.vue	\N	1	3	f	f	f	f	菜单管理
5	2026-04-05 21:30:22	2026-04-05 21:31:00	\N	接口管理	\N	/sysManagement/api	sysManagement/api/index.vue	\N	1	4	f	f	f	f	接口管理
2	2026-04-05 21:30:12	2026-04-05 21:30:14	\N	用户管理	\N	/sysManagement/user	sysManagement/user/index.vue	\N	1	1	f	f	f	f	用户管理
1	2026-04-05 21:30:06	2026-04-05 21:30:09	\N	系统管理	lock	/sysManagement	Layout	/sysManagement/user	0	1	f	f	f	f	系统管理
100	2026-04-05 21:30:41	2026-04-05 21:31:12	\N	多级菜单	menu	/cenu	Layout	/cenu/cenu1	0	2	f	f	f	f	多级菜单
41	2026-04-05 21:30:39	2026-04-05 21:31:10	\N	操作日志	\N	/sysMonitor/operationLog	sysMonitor/operationLog/index.vue	\N	40	1	f	f	f	f	操作日志
40	2026-04-05 21:30:37	2026-04-05 19:00:59.660052	\N	系统监控	monitor	/sysMonitor	Layout	/sysMonitor/operationLog	0	5	f	f	f	t	系统监控
21	2026-04-05 21:30:31	2026-04-05 21:31:07	\N	定时任务	\N	/systool/cron	sysTool/cron/index.vue	\N	20	1	f	f	f	f	定时任务
101	2026-04-05 21:30:45	2026-04-05 21:31:14	\N	cenu1	\N	/cenu/cenu1	cenu/cenu1/index.vue	/cenu/cenu1/cenu1-1	100	1	f	f	f	f	cenu1
22	2026-04-05 21:30:35	2026-04-05 21:31:09	\N	文件管理	\N	/systool/file	sysTool/file/index.vue	\N	20	2	f	f	f	f	文件管理
103	2026-04-05 21:30:48	2026-04-05 21:31:17	\N	cenu1-2	\N	/cenu/cenu1/cenu1-2	cenu/cenu1/cenu1-2/index.vue	\N	101	2	f	f	f	f	cenu1-2
107	2026-04-15 22:13:16.230994	2026-04-15 22:13:16.230994	\N			/systool/serviceToken	sysTool/service_token/index.vue		20	3	f	f	f	f	服务令牌
\.


--
-- Data for Name: sys_management_permission; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_management_permission (id, created_at, updated_at, deleted_at, name, resource, action, domain_id, domain) FROM stdin;
2	\N	\N	\N	用户管理	/sysManagement/user	view	2	menu
235	\N	\N	\N	删除文件	/file/delete	delete	235	api
227	\N	\N	\N	删除字典	/dict/delete	delete	227	api
103	\N	\N	\N	cenu1-2	/cenu/cenu1/cenu1-2	view	103	menu
21	\N	\N	\N	定时任务	/systool/cron	view	21	menu
211	\N	\N	\N	添加角色	/role/create	create	211	api
212	\N	\N	\N	删除角色	/role/delete	delete	212	api
214	\N	\N	\N	编辑角色菜单	/role/updateRoleMenu	update	214	api
217	\N	\N	\N	编辑菜单	/menu/update	update	217	api
206	\N	\N	\N	添加用户	/user/create	create	206	api
202	\N	\N	\N	登出	/logout	all	202	api
207	\N	\N	\N	编辑用户	/user/update	update	207	api
205	\N	\N	\N	删除用户	/user/delete	delete	205	api
239	\N	\N	\N	删除定时任务	/cron/delete	delete	239	api
6	\N	\N	\N	字典管理	/sysManagement/dict	view	6	menu
231	\N	\N	\N	删除字典详情	/dictDetail/delete	delete	231	api
200	\N	\N	\N	获取验证码	/captcha	all	200	api
7	\N	\N	\N	部门管理	/sysManagement/dept	view	7	menu
216	\N	\N	\N	添加菜单	/menu/create	create	216	api
218	\N	\N	\N	删除菜单	/menu/delete	delete	218	api
221	\N	\N	\N	添加API	/api/create	create	221	api
41	\N	\N	\N	操作日志	/sysMonitor/operationLog	view	41	menu
40	\N	\N	\N	系统监控	/sysMonitor	view	40	menu
233	\N	\N	\N	上传文件	/file/upload	create	233	api
4	\N	\N	\N	菜单管理	/sysManagement/menu	view	4	menu
237	\N	\N	\N	添加定时任务	/cron/create	create	237	api
225	\N	\N	\N	添加字典	/dict/create	create	225	api
213	\N	\N	\N	编辑角色	/role/update	update	213	api
20	\N	\N	\N	系统工具	/systool	view	20	menu
238	\N	\N	\N	编辑定时任务	/cron/update	update	238	api
209	\N	\N	\N	切换用户状态	/user/switchActive	update	209	api
208	\N	\N	\N	修改用户密码	/user/modifyPasswd	update	208	api
5	\N	\N	\N	接口管理	/sysManagement/api	view	5	menu
3	\N	\N	\N	角色管理	/sysManagement/role	view	3	menu
1	\N	\N	\N	系统管理	/sysManagement	view	1	menu
226	\N	\N	\N	编辑字典	/dict/update	update	226	api
102	\N	\N	\N	cenu1-1	/cenu/cenu1/cenu1-1	view	102	menu
229	\N	\N	\N	添加字典详情	/dictDetail/create	create	229	api
22	\N	\N	\N	文件管理	/systool/file	view	22	menu
100	\N	\N	\N	多级菜单	/cenu	view	100	menu
101	\N	\N	\N	cenu1	/cenu/cenu1	view	101	menu
230	\N	\N	\N	编辑字典详情	/dictDetail/update	update	230	api
242	2026-04-12 19:05:03.062708	2026-04-12 19:05:03.062708	\N	CMDB	/cmdb	view	106	menu
243	2026-04-15 22:13:16.245458	2026-04-15 22:13:16.245458	\N	服务令牌	/systool/serviceToken	view	107	menu
210	2026-04-16 23:57:47	2026-04-16 23:57:50	\N	获取所有角色	/role/list	create	210	api
201	\N	\N	\N	登录	/login	create	201	api
203	\N	\N	\N	获取用户信息	/user/getUserInfo	read	203	api
215	\N	\N	\N	获取所有菜单	/menu/list	read	215	api
244	2026-04-18 10:46:58	2026-04-18 10:47:00	\N	获取仪表盘数据	/dashboard/statistics	read	1	api
246	2026-04-18 10:54:35	\N	\N	获取系统信息	/dashboard/system-info	read	3	api
245	2026-04-18 10:53:13	\N	\N	获取最近操作记录	/dashboard/recent-operations	read	2	api
220	\N	\N	\N	获取所有API	/api/list	create	220	api
251	2026-04-18 23:39:44.614302	2026-04-18 23:54:57.508754	\N	获取部门树	/dept/getElTreeDepts	create	248	api
232	\N	\N	\N	获取文件列表	/file/list	create	232	api
240	\N	\N	\N	获取操作日志	/opl/list	create	240	api
234	\N	\N	\N	下载文件	/file/download	create	234	api
236	\N	\N	\N	获取定时任务	/cron/list	create	236	api
228	\N	\N	\N	获取字典详情	/dictDetail/list	create	228	api
224	\N	\N	\N	获取字典列表	/dict/list	create	224	api
222	\N	\N	\N	编辑API	/api/update	create	222	api
223	\N	\N	\N	删除API	/api/delete	create	223	api
247	2026-04-18 15:18:48.432194	2026-04-18 15:18:48.432194	\N	elTree	/api/elTree	create	244	api
248	2026-04-18 16:09:49.153647	2026-04-18 16:09:49.153647	\N	菜单elTree	/menu/elTree	create	246	api
249	2026-04-18 16:12:45.491423	2026-04-18 16:12:45.491423	\N	重置角色权限	/role_permission/rebuild	create	247	api
250	2026-04-18 23:08:50.06371	2026-04-18 23:08:50.06371	\N	部门列表	/dept/list	create	5	api
204	0001-01-01 00:00:00	2026-04-18 23:35:44.132263	\N	获取所有用户	/user/list	create	204	api
252	2026-04-18 23:58:06.929827	2026-04-18 23:59:17.644424	\N	获取服务令牌列表	/serviceToken/list	create	249	api
241	0001-01-01 00:00:00	2026-04-19 11:19:33.713181	\N	删除操作日志	/opl/delete	create	241	api
253	2026-04-19 11:54:43.85968	2026-04-19 11:57:59.241929	\N	批量删除操作日志	/opl/deleteByIds	create	250	api
\.


--
-- Data for Name: sys_management_role; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_management_role (id, created_at, updated_at, deleted_at, role_name, parent_id, permission_hash) FROM stdin;
1	\N	\N	\N	root	\N	\N
2	\N	\N	\N	pdd	1	
3	\N	\N	\N	zl	1	
4	2026-04-06 17:36:46.469549	2026-04-06 19:05:31.794195	\N	pddzl	2	
\.


--
-- Data for Name: sys_management_role_permissions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_management_role_permissions (role_id, permission_id, data_scope, custom_sql) FROM stdin;
1	235	all	
1	227	all	
1	211	all	
1	212	all	
1	214	all	
1	217	all	
1	206	all	
1	202	all	
1	207	all	
1	205	all	
1	239	all	
3	103	all	
3	41	all	
3	40	all	
3	102	all	
3	100	all	
3	101	all	
1	231	all	
1	200	all	
1	216	all	
1	218	all	
1	221	all	
1	233	all	
1	237	all	
2	2	all	
2	103	all	
2	21	all	
2	6	all	
2	7	all	
2	41	all	
2	40	all	
2	4	all	
2	20	all	
2	5	all	
2	3	all	
2	1	all	
2	102	all	
2	22	all	
2	100	all	
2	101	all	
1	225	all	
1	213	all	
1	238	all	
1	209	all	
1	208	all	
1	226	all	
1	229	all	
1	230	all	
1	210	all	
1	201	all	
1	203	all	
1	215	all	
1	244	all	
1	246	all	
1	245	all	
1	220	all	
1	251	all	
1	232	all	
1	240	all	
1	234	all	
1	236	all	
1	228	all	
1	224	all	
1	222	all	
1	223	all	
1	247	all	
1	248	all	
1	249	all	
1	250	all	
1	204	all	
1	252	all	
1	2	all	
1	21	all	
1	6	all	
1	7	all	
1	41	all	
1	40	all	
1	4	all	
1	20	all	
1	5	all	
1	3	all	
1	1	all	
1	22	all	
1	243	all	
1	219	all	
1	253	all	
1	241	all	
\.


--
-- Data for Name: sys_management_user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_management_user (id, created_at, updated_at, deleted_at, username, password, phone, email, active, dept_id) FROM stdin;
2	2026-04-06 18:54:30.102229	2026-04-06 18:54:30.109346	\N	edit	2c8f2f73f2629177eaf31c7e6ce07950			t	0
1	\N	2026-04-06 19:23:50.784009	\N	admin	e10adc3949ba59abbe56e057f20f883e			t	1
3	2026-04-07 09:58:05.31887	2026-04-07 09:58:05.327061	\N	x1	5a34c35f198949a76e545c540cae83d5			t	0
\.


--
-- Data for Name: sys_management_user_roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_management_user_roles (user_id, role_id) FROM stdin;
2	1
2	2
1	1
1	4
3	4
\.


--
-- Data for Name: sys_monitor_operation_log; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_monitor_operation_log (id, created_at, updated_at, deleted_at, ip, method, path, status, user_agent, req_param, resp_data, resp_time, user_id, user_name) FROM stdin;
\.


--
-- Data for Name: sys_tool_cache; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_tool_cache (id, created_at, updated_at, deleted_at, key, value, expires_at, username) FROM stdin;
1	2026-04-04 23:21:46.225932	2026-04-04 23:21:46.225932	\N	user_single_token:admin	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In1dLCJidWZmZXJUaW1lIjo0MzIwMCwiaXNzIjoicGRkemwiLCJleHAiOjE3NzU0MDI1MDYsIm5iZiI6MTc3NTMxNjEwNn0.fmjwpoeX8JSsHkcv7C491_kshKnvMQFaAmxDyevFslw	2026-04-05 23:21:46.224412	\N
16	2026-04-07 08:51:40.084958	2026-04-07 08:51:40.084958	\N	user_tokens:admin:735efe6dd2be4806	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc1NjA5NTAwLCJuYmYiOjE3NzU0NzQ2NTZ9.CWOPQuWDAfg75_k5NtsPdCFE25MCKoyJ0ow00mkJNjo	2026-04-08 08:51:40.083669	admin
18	2026-04-07 20:52:13.33599	2026-04-07 20:52:13.33599	\N	user_tokens:admin:34219ecb8ddde38a	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc1NjUyNzMzLCJuYmYiOjE3NzU1MjMxMTh9.99nOiZJmtviVm0U8S98Iuv-QW_fWXjknGeNNEWCbAT8	2026-04-08 20:52:13.335264	admin
20	2026-04-08 09:22:01.64148	2026-04-08 09:22:01.64148	\N	user_tokens:admin:28191d875f945d40	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc1Njk3NzIxLCJuYmYiOjE3NzU1NjYzNDR9.l4-IXuFMnslGtI_V-ydAT1J12O7M_TMF_PTAbIk9JtM	2026-04-09 09:22:01.640719	admin
21	2026-04-08 11:33:46.512167	2026-04-08 11:33:46.512167	\N	user_tokens:admin:ea66c04f963934ad	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc1NzA1NjI2LCJuYmYiOjE3NzU2MTkyMjZ9.mI18Uh1a2pFhOTJB-FyXxLFuNqQ2BGGevvI40Z4bCkc	2026-04-09 11:33:46.510679	admin
22	2026-04-08 23:49:56.478934	2026-04-08 23:49:56.478934	\N	user_tokens:admin:db5b67f16675a796	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc1NzQ5Nzk2LCJuYmYiOjE3NzU2NjMzOTZ9.-sBwhgDdU5yvkMm_7qK6b8N4p2lKTOskxFa5EoiKjn0	2026-04-09 23:49:56.477165	admin
23	2026-04-08 23:52:14.072782	2026-04-08 23:52:14.072782	\N	user_tokens:admin:ae076ff8886768d6	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc1NzQ5OTM0LCJuYmYiOjE3NzU2NjM1MzR9.jsUyDf8vaQ4Q4q4si8p1TGuSL7vmwLtINynokMMs2iI	2026-04-09 23:52:14.071875	admin
24	2026-04-09 00:40:22.685907	2026-04-09 00:40:22.685907	\N	user_tokens:admin:c297bcceb96e0782	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc1NzUyODIyLCJuYmYiOjE3NzU2NjY0MjJ9.fPjh6eSoW-OcWGA-8fvGtHoKydCtyGaRZ8Njx6gsluI	2026-04-10 00:40:22.684756	admin
25	2026-04-09 19:47:32.099204	2026-04-09 19:47:32.099204	\N	user_tokens:admin:0f8c7f2a799125bb	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc1ODIxNjUyLCJuYmYiOjE3NzU3MzUyNTJ9.fQw4Id2aSirL14D3x_c8CCG-tpfmNf27KCUxFpxWLgo	2026-04-10 19:47:32.098017	admin
26	2026-04-09 22:54:39.693889	2026-04-09 22:54:39.693889	\N	user_tokens:admin:430144c997d1bd3f	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc1ODMyODc5LCJuYmYiOjE3NzU3NDY0Nzl9.n0WmBl5fJYfMU_sI6UsjS1gDGJpRVBWXY350mF2gzu0	2026-04-10 22:54:39.692133	admin
27	2026-04-09 23:58:25.356011	2026-04-09 23:58:25.356011	\N	user_tokens:admin:c10e1422bd37cbef	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc1ODM2NzA1LCJuYmYiOjE3NzU3NTAzMDV9.GPtQs0AuSAZMI-_yP0NkpvLVPRcDFdfKhAMspI7puks	2026-04-10 23:58:25.354941	admin
28	2026-04-11 13:06:17.804925	2026-04-11 13:06:17.804925	\N	user_tokens:admin:45fa61e37dce911e	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc1OTcwMzc3LCJuYmYiOjE3NzU4ODM5Nzd9.gyg7I3ZI5_7yVDnON7jskwwYbePh-iD4REG3J_9HNnQ	2026-04-12 13:06:17.80261	admin
30	2026-04-12 10:52:20.367393	2026-04-12 10:52:20.367393	\N	user_tokens:admin:60341629db62a8fc	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc2MDQ4NzQwLCJuYmYiOjE3NzU4OTIxMDZ9.P5gZZNbq_ixuDsF1xyOjV8EXHd0slXdKUZaR4UMaILk	2026-04-13 10:52:20.366141	admin
31	2026-04-12 15:47:22.62121	2026-04-12 15:47:22.62121	\N	user_tokens:admin:9fd0ec0254ecf092	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc2MDY2NDQyLCJuYmYiOjE3NzU5ODAwNDJ9.eiI5Ru6WYimbbI57ybAiJTzV3JXRzq3EBpSjBgneHYs	2026-04-13 15:47:22.619928	admin
32	2026-04-15 22:08:14.787528	2026-04-15 22:08:14.787529	\N	user_tokens:admin:b094fc60a5b271ad	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc2MzQ4NDk0LCJuYmYiOjE3NzYyNjIwOTR9.95tcB1GwWv-GJmwYlxG_MgAa55B-ejvJ2iZd8htx368	2026-04-16 22:08:14.786213	admin
46	2026-04-17 18:34:29.609772	2026-04-17 18:34:29.609773	\N	user_tokens:admin:d019594bbac429ce	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc2NTA4NDY5LCJuYmYiOjE3NzY0MjIwNjl9.h0c1wt_yT6jmQIoL7yaBlG-nQHK3EHKaTnoxaGlQVlQ	2026-04-18 18:34:29.608809	admin
47	2026-04-18 10:38:50.478918	2026-04-18 10:38:50.478918	\N	user_tokens:admin:76532cb3996d5d1b	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc2NTY2MzMwLCJuYmYiOjE3NzY0Nzk5MzB9.7r2yuoNGIQLW-4RvZcnrTwpTNhVpIg72JlqtbhHDBRo	2026-04-19 10:38:50.477781	admin
48	2026-04-18 10:39:03.254063	2026-04-18 10:39:03.254064	\N	user_tokens:admin:a07d6fc73e04313f	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc2NTY2MzQzLCJuYmYiOjE3NzY0Nzk5NDN9.MJJBOUdjWGe4cujvNpWmsoX0J-G7GbYLDKnT8EPX-i8	2026-04-19 10:39:03.253299	admin
49	2026-04-18 10:41:05.528052	2026-04-18 10:41:05.528052	\N	user_tokens:admin:0e052f92fbbeaf97	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc2NTY2NDY1LCJuYmYiOjE3NzY0ODAwNjV9.J_cTsglP8GfoOxekTu9eQxI5WlWpSgGGr-mf55bOseM	2026-04-19 10:41:05.526961	admin
51	2026-04-18 23:01:27.029735	2026-04-18 23:01:27.029735	\N	user_tokens:admin:b46d17d3b43ac3c2	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc2NjEwODg3LCJuYmYiOjE3NzY0ODAyMTV9.1RyBZ4-Qt6Lp0UIxRDIuP6NMmzJzRx45mEF5UwAMN9g	2026-04-19 23:01:27.029056	admin
53	2026-04-19 11:14:29.181914	2026-04-19 11:14:29.181915	\N	user_tokens:admin:cf49d360169e16f0	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc2NjU0ODY5LCJuYmYiOjE3NzY1MjQ1Mzh9.a9ewq5cH8fUfxuQfNYpcFOTIY6EpRT9XuUu669UjlUc	2026-04-20 11:14:29.174871	admin
54	2026-04-19 11:15:27.966872	2026-04-19 11:15:27.966872	\N	user_tokens:admin:d814af3ef8656d24	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MSwidXNlcm5hbWUiOiJhZG1pbiIsInJvbGVzIjpbeyJpZCI6MSwicm9sZU5hbWUiOiJyb290In0seyJpZCI6NCwicm9sZU5hbWUiOiJwZGR6bCJ9XSwiYnVmZmVyVGltZSI6NDMyMDAsImlzcyI6InBkZHpsIiwiZXhwIjoxNzc2NjU0OTI3LCJuYmYiOjE3NzY1Njg1Mjd9.AILoq1Oqde2wx9mqQ6mSaUHJAgKe_DuHk1mkGwa5kps	2026-04-20 11:15:27.966459	admin
\.


--
-- Data for Name: sys_tool_cron; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_tool_cron (id, created_at, updated_at, deleted_at, name, method, expression, strategy, open, "extraParams", "entryId", comment) FROM stdin;
\.


--
-- Data for Name: sys_tool_file; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_tool_file (id, created_at, updated_at, deleted_at, file_name, full_path, mime) FROM stdin;
\.


--
-- Data for Name: sys_tool_service_token; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_tool_service_token (id, created_at, updated_at, deleted_at, name, token_hash, status, expires_at) FROM stdin;
3	2026-04-17 12:19:24.337707	2026-04-17 17:25:51.571407	\N	test	75749e4f6f96b1e230e5c94a67d68f32ccaf98be50046ea1e593ec25092cb463	t	\N
\.


--
-- Data for Name: sys_tool_token_permission; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_tool_token_permission (token_id, permission_id) FROM stdin;
3	210
\.


--
-- Data for Name: sys_tool_token_permissions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_tool_token_permissions (token_id, permission_id) FROM stdin;
\.


--
-- Name: sys_management_api_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_management_api_id_seq', 250, true);


--
-- Name: sys_management_dept_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_management_dept_id_seq', 3, true);


--
-- Name: sys_management_dict_detail_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_management_dict_detail_id_seq', 1, false);


--
-- Name: sys_management_dict_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_management_dict_id_seq', 1, false);


--
-- Name: sys_management_menu_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_management_menu_id_seq', 107, true);


--
-- Name: sys_management_permission_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_management_permission_id_seq', 253, true);


--
-- Name: sys_management_role_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_management_role_id_seq', 4, true);


--
-- Name: sys_management_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_management_user_id_seq', 3, true);


--
-- Name: sys_monitor_operation_log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_monitor_operation_log_id_seq', 2279, true);


--
-- Name: sys_tool_cache_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_tool_cache_id_seq', 54, true);


--
-- Name: sys_tool_cron_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_tool_cron_id_seq', 1, false);


--
-- Name: sys_tool_file_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_tool_file_id_seq', 1, false);


--
-- Name: sys_tool_service_token_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_tool_service_token_id_seq', 3, true);


--
-- Name: sys_management_api sys_management_api_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_api
    ADD CONSTRAINT sys_management_api_pkey PRIMARY KEY (id);


--
-- Name: sys_management_dept sys_management_dept_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_dept
    ADD CONSTRAINT sys_management_dept_pkey PRIMARY KEY (id);


--
-- Name: sys_management_dict_detail sys_management_dict_detail_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_dict_detail
    ADD CONSTRAINT sys_management_dict_detail_pkey PRIMARY KEY (id);


--
-- Name: sys_management_dict sys_management_dict_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_dict
    ADD CONSTRAINT sys_management_dict_pkey PRIMARY KEY (id);


--
-- Name: sys_management_menu sys_management_menu_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_menu
    ADD CONSTRAINT sys_management_menu_pkey PRIMARY KEY (id);


--
-- Name: sys_management_permission sys_management_permission_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_permission
    ADD CONSTRAINT sys_management_permission_pkey PRIMARY KEY (id);


--
-- Name: sys_management_role_permissions sys_management_role_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_role_permissions
    ADD CONSTRAINT sys_management_role_permissions_pkey PRIMARY KEY (role_id, permission_id);


--
-- Name: sys_management_role sys_management_role_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_role
    ADD CONSTRAINT sys_management_role_pkey PRIMARY KEY (id);


--
-- Name: sys_management_user sys_management_user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_user
    ADD CONSTRAINT sys_management_user_pkey PRIMARY KEY (id);


--
-- Name: sys_management_user_roles sys_management_user_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_user_roles
    ADD CONSTRAINT sys_management_user_roles_pkey PRIMARY KEY (user_id, role_id);


--
-- Name: sys_monitor_operation_log sys_monitor_operation_log_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_monitor_operation_log
    ADD CONSTRAINT sys_monitor_operation_log_pkey PRIMARY KEY (id);


--
-- Name: sys_tool_cache sys_tool_cache_key_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_tool_cache
    ADD CONSTRAINT sys_tool_cache_key_key UNIQUE (key);


--
-- Name: sys_tool_cache sys_tool_cache_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_tool_cache
    ADD CONSTRAINT sys_tool_cache_pkey PRIMARY KEY (id);


--
-- Name: sys_tool_cron sys_tool_cron_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_tool_cron
    ADD CONSTRAINT sys_tool_cron_pkey PRIMARY KEY (id);


--
-- Name: sys_tool_file sys_tool_file_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_tool_file
    ADD CONSTRAINT sys_tool_file_pkey PRIMARY KEY (id);


--
-- Name: sys_tool_service_token sys_tool_service_token_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_tool_service_token
    ADD CONSTRAINT sys_tool_service_token_pkey PRIMARY KEY (id);


--
-- Name: sys_tool_token_permission sys_tool_token_permission_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_tool_token_permission
    ADD CONSTRAINT sys_tool_token_permission_pkey PRIMARY KEY (token_id, permission_id);


--
-- Name: sys_tool_token_permissions sys_tool_token_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_tool_token_permissions
    ADD CONSTRAINT sys_tool_token_permissions_pkey PRIMARY KEY (token_id, permission_id);


--
-- Name: sys_management_dept uni_sys_management_dept_dept_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_dept
    ADD CONSTRAINT uni_sys_management_dept_dept_name UNIQUE (dept_name);


--
-- Name: sys_management_dict uni_sys_management_dict_cn_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_dict
    ADD CONSTRAINT uni_sys_management_dict_cn_name UNIQUE (cn_name);


--
-- Name: sys_management_dict uni_sys_management_dict_en_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_dict
    ADD CONSTRAINT uni_sys_management_dict_en_name UNIQUE (en_name);


--
-- Name: sys_management_menu uni_sys_management_menu_menu_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_menu
    ADD CONSTRAINT uni_sys_management_menu_menu_name UNIQUE (menu_name);


--
-- Name: sys_management_menu uni_sys_management_menu_path; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_menu
    ADD CONSTRAINT uni_sys_management_menu_path UNIQUE (path);


--
-- Name: sys_management_menu uni_sys_management_menu_title; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_menu
    ADD CONSTRAINT uni_sys_management_menu_title UNIQUE (title);


--
-- Name: sys_management_permission uni_sys_management_permission_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_permission
    ADD CONSTRAINT uni_sys_management_permission_name UNIQUE (name);


--
-- Name: sys_management_permission uni_sys_management_permission_resource; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_permission
    ADD CONSTRAINT uni_sys_management_permission_resource UNIQUE (resource);


--
-- Name: sys_management_role uni_sys_management_role_role_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_role
    ADD CONSTRAINT uni_sys_management_role_role_name UNIQUE (role_name);


--
-- Name: sys_management_user uni_sys_management_user_username; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_user
    ADD CONSTRAINT uni_sys_management_user_username UNIQUE (username);


--
-- Name: sys_tool_cron uni_sys_tool_cron_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_tool_cron
    ADD CONSTRAINT uni_sys_tool_cron_name UNIQUE (name);


--
-- Name: sys_tool_service_token uni_sys_tool_service_token_name; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_tool_service_token
    ADD CONSTRAINT uni_sys_tool_service_token_name UNIQUE (name);


--
-- Name: sys_tool_service_token uni_sys_tool_service_token_token_hash; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_tool_service_token
    ADD CONSTRAINT uni_sys_tool_service_token_token_hash UNIQUE (token_hash);


--
-- Name: sys_management_api unique_path_method; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_management_api
    ADD CONSTRAINT unique_path_method UNIQUE (path, method);


--
-- Name: idx_sys_management_api_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_management_api_deleted_at ON public.sys_management_api USING btree (deleted_at);


--
-- Name: idx_sys_management_dept_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_management_dept_deleted_at ON public.sys_management_dept USING btree (deleted_at);


--
-- Name: idx_sys_management_dept_parent_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_management_dept_parent_id ON public.sys_management_dept USING btree (parent_id);


--
-- Name: idx_sys_management_dept_path; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_management_dept_path ON public.sys_management_dept USING btree (path);


--
-- Name: idx_sys_management_dict_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_management_dict_deleted_at ON public.sys_management_dict USING btree (deleted_at);


--
-- Name: idx_sys_management_dict_detail_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_management_dict_detail_deleted_at ON public.sys_management_dict_detail USING btree (deleted_at);


--
-- Name: idx_sys_management_menu_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_management_menu_deleted_at ON public.sys_management_menu USING btree (deleted_at);


--
-- Name: idx_sys_management_menu_parent_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_management_menu_parent_id ON public.sys_management_menu USING btree (parent_id);


--
-- Name: idx_sys_management_permission_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_management_permission_deleted_at ON public.sys_management_permission USING btree (deleted_at);


--
-- Name: idx_sys_management_permission_domain_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_management_permission_domain_id ON public.sys_management_permission USING btree (domain_id);


--
-- Name: idx_sys_management_permission_resource; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_management_permission_resource ON public.sys_management_permission USING btree (resource);


--
-- Name: idx_sys_management_role_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_management_role_deleted_at ON public.sys_management_role USING btree (deleted_at);


--
-- Name: idx_sys_management_role_parent_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_management_role_parent_id ON public.sys_management_role USING btree (parent_id);


--
-- Name: idx_sys_management_user_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_management_user_deleted_at ON public.sys_management_user USING btree (deleted_at);


--
-- Name: idx_sys_monitor_operation_log_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_monitor_operation_log_deleted_at ON public.sys_monitor_operation_log USING btree (deleted_at);


--
-- Name: idx_sys_tool_cache_expires_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_tool_cache_expires_at ON public.sys_tool_cache USING btree (expires_at);


--
-- Name: idx_sys_tool_cron_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_tool_cron_deleted_at ON public.sys_tool_cron USING btree (deleted_at);


--
-- Name: idx_sys_tool_file_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_tool_file_deleted_at ON public.sys_tool_file USING btree (deleted_at);


--
-- Name: idx_sys_tool_service_token_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_tool_service_token_deleted_at ON public.sys_tool_service_token USING btree (deleted_at);


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

\unrestrict 4WWg2v0b89ELCZoC7rixwJMnlXF8OjechMDjAb7DVqne6yGHnhAy8NuEabt9kPy

