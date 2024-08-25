export function getImage(name) {
    return () => import(`@/assets/category/${name}.png`)
  }