/** 表单电话号码校验 */
export function useValidatePhone(rule: any, value: any, callback: any) {
  if (value === "") {
    callback()
  } else {
    const phoneReg = /^1\d{10}$/
    if (!phoneReg.test(value)) {
      callback(new Error("手机号码不合规"))
    } else {
      callback()
    }
  }
}

/** 邮箱校验 */
export function useValidateEmail(rule: any, value: any, callback: any) {
  if (value === "") {
    callback()
  } else {
    const emailReg = /^[\w-]+@[\w-]+(\.[\w-]+)+$/
    if (!emailReg.test(value)) {
      callback(new Error("邮箱不合规"))
    } else {
      callback()
    }
  }
}
