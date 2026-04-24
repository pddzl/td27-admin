package sysTool

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"server/internal/global"
	modelSysManagement "server/internal/model/sysManagement"
	modelSysTool "server/internal/model/sysTool"
	"server/internal/testutil"
)

func TestServiceTokenService_AuthenticateToken(t *testing.T) {
	db := testutil.NewTestDB(t)
	global.TD27_DB = db
	db.AutoMigrate(&modelSysTool.ServiceToken{}, &modelSysTool.TokenPermission{}, &modelSysManagement.PermissionModel{})

	svc := NewServiceTokenService()

	// Create a token
	resp, err := svc.Create(&modelSysTool.CreateServiceTokenReq{
		Name: "test-token",
	})
	require.NoError(t, err)
	require.NotEmpty(t, resp.Token)

	// Authenticate with plain token
	id, err := svc.AuthenticateToken(resp.Token)
	assert.NoError(t, err)
	assert.Equal(t, resp.ID, id)

	// Wrong token
	_, err = svc.AuthenticateToken("sk-wrong")
	assert.Error(t, err)

	// Expired token
	oldTime := time.Now().Add(-1 * time.Hour).Unix()
	expired, _ := svc.Create(&modelSysTool.CreateServiceTokenReq{
		Name:      "expired",
		ExpiresAt: &oldTime,
	})
	_, err = svc.AuthenticateToken(expired.Token)
	assert.Error(t, err)
}

func TestServiceTokenService_ValidateToken(t *testing.T) {
	db := testutil.NewTestDB(t)
	global.TD27_DB = db
	db.AutoMigrate(
		&modelSysTool.ServiceToken{},
		&modelSysTool.TokenPermission{},
		&modelSysManagement.PermissionModel{},
		&modelSysManagement.RolePermissionModel{},
	)
	global.TD27_CONFIG.Casbin.CacheTTL = 3600

	svc := NewServiceTokenService()

	// Create permission
	perm := &modelSysManagement.PermissionModel{
		Name:     "test-api",
		Domain:   modelSysManagement.PermissionDomainAPI,
		Resource: "/test",
		Action:   modelSysManagement.ActionRead,
		DomainID: 1,
	}
	db.Create(perm)

	// Create token with permission
	resp, err := svc.Create(&modelSysTool.CreateServiceTokenReq{
		Name:   "valid-token",
		ApiIDs: []uint{1},
	})
	require.NoError(t, err)

	// Validate correct path/method
	ok, err := svc.ValidateToken(resp.ID, "/test", "GET")
	assert.NoError(t, err)
	assert.True(t, ok)

	// Validate wrong path
	ok, err = svc.ValidateToken(resp.ID, "/wrong", "GET")
	assert.NoError(t, err)
	assert.False(t, ok)
}
