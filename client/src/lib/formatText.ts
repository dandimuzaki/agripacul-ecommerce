export const capitalize = (text: string): string => {
  return text.charAt(0).toUpperCase() + text.slice(1)
}

export const capitalizeAll = (text: string): string => {
  const words = text.split(' ')
  const capitalized = words.map(w => w.charAt(0).toUpperCase() + w.slice(1))
  return capitalized.join(' ')
}