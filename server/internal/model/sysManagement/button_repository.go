package sysManagement

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"

	"server/internal/global"
)

type ButtonRepository interface {
	Create(ctx context.Context, button *ButtonModel) error
	Update(ctx context.Context, button *ButtonModel) error
	Delete(ctx context.Context, id uint) error
	FindByID(ctx context.Context, id uint) (*ButtonModel, error)
	FindByCode(ctx context.Context, code string) (*ButtonModel, error)
	List(ctx context.Context, req *ListButtonReq) ([]ButtonModel, int64, error)
	GetByPagePath(ctx context.Context, pagePath string) ([]ButtonModel, error)
}

type buttonRepo struct{}

func NewButtonRepo() ButtonRepository {
	return &buttonRepo{}
}

func (r *buttonRepo) Create(ctx context.Context, button *ButtonModel) error {
	if err := global.TD27_DB.WithContext(ctx).Create(button).Error; err != nil {
		return fmt.Errorf("create button failed: %w", err)
	}
	return nil
}

func (r *buttonRepo) Update(ctx context.Context, button *ButtonModel) error {
	if err := global.TD27_DB.WithContext(ctx).Save(button).Error; err != nil {
		return fmt.Errorf("update button failed: %w", err)
	}
	return nil
}

func (r *buttonRepo) Delete(ctx context.Context, id uint) error {
	result := global.TD27_DB.WithContext(ctx).Delete(&ButtonModel{}, id)
	if result.Error != nil {
		return fmt.Errorf("delete button failed: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return errors.New("button not found")
	}
	return nil
}

func (r *buttonRepo) FindByID(ctx context.Context, id uint) (*ButtonModel, error) {
	var button ButtonModel
	if err := global.TD27_DB.WithContext(ctx).First(&button, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("button not found")
		}
		return nil, fmt.Errorf("find button failed: %w", err)
	}
	return &button, nil
}

func (r *buttonRepo) FindByCode(ctx context.Context, code string) (*ButtonModel, error) {
	var button ButtonModel
	if err := global.TD27_DB.WithContext(ctx).Where("button_code = ?", code).First(&button).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("button not found")
		}
		return nil, fmt.Errorf("find button by code failed: %w", err)
	}
	return &button, nil
}

func (r *buttonRepo) List(ctx context.Context, req *ListButtonReq) ([]ButtonModel, int64, error) {
	var buttons []ButtonModel
	var total int64

	db := global.TD27_DB.WithContext(ctx).Model(&ButtonModel{})

	if req.PagePath != "" {
		db = db.Where("page_path = ?", req.PagePath)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("count buttons failed: %w", err)
	}

	req.Normalize()

	if err := db.Order("page_path, id").Limit(req.PageSize).Offset(req.Offset()).Find(&buttons).Error; err != nil {
		return nil, 0, fmt.Errorf("list buttons failed: %w", err)
	}

	return buttons, total, nil
}

func (r *buttonRepo) GetByPagePath(ctx context.Context, pagePath string) ([]ButtonModel, error) {
	var buttons []ButtonModel
	if err := global.TD27_DB.WithContext(ctx).Where("page_path = ?", pagePath).Find(&buttons).Error; err != nil {
		return nil, fmt.Errorf("get buttons by page path failed: %w", err)
	}
	return buttons, nil
}
