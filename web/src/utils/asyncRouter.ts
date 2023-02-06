const modules = import.meta.glob("../views/**/*.vue", { eager: true })

export function dynamicImport(component: string) {
  console.log("here")
  return modules[`../views/${component}`]
}
