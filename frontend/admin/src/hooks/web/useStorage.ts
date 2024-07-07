// 获取传入的值的类型
const getValueType = (value: any) => {
  const type = Object.prototype.toString.call(value)
  return type.slice(8, -1)
}

export type StorageKeyValue = {
  lang: LocaleType | null
}

export const useStorage = (type: 'sessionStorage' | 'localStorage' = 'sessionStorage') => {
  const setStorage = <K extends keyof StorageKeyValue>(key: K, value: StorageKeyValue[K]) => {
    const valueType = getValueType(value)
    window[type].setItem(key, JSON.stringify({ type: valueType, value }))
  }

  const getStorage = <K extends keyof StorageKeyValue>(key: K): StorageKeyValue[K] => {
    const value = window[type].getItem(key)
    if (value) {
      const { value: val } = JSON.parse(value)
      return val
    } else {
      console.error('get invalid value type:', value)
      return null
    }
  }

  const removeStorage = (key: string) => {
    window[type].removeItem(key)
  }

  const clear = (excludes?: string[]) => {
    const keys = Object.keys(window[type])
    const defaultExcludes = ['dynamicRouter', 'serverDynamicRouter']
    const excludesArr = excludes ? [...excludes, ...defaultExcludes] : defaultExcludes
    const excludesKeys = excludesArr ? keys.filter((key) => !excludesArr.includes(key)) : keys

    excludesKeys.forEach((key) => {
      window[type].removeItem(key)
    })
    // window[type].clear()
  }

  return {
    setStorage,
    getStorage,
    removeStorage,
    clear
  }
}
