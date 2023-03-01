/** 表单电话号码校验 */
export const useValidatePhone = (rule: any, value: any, callback: any) => {
  if (value === "") {
    callback()
  } else {
    const phoneReg = /^[1][0-9]{10}$/
    if (!phoneReg.test(value)) {
      callback(new Error("手机号码不合规"))
    } else {
      callback()
    }
  }
}

/** 邮箱校验 */
export const useValidateEmail = (rule: any, value: any, callback: any) => {
  if (value === "") {
    callback()
  } else {
    const emailReg = /^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/
    if (!emailReg.test(value)) {
      callback(new Error("邮箱不合规"))
    } else {
      callback()
    }
  }
}
