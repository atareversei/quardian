export type TestCase<D, E = string> = {
  goal?: string;
  data: D;
  expected: E;
};
