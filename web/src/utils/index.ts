/** 将全局 CSS 变量导入 JS 中使用 */
export const getCssVariableValue = (cssVariableName: string) => {
  let cssVariableValue = ""
  try {
    // 没有拿到值时，会返回空串
    cssVariableValue = getComputedStyle(document.documentElement).getPropertyValue(cssVariableName)
  } catch (error) {
    console.error(error)
  }
  return cssVariableValue
}
