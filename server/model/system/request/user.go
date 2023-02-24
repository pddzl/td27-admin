package request

// Login User login structure
type Login struct {
	Username  string `json:"username" validate:"required"`  // 用户名
	Password  string `json:"password" validate:"required"`  // 密码
	Captcha   string `json:"captcha" validate:"required"`   // 验证码
	CaptchaId string `json:"captchaId" validate:"required"` // 验证码ID
}

type AddUser struct {
	Username string `json:"username" validate:"required"` // 用户名
	Password string `json:"password" validate:"required"` // 密码
	//Phone       string `json:"phone" validate:"required,regexp=^[1][0-9]{10}$"` // 手机号
	Phone       string `json:"phone"`                      // 手机号
	Email       string `json:"email" validate:"email"`     // 邮箱
	Active      bool   `json:"active"`                     // 是否活跃
	RoleModelID uint   `json:"roleId" validate:"required"` // 角色ID
}

type EditUser struct {
	Id          uint   `json:"id" validate:"required"`
	Username    string `json:"username" validate:"required"` // 用户名
	Phone       string `json:"phone"`                        // 手机号
	Email       string `json:"email"`                        // 邮箱
	Active      bool   `json:"active"`                       // 是否活跃
	RoleModelID uint   `json:"roleId" validate:"required"`   // 角色ID
}

type ModifyPass struct {
	Id          uint   `json:"id" validate:"required"`
	OldPassword string `json:"oldPassword" validate:"required"` // 旧密码
	NewPassword string `json:"newPassword" validate:"required"` // 新密码
}

type SwitchActive struct {
	Id     uint `json:"id" validate:"required"`
	Active bool `json:"active"`
}
