package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"time"

	"server/global"
	"server/model/common/response"
	"server/model/system"
	systemReq "server/model/system/request"
	systemRes "server/model/system/response"
	"server/utils"
)

// 当开启多服务器部署时，替换下面的配置，使用redis共享存储验证码
// var store = captcha.NewDefaultRedisStore()
var store = base64Captcha.DefaultMemStore

type BaseApi struct{}

// Captcha
// @Tags      Base
// @Summary   生成验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=systemRes.SysCaptchaResponse,msg=string}  "生成验证码,返回包括随机数id,base64,验证码长度"
// @Router    /base/captcha [post]
func (ba *BaseApi) Captcha(c *gin.Context) {
	// 字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(global.TD27_CONFIG.Captcha.ImgHeight, global.TD27_CONFIG.Captcha.ImgWidth, global.TD27_CONFIG.Captcha.KeyLong, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // v8下使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		global.TD27_LOG.Error("验证码获取失败!", zap.Error(err))
		response.FailWithMessage("验证码获取失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysCaptchaResponse{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: global.TD27_CONFIG.Captcha.KeyLong,
	}, "验证码获取成功", c)
}

// Login
// @Tags     Base
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      systemReq.Login                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  response.Response{data=systemRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /base/login [post]
func (ba *BaseApi) Login(c *gin.Context) {
	var login systemReq.Login
	_ = c.ShouldBindJSON(&login)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&login); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.TD27_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if store.Verify(login.CaptchaId, login.Captcha, true) {
		u := &system.UserModel{Username: login.Username, Password: login.Password}
		user, err := userService.Login(u)
		if err != nil {
			global.TD27_LOG.Error("登陆失败", zap.Error(err))
			response.FailWithMessage("登陆失败", c)
			return
		}
		tokenNext(c, user)
	} else {
		response.FailWithMessage("验证码错误", c)
	}
}

func tokenNext(c *gin.Context, user *system.UserModel) {
	j := &utils.JWT{SigningKey: []byte(global.TD27_CONFIG.JWT.SigningKey)} // 唯一签名

	claims := systemReq.CustomClaims{
		UUID:       user.UUID,
		ID:         user.ID,
		Username:   user.Username,
		BufferTime: global.TD27_CONFIG.JWT.BufferTime,
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(-time.Duration(1000))),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(global.TD27_CONFIG.JWT.ExpiresTime) * time.Second)),
			Issuer:    global.TD27_CONFIG.JWT.Issuer,
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		response.FailWithMessage("获取token失败", c)
		global.TD27_LOG.Error("获取token失败", zap.Error(err))
		return
	}

	if !global.TD27_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      *user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix(),
		}, "登录成功", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(user.Username, token); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			global.TD27_LOG.Error("设置登录状态失败", zap.Error(err))
			return
		}

		response.OkWithDetailed(systemRes.LoginResponse{
			User:      *user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix(),
		}, "登录成功", c)
	} else if err != nil {
		response.FailWithMessage("设置登录状态失败", c)
		global.TD27_LOG.Error("设置登录状态失败!", zap.Error(err))
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JoinInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(user.Username, token); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      *user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix(),
		}, "登录成功", c)
	}
}
