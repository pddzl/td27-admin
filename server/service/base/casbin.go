package base

import (
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"strconv"
	"sync"

	"server/global"
	baseReq "server/model/base/request"
)

type CasbinService struct{}

var CasbinServiceApp = new(CasbinService)

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

func (cs *CasbinService) Casbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(global.TD27_DB)
		if err != nil {
			zap.L().Error("适配数据库失败请检查casbin表是否为InnoDB引擎!", zap.Error(err))
			return
		}
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			zap.L().Error("字符串加载模型失败!", zap.Error(err))
			return
		}
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, a)
		syncedCachedEnforcer.SetExpireTime(60 * 60)
		_ = syncedCachedEnforcer.LoadPolicy()
	})
	return syncedCachedEnforcer
}

// UpdateCasbinApi 更新api权限
func (cs *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := global.TD27_DB.Model(&gormadapter.CasbinRule{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	e := cs.Casbin()
	err = e.LoadPolicy()
	if err != nil {
		return err
	}
	return err
}

// ClearCasbin 清除casbin rule
func (cs *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := cs.Casbin()
	ok, _ := e.RemoveFilteredPolicy(v, p...)
	return ok
}

// EditCasbin 更新casbin rule
func (cs *CasbinService) EditCasbin(roleId uint, casbinInfos []baseReq.CasbinInfo) (err error) {
	authorityId := strconv.Itoa(int(roleId))
	cs.ClearCasbin(0, authorityId)
	var rules [][]string
	// casbin rule 去重
	deduplicateMap := make(map[string]struct{})
	for _, v := range casbinInfos {
		key := authorityId + v.Path + v.Method
		if _, ok := deduplicateMap[key]; !ok {
			deduplicateMap[key] = struct{}{}
			rules = append(rules, []string{authorityId, v.Path, v.Method})
		}
	}
	e := cs.Casbin()
	ok, _ := e.AddPolicies(rules)
	if !ok {
		return errors.New("存在相同api")
	}
	return
}
