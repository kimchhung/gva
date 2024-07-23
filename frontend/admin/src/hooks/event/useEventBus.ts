import mitt from 'mitt'
import { onBeforeUnmount } from 'vue'

type Option = {
  name: string // Event name
  callback: Fn // Call back
}

const emitter = mitt()

export const useEventBus = (option?: Option) => {
  if (option) {
    emitter.on(option.name, option.callback)

    onBeforeUnmount(() => {
      emitter.off(option.name)
    })
  }

  return {
    on: emitter.on,
    off: emitter.off,
    emit: emitter.emit,
    all: emitter.all
  }
}
