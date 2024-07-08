import { NowAPI } from '.'

declare global {
  interface API {
    now: NowAPI
  }
}
