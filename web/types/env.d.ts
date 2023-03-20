/// <reference types="vite/client" />

/** 声明 vite 环境变量的类型（如果未声明则默认是 any） */
declare interface ImportMetaEnv {
  readonly VITE_BASE_API: string
  readonly VITE_CLI_PORT: number
  readonly VITE_BASE_PATH: string
  readonly VITE_SERVER_PORT: number
  readonly VITE_ROUTER_HISTORY: "hash" | "html5"
  readonly VITE_PUBLIC_PATH: string
}
