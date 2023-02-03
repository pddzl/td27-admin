const modules = import.meta.glob("../views/**/*.vue")

export function dynamicImport(component: string) {
  return modules[`../views/${component}`]
}
