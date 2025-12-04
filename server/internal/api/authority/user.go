package authority

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"server/internal/global"
	commonReq "server/internal/model/common/request"
	commonResp "server/internal/model/common/response"
	authorityReq "server/internal/model/entity/authority/request"
	_ "server/internal/model/entity/authority/response"
	"server/internal/pkg"
	serviceAuthority "server/internal/service/authority"
)

type UserApi struct {
	userService *serviceAuthority.UserService
}

func NewUserApi() *UserApi {
	return &UserApi{userService: serviceAuthority.NewUserService()}
}

// GetUserInfo
// @Tags      UserApi
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200   {object}  commonResp.Response{msg=string}
// @Router    /user/getUserInfo [post]
func (ua *UserApi) GetUserInfo(c *gin.Context) {
	userInfo, err := pkg.GetUserInfo(c)
	if err != nil {
		commonResp.FailWithMessage("获取失败", c)
	}

	if user, err := ua.userService.GetUserInfo(userInfo.ID); err != nil {
		commonResp.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败", zap.Error(err))
	} else {
		commonResp.OkWithDetailed(user, "获取成功", c)
	}
}

// GetUsers
// @Tags      UserApi
// @Summary   分页获取用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      commonReq.PageInfo true "请求参数"
// @Success   200   {object}  commonResp.Response{data=commonResp.Page{list=[]response.UserResult},msg=string}
// @Router    /user/getUsers [post]
func (ua *UserApi) GetUsers(c *gin.Context) {
	var pageInfo commonReq.PageInfo
	if err := c.ShouldBindJSON(&pageInfo); err != nil {
		commonResp.FailReq(err.Error(), c)
		return
	}

	if list, total, err := ua.userService.GetUsers(pageInfo); err != nil {
		commonResp.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取users失败", zap.Error(err))
	} else {
		commonResp.OkWithDetailed(commonResp.Page{
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
// @Param     data  body      commonReq.CId true "请求参数"
// @Success   200   {object}  commonResp.Response{msg=string}
// @Router    /user/deleteUser [post]
func (ua *UserApi) DeleteUser(c *gin.Context) {
	var cId commonReq.CId
	if err := c.ShouldBindJSON(&cId); err != nil {
		commonResp.FailReq(err.Error(), c)
		return
	}

	if err := ua.userService.DeleteUser(cId.ID); err != nil {
		commonResp.Fail(c)
		global.TD27_LOG.Error("删除失败", zap.Error(err))
	} else {
		commonResp.Ok(c)
	}
}

// AddUser
// @Tags      UserApi
// @Summary   添加用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.AddUser true "请求参数"
// @Success   200   {object}  commonResp.Response{msg=string}
// @Router    /user/addUser [post]
func (ua *UserApi) AddUser(c *gin.Context) {
	// 注册自定义校验函数
	validate := validator.New()
	err := validate.RegisterValidation("phone", authorityReq.PhoneValidation)
	if err != nil {
		commonResp.FailReq(err.Error(), c)
		return
	}

	// 使用 Gin 的验证器替换为自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err = v.RegisterValidation("phone", authorityReq.PhoneValidation)
		if err != nil {
			commonResp.FailReq(err.Error(), c)
			return
		}
	}

	var addUser authorityReq.AddUser
	if err = c.ShouldBindJSON(&addUser); err != nil {
		commonResp.FailReq(err.Error(), c)
		return
	}

	if err = ua.userService.AddUser(&addUser); err != nil {
		commonResp.FailWithMessage("添加失败", c)
		global.TD27_LOG.Error("添加失败", zap.Error(err))
	} else {
		commonResp.Ok(c)
	}
}

// EditUser
// @Tags      UserApi
// @Summary   编辑用户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.EditUser true "请求参数"
// @Success   200   {object}  commonResp.Response{msg=string}
// @Router    /user/editUser [post]
func (ua *UserApi) EditUser(c *gin.Context) {
	// 注册自定义校验函数
	validate := validator.New()
	err := validate.RegisterValidation("phone", authorityReq.PhoneValidation)
	if err != nil {
		commonResp.FailReq(err.Error(), c)
		return
	}

	// 使用 Gin 的验证器替换为自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err = v.RegisterValidation("phone", authorityReq.PhoneValidation)
		if err != nil {
			commonResp.FailReq(err.Error(), c)
			return
		}
	}

	var editUser authorityReq.EditUser
	if err = c.ShouldBindJSON(&editUser); err != nil {
		commonResp.FailReq(err.Error(), c)
		return
	}

	if instance, err := ua.userService.EditUser(&editUser); err != nil {
		commonResp.Fail(c)
		global.TD27_LOG.Error("编辑失败", zap.Error(err))
	} else {
		commonResp.OkWithDetailed(instance, "编辑成功", c)
	}
}

// ModifyPass
// @Tags      UserApi
// @Summary   修改用户密码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.ModifyPass true "请求参数"
// @Success   200   {object}  commonResp.Response{msg=string}
// @Router    /user/modifyPass [post]
func (ua *UserApi) ModifyPass(c *gin.Context) {
	var mp authorityReq.ModifyPass
	if err := c.ShouldBindJSON(&mp); err != nil {
		commonResp.FailReq(err.Error(), c)
		return
	}

	if err := ua.userService.ModifyPass(&mp); err != nil {
		commonResp.FailWithMessage(err.Error(), c)
		global.TD27_LOG.Error("修改失败", zap.Error(err))
	} else {
		commonResp.Ok(c)
	}
}

// SwitchActive
// @Tags      UserApi
// @Summary   切换启用状态
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      authorityReq.SwitchActive true "请求参数"
// @Success   200   {object}  commonResp.Response{msg=string}
// @Router    /user/switchActive [post]
func (ua *UserApi) SwitchActive(c *gin.Context) {
	var sa authorityReq.SwitchActive
	if err := c.ShouldBindJSON(&sa); err != nil {
		commonResp.FailReq(err.Error(), c)
		return
	}

	if err := ua.userService.SwitchActive(&sa); err != nil {
		commonResp.Fail(c)
		global.TD27_LOG.Error("切换失败", zap.Error(err))
	} else {
		commonResp.Ok(c)
	}
}
