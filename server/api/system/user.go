package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/request"
	"server/model/common/response"
	systemReq "server/model/system/request"
	"server/utils"
)

type UserApi struct{}

// GetUserInfo
// @Tags      UserApi
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200   {object}  response.Response{msg=string}
// @Router    /user/getUserInfo [post]
func (ua *UserApi) GetUserInfo(c *gin.Context) {
	userInfo, err := utils.GetUserInfo(c)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	}

	if user, err := userService.GetUserInfo(userInfo.ID); err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(user, "获取成功", c)
	}
}

// GetUsers
// @Tags      UserApi
// @Summary   分页获取用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.PageInfo true "请求参数"
// @Success   200   {object}  response.Response{data=response.PageResult{list=[]response.UserResult},msg=string}
// @Router    /user/getUsers [post]
func (ua *UserApi) GetUsers(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)

	if list, total, err := userService.GetUsers(pageInfo); err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取users失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
			Total:    total,
			List:     list,
		}, "获取成功", c)
	}
}

// DeleteUser
// @Tags      UserApi
// @Summary   删除用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.CId true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /user/deleteUser [post]
func (ua *UserApi) DeleteUser(c *gin.Context) {
	var cId request.CId
	_ = c.ShouldBindJSON(&cId)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&cId); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := userService.DeleteUser(cId.ID); err != nil {
		response.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// AddUser
// @Tags      UserApi
// @Summary   添加用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.AddUser true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /user/addUser [post]
func (ua *UserApi) AddUser(c *gin.Context) {
	var addUser systemReq.AddUser
	_ = c.ShouldBindJSON(&addUser)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&addUser); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := userService.AddUser(addUser); err != nil {
		response.FailWithMessage("添加失败", c)
		global.TD27_LOG.Error("添加失败", zap.Error(err))
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

// EditUser
// @Tags      UserApi
// @Summary   编辑用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.EditUser true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /user/editUser [post]
func (ua *UserApi) EditUser(c *gin.Context) {
	var editUser systemReq.EditUser
	_ = c.ShouldBindJSON(&editUser)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&editUser); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if user, err := userService.EditUser(editUser); err != nil {
		response.FailWithMessage("编辑失败", c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		response.OkWithDetailed(user, "编辑成功", c)
	}
}

// ModifyPass
// @Tags      UserApi
// @Summary   修改用户密码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.ModifyPass true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /user/modifyPass [post]
func (ua *UserApi) ModifyPass(c *gin.Context) {
	var mp systemReq.ModifyPass
	_ = c.ShouldBindJSON(&mp)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&mp); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := userService.ModifyPass(mp); err != nil {
		response.FailWithMessage("修改失败", c)
		global.TD27_LOG.Error("修改失败", zap.Error(err))
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// SwitchActive
// @Tags      UserApi
// @Summary   切换启用状态
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      systemReq.SwitchActive true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /user/switchActive [post]
func (ua *UserApi) SwitchActive(c *gin.Context) {
	var sa systemReq.SwitchActive
	_ = c.ShouldBindJSON(&sa)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&sa); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := userService.SwitchActive(sa); err != nil {
		response.FailWithMessage("切换失败", c)
		global.TD27_LOG.Error("切换失败", zap.Error(err))
	} else {
		response.OkWithMessage("切换成功", c)
	}
}
