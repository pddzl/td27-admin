package authority

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"server/global"
	authorityReq "server/model/authority/request"
	commonReq "server/model/common/request"
	commonRes "server/model/common/response"
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
		commonRes.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	}

	if user, err := userService.GetUserInfo(userInfo.ID); err != nil {
		commonRes.FailWithMessage("获取失败", c)
	} else {
		commonRes.OkWithDetailed(user, "获取成功", c)
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
	var pageInfo commonReq.PageInfo
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := userService.GetUsers(pageInfo); err != nil {
		commonRes.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取users失败", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(commonRes.PageResult{
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
	var cId commonReq.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}

	if err := userService.DeleteUser(cId.ID); err != nil {
		commonRes.FailWithMessage("删除失败", c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		commonRes.OkWithMessage("删除成功", c)
	}
}

// AddUser
// @Tags      UserApi
// @Summary   添加用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.AddUser true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /user/addUser [post]
func (ua *UserApi) AddUser(c *gin.Context) {
	var addUser authorityReq.AddUser
	if err := c.ShouldBindJSON(&addUser); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}

	if err := userService.AddUser(addUser); err != nil {
		commonRes.FailWithMessage("添加失败", c)
		global.TD27_LOG.Error("添加失败", zap.Error(err))
	} else {
		commonRes.OkWithMessage("添加成功", c)
	}
}

// EditUser
// @Tags      UserApi
// @Summary   编辑用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.EditUser true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /user/editUser [post]
func (ua *UserApi) EditUser(c *gin.Context) {
	var editUser authorityReq.EditUser
	if err := c.ShouldBindJSON(&editUser); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}

	if user, err := userService.EditUser(editUser); err != nil {
		commonRes.FailWithMessage("编辑失败", c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		commonRes.OkWithDetailed(user, "编辑成功", c)
	}
}

// ModifyPass
// @Tags      UserApi
// @Summary   修改用户密码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.ModifyPass true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /user/modifyPass [post]
func (ua *UserApi) ModifyPass(c *gin.Context) {
	var mp authorityReq.ModifyPass
	if err := c.ShouldBindJSON(&mp); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}

	if err := userService.ModifyPass(mp); err != nil {
		commonRes.FailWithMessage("修改失败", c)
		global.TD27_LOG.Error("修改失败", zap.Error(err))
	} else {
		commonRes.OkWithMessage("修改成功", c)
	}
}

// SwitchActive
// @Tags      UserApi
// @Summary   切换启用状态
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.SwitchActive true "请求参数"
// @Success   200   {object}  response.Response{msg=string}
// @Router    /user/switchActive [post]
func (ua *UserApi) SwitchActive(c *gin.Context) {
	var sa authorityReq.SwitchActive
	if err := c.ShouldBindJSON(&sa); err != nil {
		commonRes.FailWithMessage(err.Error(), c)
		return
	}

	if err := userService.SwitchActive(sa); err != nil {
		commonRes.FailWithMessage("切换失败", c)
		global.TD27_LOG.Error("切换失败", zap.Error(err))
	} else {
		commonRes.OkWithMessage("切换成功", c)
	}
}
