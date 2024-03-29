package base

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"time"

	"server/global"
	modelAuthority "server/model/authority"
	modelBase "server/model/base"
	baseReq "server/model/base/request"
	baseRes "server/model/base/response"
	commonRes "server/model/common/response"
	"server/utils"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

type LogRegApi struct{}

// Captcha
// @Tags      LogRegApi
// @Summary   生成验证码
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200  {object}  response.Response{data=baseReq.CaptchaResponse,msg=string}
// @Router    /logReg/captcha [post]
func (ba *LogRegApi) Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.TD27_CONFIG.Captcha.ImgHeight, global.TD27_CONFIG.Captcha.ImgWidth, global.TD27_CONFIG.Captcha.KeyLong, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v9下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		global.TD27_LOG.Error("验证码获取失败!", zap.Error(err))
		commonRes.FailWithMessage("验证码获取失败", c)
		return
	}
	commonRes.OkWithDetailed(baseReq.CaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.TD27_CONFIG.Captcha.KeyLong,
	}, "验证码获取成功", c)
}

// Login
// @Tags     LogRegApi
// @Summary  用户登录
// @accept    application/json
// @Produce   application/json
// @Param    data  body      baseReq.Login true "请求参数"
// @Success  200   {object}  response.Response{data=baseRes.LoginResponse,msg=string}
// @Router   /logReg/login [post]
func (ba *LogRegApi) Login(c *gin.Context) {
	var login baseReq.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		commonRes.FailReq(err.Error(), c)
		return
	}

	// 验证码
	if store.Verify(login.CaptchaId, login.Captcha, true) {
		u := &modelAuthority.UserModel{Username: login.Username, Password: login.Password}
		user, err := logRegService.Login(u)
		if err != nil {
			commonRes.FailWithMessage(fmt.Sprintf("登录失败: %s", err.Error()), c)
			global.TD27_LOG.Error("登录失败", zap.Error(err))
			return
		}
		// 获取token
		tokenNext(c, user)
	} else {
		commonRes.FailWithMessage("验证码错误", c)
	}
}

// 生成jwt token
func tokenNext(c *gin.Context, user *modelAuthority.UserModel) {
	j := &utils.JWT{SigningKey: []byte(global.TD27_CONFIG.JWT.SigningKey)} // 唯一签名

	claims := baseReq.CustomClaims{
		ID:         user.ID,
		Username:   user.Username,
		RoleId:     user.RoleModelID,
		BufferTime: global.TD27_CONFIG.JWT.BufferTime,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(-time.Duration(1000))),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.TD27_CONFIG.JWT.ExpiresTime) * time.Second)),
			Issuer:    global.TD27_CONFIG.JWT.Issuer,
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		commonRes.FailWithMessage("创建token失败", c)
		global.TD27_LOG.Error("创建token失败", zap.Error(err))
		return
	}

	// 是否开启多点登录
	// true: 只允许账号单点登录，后续登录的会挤掉前面的
	// false: 允许账号多点登录
	if !global.TD27_CONFIG.System.UseMultipoint {
		commonRes.OkWithDetailed(baseRes.LoginResponse{
			User:      *user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix(),
		}, "登录成功", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(user.Username, token); err != nil {
			commonRes.FailWithMessage("设置登录状态失败", c)
			global.TD27_LOG.Error("设置登录状态失败", zap.Error(err))
			return
		}

		commonRes.OkWithDetailed(baseRes.LoginResponse{
			User:      *user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix(),
		}, "登录成功", c)
	} else if err != nil {
		commonRes.FailWithMessage("设置登录状态失败", c)
		global.TD27_LOG.Error("设置登录状态失败!", zap.Error(err))
	} else {
		var blackJWT modelBase.JwtBlackListModel
		blackJWT.Jwt = jwtStr
		if err = jwtService.JoinInBlacklist(blackJWT); err != nil {
			commonRes.FailWithMessage("jwt作废失败", c)
			return
		}
		if err = jwtService.SetRedisJWT(user.Username, token); err != nil {
			commonRes.FailWithMessage("设置登录状态失败", c)
			return
		}
		commonRes.OkWithDetailed(baseRes.LoginResponse{
			User:      *user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix(),
		}, "登录成功", c)
	}
}
