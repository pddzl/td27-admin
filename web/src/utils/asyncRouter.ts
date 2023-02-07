const modules = import.meta.glob("../views/**/*.vue", { eager: true })

export function dynamicImport(component: string) {
  return modules[`../views/${component}`]
}
