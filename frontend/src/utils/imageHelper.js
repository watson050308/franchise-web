export function getImage(name) {
    return () => import(`@/assets/${name}.png`)
  }