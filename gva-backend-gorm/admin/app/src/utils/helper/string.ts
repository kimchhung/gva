export function singularize(str: string) {
  return str.replace(/ies$/, 'y').replace(/s$/, '');
}
