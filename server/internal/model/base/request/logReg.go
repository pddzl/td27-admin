package request

type Login struct {
	Username  string `json:"username" binding:"required"`  // 用户名
	Password  string `json:"password" binding:"required"`  // 密码
	Captcha   string `json:"captcha" binding:"required"`   // 验证码
	CaptchaId string `json:"captchaId" binding:"required"` // 验证码ID
}

type CaptchaResponse struct {
	CaptchaId     string `json:"captchaId"`
	PicPath       string `json:"picPath"`
	CaptchaLength int    `json:"captchaLength"`
}
