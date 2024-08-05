import { type ComputedRef, type MaybeRef } from 'vue';

/**
 * Deep recursive all attributes are optional
 */
type DeepPartial<T> = T extends object
  ? {
      [P in keyof T]?: DeepPartial<T[P]>;
    }
  : T;

/**
 * Deep recursive all attributes are read only
 */
type DeepReadonly<T> = {
  readonly [P in keyof T]: T[P] extends object ? DeepReadonly<T[P]> : T[P];
};

/**
 * Any type of asynchronous function
 */

type AnyPromiseFunction<T extends any[] = any[], R = void> = (
  ...arg: T
) => PromiseLike<R>;

/**
 * Any type of common function
 */
type AnyNormalFunction<T extends any[] = any[], R = void> = (...arg: T) => R;

/**
 * Any type of function
 */
type AnyFunction<T extends any[] = any[], R = void> =
  | AnyNormalFunction<T, R>
  | AnyPromiseFunction<T, R>;

/**
 *  T | null 包装
 */
type Nullable<T> = null | T;

/**
 * T | Not null 包装
 */
type NonNullable<T> = T extends null | undefined ? never : T;

/**
 * 字符串类型对象
 */
type Recordable<T> = Record<string, T>;

/**
 * 字符串类型对象（只读）
 */
interface ReadonlyRecordable<T = any> {
  readonly [key: string]: T;
}

/**
 * setTimeout Return value type
 */
type TimeoutHandle = ReturnType<typeof setTimeout>;

/**
 * setInterval Return value type
 */
type IntervalHandle = ReturnType<typeof setInterval>;

/**
 * 也许它是一个计算的 ref，或者一个 getter 函数
 *
 */
type MaybeReadonlyRef<T> = (() => T) | ComputedRef<T>;

/**
 *Maybe it is a Ref, or a normal value, or a getter function
 *
 */
type MaybeComputedRef<T> = MaybeReadonlyRef<T> | MaybeRef<T>;

type Merge<O extends object, T extends object> = {
  [K in keyof O | keyof T]: K extends keyof T
    ? T[K]
    : K extends keyof O
      ? O[K]
      : never;
};

/**
 * T = [
 *  { name: string; age: number; },
 *  { sex: 'male' | 'female'; age: string }
 * ]
 * =>
 * MergeAll<T> = {
 *  name: string;
 *  sex: 'male' | 'female';
 *  age: string
 * }
 */
type MergeAll<
  T extends object[],
  R extends object = Record<string, any>,
> = T extends [infer F extends object, ...infer Rest extends object[]]
  ? MergeAll<Rest, Merge<R, F>>
  : R;

export {
  type AnyFunction,
  type AnyNormalFunction,
  type AnyPromiseFunction,
  type DeepPartial,
  type DeepReadonly,
  type IntervalHandle,
  type MaybeComputedRef,
  type MaybeReadonlyRef,
  type Merge,
  type MergeAll,
  type NonNullable,
  type Nullable,
  type ReadonlyRecordable,
  type Recordable,
  type TimeoutHandle,
};
