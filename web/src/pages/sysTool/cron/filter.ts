const strategyMap: Record<string, string> = {
  always: "重复执行",
  once: "执行一次"
}

export const strategyFilter = (strategy: string) => {
  return strategyMap[strategy] || "未知"
}
