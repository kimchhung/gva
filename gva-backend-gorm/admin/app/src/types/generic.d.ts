declare type StringValueKeys<T> = {
  [P in keyof T]: T[P] extends string ? T[P] : never;
};

// all keys of the above type
declare type Key<T> = keyof StringValueKeys<T>;

type Decrement<N extends number> = [
  -1,
  0,
  1,
  2,
  3,
  4,
  5,
  6,
  7,
  8,
  9,
  10,
  11,
  12,
  13,
  14,
  15,
  16,
  17,
  18,
  19,
  20,
  ...0[],
][N];

type AllKeys<
  T extends Record<string, unknown>,
  Key extends string = keyof T,
  Depth extends number = 2,
> = Key extends string
  ? T[Key] extends Record<string, unknown>
    ? Depth extends 0
      ? `${Key}`
      : `${Key}__${AllKeys<T[Key], keyof T[Key], Decrement<Depth>>}`
    : `${Key}`
  : never;

type PickKey<T extends object, Keys extends AllKeys<T> = AllKeys<T>> = Keys;
