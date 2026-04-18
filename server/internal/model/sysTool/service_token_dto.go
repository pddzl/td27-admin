package sysTool

import (
	"server/internal/model/common"
)

// CreateServiceTokenReq 创建服务令牌请求
type CreateServiceTokenReq struct {
	Name      string   `json:"name" binding:"required,max=100"`
	ExpiresAt *int64   `json:"expiresAt"` // 过期时间戳(秒)，nil表示永不过期
	ApiIDs    []uint   `json:"apiIds"`    // 关联的API权限ID列表
}

// UpdateServiceTokenReq 更新服务令牌请求
type UpdateServiceTokenReq struct {
	ID        uint     `json:"id" binding:"required"`
	Name      string   `json:"name" binding:"required,max=100"`
	Status    bool     `json:"status"`
	ExpiresAt *int64   `json:"expiresAt"`
	ApiIDs    []uint   `json:"apiIds"`
}

// ServiceTokenResp 服务令牌响应（不包含敏感信息）
type ServiceTokenResp struct {
	ID        uint         `json:"id"`
	Name      string       `json:"name"`
	Status    bool         `json:"status"`
	ExpiresAt *int64       `json:"expiresAt"`
	ApiCount  int          `json:"apiCount"`  // 关联的API数量
	CreatedAt int64        `json:"createdAt"`
}

// ServiceTokenDetailResp 服务令牌详情响应
type ServiceTokenDetailResp struct {
	ServiceTokenResp
	ApiIDs []uint `json:"apiIds"`
}

// CreateServiceTokenResp 创建服务令牌响应（仅创建时返回完整token）
type CreateServiceTokenResp struct {
	ServiceTokenResp
	Token string `json:"token"` // 完整令牌（仅创建时返回一次）
}

// ListServiceTokenReq 列表查询请求
type ListServiceTokenReq struct {
	common.PageInfo
	Name   string `json:"name" form:"name"`
	Status *bool  `json:"status" form:"status"`
}

// ServiceTokenListResp 列表响应
type ServiceTokenListResp struct {
	List  []ServiceTokenResp `json:"list"`
	Total int64              `json:"total"`
}

// ValidateTokenReq 验证令牌请求
type ValidateTokenReq struct {
	Token string `json:"token" binding:"required"`
}
