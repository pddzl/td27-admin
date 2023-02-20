package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/request"
	"server/model/common/response"
	systemModel "server/model/system"
	"server/utils"
)

type UserApi struct{}

func (ua *UserApi) GetUserInfo(c *gin.Context) {
	if userInfo, err := utils.GetUserInfo(c); err != nil {
		response.FailWithMessage("获取失败", c)
		global.TD27_LOG.Error("获取失败!", zap.Error(err))
	} else {
		response.OkWithDetailed(userInfo, "获取成功", c)
	}
}

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

// AddUser 添加用户
func (ua *UserApi) AddUser(c *gin.Context) {
	var addUser systemModel.AddUser
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
