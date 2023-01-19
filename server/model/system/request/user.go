package request

// Login User login structure
type Login struct {
	Username  string `json:"username" validate:"required"`  // 用户名
	Password  string `json:"password" validate:"required"`  // 密码
	Captcha   string `json:"captcha" validate:"required"`   // 验证码
	CaptchaId string `json:"captchaId" validate:"required"` // 验证码ID
}
