package sysManagement

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"server/internal/global"
	"server/internal/model/common"
	modelSysManagement "server/internal/model/sysManagement"
	serviceSysManagement "server/internal/service/sysManagement"
)

type UserApi struct {
	userService *serviceSysManagement.UserService
}

func NewUserApi() *UserApi {
	return &UserApi{userService: serviceSysManagement.NewUserService()}
}

// GetUserInfo
// @Tags      UserApi
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200   {object}  common.Response{msg=string}
// @Router    /user/getUserInfo [post]
func (ua *UserApi) GetUserInfo(c *gin.Context) {
	userInfo, err := GetUserInfo(c)
	if err != nil {
		common.FailWithMessage("获取失败", c)
	}

	if user, err := ua.userService.GetUserInfo(userInfo.ID); err != nil {
		common.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败", zap.Error(err))
	} else {
		common.OkWithDetailed(user, "获取成功", c)
	}
}

// List
// @Tags      UserApi
// @Summary   分页获取用户（支持数据权限）
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      common.PageInfo true "请求参数"
// @Success   200   {object}  common.Response{data=[]modelSysManagement.UserResp,msg=string}
// @Router    /user/list [post]
func (ua *UserApi) List(c *gin.Context) {
	var pageInfo common.PageInfo
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	// 获取当前用户ID
	userInfo, err := GetUserInfo(c)
	if err != nil {
		common.FailWithMessage("获取当前用户失败", c)
		return
	}

	if list, total, err := ua.userService.List(&pageInfo, userInfo.ID); err != nil {
		common.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("get users failed", zap.Error(err))
	} else {
		common.OkWithDetailed(common.Page{
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
			Total:    total,
			List:     list,
		}, "获取成功", c)
	}
}

// Delete
// @Tags      UserApi
// @Summary   删除用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      common.CId true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /user/delete [post]
func (ua *UserApi) Delete(c *gin.Context) {
	var cId common.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := ua.userService.Delete(cId.ID); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		common.Ok(c)
	}
}

// Create
// @Tags      UserApi
// @Summary   添加用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysManagement.AddUserReq true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /user/create [post]
func (ua *UserApi) Create(c *gin.Context) {
	// 注册自定义校验函数
	validate := validator.New()
	err := validate.RegisterValidation("phone", modelSysManagement.PhoneValidation)
	if err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	// 使用 Gin 的验证器替换为自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err = v.RegisterValidation("phone", modelSysManagement.PhoneValidation)
		if err != nil {
			common.FailReq(err.Error(), c)
			return
		}
	}

	var req modelSysManagement.AddUserReq
	if err = c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err = ua.userService.Create(&req); err != nil {
		common.FailWithMessage("添加失败", c)
		global.TD27_LOG.Error("添加失败", zap.Error(err))
	} else {
		common.Ok(c)
	}
}

// Update
// @Tags      UserApi
// @Summary   编辑用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysManagement.UpdateUserReq true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /user/update [post]
func (ua *UserApi) Update(c *gin.Context) {
	// 注册自定义校验函数
	validate := validator.New()
	err := validate.RegisterValidation("phone", modelSysManagement.PhoneValidation)
	if err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	// 使用 Gin 的验证器替换为自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err = v.RegisterValidation("phone", modelSysManagement.PhoneValidation)
		if err != nil {
			common.FailReq(err.Error(), c)
			return
		}
	}

	var req modelSysManagement.UpdateUserReq
	if err = c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if instance, err := ua.userService.Update(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		common.OkWithDetailed(instance, "编辑成功", c)
	}
}

// ModifyPasswd
// @Tags      UserApi
// @Summary   修改用户密码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysManagement.ModifyPasswdReq true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /user/modifyPasswd [post]
func (ua *UserApi) ModifyPasswd(c *gin.Context) {
	var req modelSysManagement.ModifyPasswdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := ua.userService.ModifyPasswd(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("修改失败", zap.Error(err))
	} else {
		common.Ok(c)
	}
}

// SwitchActive
// @Tags      UserApi
// @Summary   切换启用状态
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      modelSysManagement.SwitchActiveReq true "请求参数"
// @Success   200   {object}  common.Response{msg=string}
// @Router    /user/switchActive [post]
func (ua *UserApi) SwitchActive(c *gin.Context) {
	var req modelSysManagement.SwitchActiveReq
	if err := c.ShouldBindJSON(&req); err != nil {
		common.FailReq(err.Error(), c)
		return
	}

	if err := ua.userService.SwitchActive(&req); err != nil {
		common.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("切换失败", zap.Error(err))
	} else {
		common.Ok(c)
	}
}
