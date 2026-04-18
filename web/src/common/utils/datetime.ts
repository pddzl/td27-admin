import dayjs from "dayjs"

const INVALID_DATE = "N/A"

/** 格式化日期时间 */
export function formatDateTime(
  datetime: string | number | Date = "",
  template: string = "YYYY-MM-DD HH:mm:ss"
) {
  let value = datetime

  // 👇 handle unix seconds
  if (typeof datetime === "number" && datetime < 1e12) {
    value = datetime * 1000
  }

  const day = dayjs(value)
  return day.isValid() ? day.format(template) : INVALID_DATE
}
