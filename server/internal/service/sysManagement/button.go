package sysManagement

import (
	"context"
	"fmt"

	"server/internal/global"
	modelSysManagement "server/internal/model/sysManagement"
)

type ButtonService struct {
	buttonRepo modelSysManagement.ButtonRepository
	ctx        context.Context
}

func NewButtonService() *ButtonService {
	return &ButtonService{
		buttonRepo: modelSysManagement.NewButtonRepo(),
		ctx:        context.Background(),
	}
}

func (s *ButtonService) Create(req *modelSysManagement.CreateButtonReq) (*modelSysManagement.ButtonModel, error) {
	_, err := s.buttonRepo.FindByCode(s.ctx, req.ButtonCode)
	if err == nil {
		return nil, fmt.Errorf("button code already exists: %s", req.ButtonCode)
	}

	button := &modelSysManagement.ButtonModel{
		ButtonCode:  req.ButtonCode,
		ButtonName:  req.ButtonName,
		Description: req.Description,
		PagePath:    req.PagePath,
	}

	if err := s.buttonRepo.Create(s.ctx, button); err != nil {
		return nil, err
	}

	permission := &modelSysManagement.PermissionModel{
		Name:     req.ButtonName,
		Domain:   modelSysManagement.PermissionDomainButton,
		Resource: req.ButtonCode,
		Action:   modelSysManagement.ActionExecute,
		DomainID: button.ID,
	}

	if err := global.TD27_DB.Create(permission).Error; err != nil {
		s.buttonRepo.Delete(s.ctx, button.ID)
		return nil, fmt.Errorf("create permission failed: %w", err)
	}

	return button, nil
}

func (s *ButtonService) Update(req *modelSysManagement.UpdateButtonReq) error {
	button, err := s.buttonRepo.FindByID(s.ctx, req.ID)
	if err != nil {
		return err
	}

	if button.ButtonCode != req.ButtonCode {
		_, err := s.buttonRepo.FindByCode(s.ctx, req.ButtonCode)
		if err == nil {
			return fmt.Errorf("button code already exists: %s", req.ButtonCode)
		}
	}

	button.ButtonCode = req.ButtonCode
	button.ButtonName = req.ButtonName
	button.Description = req.Description
	button.PagePath = req.PagePath

	if err := s.buttonRepo.Update(s.ctx, button); err != nil {
		return err
	}

	global.TD27_DB.Model(&modelSysManagement.PermissionModel{}).
		Where("domain_id = ? AND domain = ?", button.ID, modelSysManagement.PermissionDomainButton).
		Update("name", req.ButtonName)

	return nil
}

func (s *ButtonService) Delete(id uint) error {
	global.TD27_DB.Where("domain_id = ? AND domain = ?", id, modelSysManagement.PermissionDomainButton).
		Delete(&modelSysManagement.PermissionModel{})
	return s.buttonRepo.Delete(s.ctx, id)
}

func (s *ButtonService) List(req *modelSysManagement.ListButtonReq) ([]modelSysManagement.ButtonModel, int64, error) {
	return s.buttonRepo.List(s.ctx, req)
}

func (s *ButtonService) GetPageButtons(pagePath string, roleIDs []uint) ([]modelSysManagement.ButtonDto, error) {
	buttons, err := s.buttonRepo.GetByPagePath(s.ctx, pagePath)
	if err != nil {
		return nil, err
	}

	var permittedCodes []string
	if len(roleIDs) > 0 {
		global.TD27_DB.Raw(`
			SELECT DISTINCT p.resource
			FROM sys_management_permission p
			JOIN sys_management_role_permissions rp ON p.id = rp.permission_id
			WHERE rp.role_id IN ? AND p.domain = 'button'
		`, roleIDs).Scan(&permittedCodes)
	}

	permittedMap := make(map[string]bool)
	for _, code := range permittedCodes {
		permittedMap[code] = true
	}

	result := make([]modelSysManagement.ButtonDto, 0, len(buttons))
	for _, btn := range buttons {
		result = append(result, modelSysManagement.ButtonDto{
			ID:            btn.ID,
			ButtonCode:    btn.ButtonCode,
			ButtonName:    btn.ButtonName,
			Description:   btn.Description,
			PagePath:      btn.PagePath,
			HasPermission: permittedMap[btn.ButtonCode],
		})
	}

	return result, nil
}

func (s *ButtonService) CheckPermission(buttonCode string, roleIDs []uint) bool {
	if len(roleIDs) == 0 {
		return false
	}
	var count int64
	global.TD27_DB.Raw(`
		SELECT COUNT(*) FROM sys_management_permission p
		JOIN sys_management_role_permissions rp ON p.id = rp.permission_id
		WHERE rp.role_id IN ? AND p.domain = 'button' AND p.resource = ?
	`, roleIDs, buttonCode).Scan(&count)
	return count > 0
}

func (s *ButtonService) BatchCheckPermission(buttonCodes []string, roleIDs []uint) map[string]bool {
	result := make(map[string]bool)
	for _, code := range buttonCodes {
		result[code] = false
	}
	if len(roleIDs) == 0 || len(buttonCodes) == 0 {
		return result
	}
	var permittedCodes []string
	global.TD27_DB.Raw(`
		SELECT DISTINCT p.resource FROM sys_management_permission p
		JOIN sys_management_role_permissions rp ON p.id = rp.permission_id
		WHERE rp.role_id IN ? AND p.domain = 'button' AND p.resource IN ?
	`, roleIDs, buttonCodes).Scan(&permittedCodes)
	for _, code := range permittedCodes {
		result[code] = true
	}
	return result
}

func (s *ButtonService) GetUserButtons(roleIDs []uint) ([]string, error) {
	if len(roleIDs) == 0 {
		return []string{}, nil
	}
	var buttonCodes []string
	err := global.TD27_DB.Raw(`
		SELECT DISTINCT p.resource FROM sys_management_permission p
		JOIN sys_management_role_permissions rp ON p.id = rp.permission_id
		WHERE rp.role_id IN ? AND p.domain = 'button'
	`, roleIDs).Scan(&buttonCodes).Error
	return buttonCodes, err
}
