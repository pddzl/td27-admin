package sysManagement

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"server/internal/global"
	"server/internal/model/sysManagement"
	modelSysTool "server/internal/model/sysTool"
	"server/internal/testutil"
)

func setupCasbinTest(t *testing.T) {
	db := testutil.NewTestDB(t)
	global.TD27_DB = db
	db.AutoMigrate(&sysManagement.RolePermissionModel{}, &sysManagement.PermissionModel{}, &modelSysTool.TokenPermission{})
}

func TestCasbinService_EnforceSubject(t *testing.T) {
	setupCasbinTest(t)
	cs := NewCasbinService()
	e := cs.Casbin()
	require.NotNil(t, e)

	// Add a policy for token:99
	_, err := e.AddPolicy("token:99", "/api/test", "read")
	require.NoError(t, err)

	ok, err := cs.EnforceSubject("token:99", "/api/test", "GET")
	assert.NoError(t, err)
	assert.True(t, ok)

	// Wrong method
	ok, err = cs.EnforceSubject("token:99", "/api/test", "POST")
	assert.NoError(t, err)
	assert.False(t, ok)

	// Non-existent subject
	ok, err = cs.EnforceSubject("token:999", "/api/test", "GET")
	assert.NoError(t, err)
	assert.False(t, ok)
}

func TestCasbinService_RebuildSubjectPolicies(t *testing.T) {
	setupCasbinTest(t)
	cs := NewCasbinService()
	e := cs.Casbin()
	require.NotNil(t, e)

	err := cs.RebuildSubjectPolicies("token:42", [][]string{
		{"token:42", "/a", "read"},
		{"token:42", "/b", "create"},
	})
	assert.NoError(t, err)

	policies, _ := e.GetFilteredPolicy(0, "token:42")
	assert.Len(t, policies, 2)

	// Rebuild empty should clear all
	err = cs.RebuildSubjectPolicies("token:42", [][]string{})
	assert.NoError(t, err)

	policies, _ = e.GetFilteredPolicy(0, "token:42")
	assert.Len(t, policies, 0)
}

func TestCasbinService_UpdateResourcePolicies(t *testing.T) {
	setupCasbinTest(t)
	cs := NewCasbinService()
	e := cs.Casbin()
	require.NotNil(t, e)

	// Seed policies with old resource
	_, _ = e.AddPolicy("role1", "/old/path", "create")
	_, _ = e.AddPolicy("token:1", "/old/path", "create")

	err := cs.UpdateResourcePolicies("/old/path", "create", "/new/path", "read")
	assert.NoError(t, err)

	oldPolicies, _ := e.GetFilteredPolicy(1, "/old/path", "create")
	assert.Len(t, oldPolicies, 0)

	newPolicies, _ := e.GetFilteredPolicy(1, "/new/path", "read")
	assert.Len(t, newPolicies, 2)
}
