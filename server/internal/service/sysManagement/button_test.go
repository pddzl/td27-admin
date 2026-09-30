package sysManagement

import (
	"io"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"server/internal/global"
	modelSysManagement "server/internal/model/sysManagement"
	"server/internal/testutil"
)

func setupButtonTest(t *testing.T) {
	t.Helper()
	db := testutil.NewTestDB(t)
	global.TD27_DB = db
	global.TD27_LOG = slog.New(slog.NewTextHandler(io.Discard, nil))
	require.NoError(t, db.AutoMigrate(
		&modelSysManagement.ButtonModel{},
		&modelSysManagement.PermissionModel{},
		&modelSysManagement.RolePermissionModel{},
	))
}

// seedButtonPermission creates a button-domain permission and grants it to a role.
func seedButtonPermission(t *testing.T, resource string, roleID uint) uint {
	t.Helper()
	perm := &modelSysManagement.PermissionModel{
		Name:     "perm-" + resource,
		Domain:   modelSysManagement.PermissionDomainButton,
		Resource: resource,
		Action:   modelSysManagement.ActionExecute,
	}
	require.NoError(t, global.TD27_DB.Create(perm).Error)
	rp := &modelSysManagement.RolePermissionModel{RoleID: roleID, PermissionID: perm.ID}
	require.NoError(t, global.TD27_DB.Create(rp).Error)
	return perm.ID
}

func TestButtonService_Create(t *testing.T) {
	setupButtonTest(t)
	svc := NewButtonService()

	button, err := svc.Create(&modelSysManagement.CreateButtonReq{
		ButtonCode: "user:btn-create-1",
		ButtonName: "创建用户",
		PagePath:   "/user",
	})
	require.NoError(t, err)
	require.NotZero(t, button.ID)

	// A button-domain permission row must be created and linked to the button
	var perm modelSysManagement.PermissionModel
	err = global.TD27_DB.Where("domain_id = ? AND domain = ?", button.ID, modelSysManagement.PermissionDomainButton).
		First(&perm).Error
	require.NoError(t, err)
	assert.Equal(t, "user:btn-create-1", perm.Resource)
	assert.Equal(t, modelSysManagement.ActionExecute, perm.Action)

	// Duplicate code must be rejected
	_, err = svc.Create(&modelSysManagement.CreateButtonReq{
		ButtonCode: "user:btn-create-1",
		ButtonName: "duplicate",
		PagePath:   "/user",
	})
	assert.Error(t, err)
}

func TestButtonService_CheckPermission(t *testing.T) {
	setupButtonTest(t)
	svc := NewButtonService()

	seedButtonPermission(t, "user:btn-check-1", 10)
	seedButtonPermission(t, "user:btn-check-2", 11)

	assert.True(t, svc.CheckPermission("user:btn-check-1", []uint{10}))
	assert.False(t, svc.CheckPermission("user:btn-check-1", []uint{11}))
	// No roles at all must always be false (fail-closed)
	assert.False(t, svc.CheckPermission("user:btn-check-1", []uint{}))
	assert.False(t, svc.CheckPermission("nonexistent-code", []uint{10}))
}

func TestButtonService_BatchCheckPermission(t *testing.T) {
	setupButtonTest(t)
	svc := NewButtonService()

	seedButtonPermission(t, "user:btn-batch-yes", 20)

	result := svc.BatchCheckPermission(
		[]string{"user:btn-batch-yes", "user:btn-batch-no"},
		[]uint{20},
	)
	assert.Equal(t, map[string]bool{
		"user:btn-batch-yes": true,
		"user:btn-batch-no":  false,
	}, result)

	// Unknown codes and empty roles must all be false
	result = svc.BatchCheckPermission([]string{"user:btn-batch-yes"}, []uint{})
	assert.Equal(t, map[string]bool{"user:btn-batch-yes": false}, result)

	result = svc.BatchCheckPermission(nil, []uint{20})
	assert.Empty(t, result)
}

func TestButtonService_GetUserButtons(t *testing.T) {
	setupButtonTest(t)
	svc := NewButtonService()

	seedButtonPermission(t, "user:btn-my-1", 30)
	seedButtonPermission(t, "user:btn-my-2", 30)
	seedButtonPermission(t, "user:btn-other", 31)

	codes, err := svc.GetUserButtons([]uint{30})
	require.NoError(t, err)
	assert.ElementsMatch(t, []string{"user:btn-my-1", "user:btn-my-2"}, codes)

	// Multiple roles must be merged without duplicates
	codes, err = svc.GetUserButtons([]uint{30, 31})
	require.NoError(t, err)
	assert.ElementsMatch(t, []string{"user:btn-my-1", "user:btn-my-2", "user:btn-other"}, codes)

	// No roles must return an empty slice
	codes, err = svc.GetUserButtons([]uint{})
	require.NoError(t, err)
	assert.Empty(t, codes)
}

func TestButtonService_GetPageButtons(t *testing.T) {
	setupButtonTest(t)
	svc := NewButtonService()

	allowed, err := svc.Create(&modelSysManagement.CreateButtonReq{
		ButtonCode: "user:btn-page-yes",
		ButtonName: "允许按钮",
		PagePath:   "/page-buttons-test",
	})
	require.NoError(t, err)
	denied, err := svc.Create(&modelSysManagement.CreateButtonReq{
		ButtonCode: "user:btn-page-no",
		ButtonName: "拒绝按钮",
		PagePath:   "/page-buttons-test",
	})
	require.NoError(t, err)

	// Grant the role only the first button via its auto-created permission
	var perm modelSysManagement.PermissionModel
	require.NoError(t, global.TD27_DB.Where("domain_id = ? AND domain = ?", allowed.ID, modelSysManagement.PermissionDomainButton).
		First(&perm).Error)
	require.NoError(t, global.TD27_DB.Create(&modelSysManagement.RolePermissionModel{RoleID: 40, PermissionID: perm.ID}).Error)
	assert.NotZero(t, denied.ID)

	buttons, err := svc.GetPageButtons("/page-buttons-test", []uint{40})
	require.NoError(t, err)
	require.Len(t, buttons, 2)
	for _, btn := range buttons {
		switch btn.ButtonCode {
		case "user:btn-page-yes":
			assert.True(t, btn.HasPermission)
		case "user:btn-page-no":
			assert.False(t, btn.HasPermission)
		}
	}

	// No roles: every button must be marked as not permitted
	buttons, err = svc.GetPageButtons("/page-buttons-test", []uint{})
	require.NoError(t, err)
	require.Len(t, buttons, 2)
	for _, btn := range buttons {
		assert.False(t, btn.HasPermission)
	}
}
