const images = import.meta.glob('../assets/**/*.{png,jpg,jpeg,webp}', { eager: true })

export function getImage(name) {
  try {
    // find the path include the name.
    const imagePath = Object.keys(images).find(path => path.includes(name))
    if (!imagePath) {
      console.error(`Image not found: ${name}`)
      return ''
    }
    
    return images[imagePath].default
  } catch (error) {
    console.error(`Error loading image: ${name}`, error)
    return ''
  }
}