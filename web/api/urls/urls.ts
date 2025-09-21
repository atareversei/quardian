import { MainURLs } from './main';
import { NextURLs } from './next';

const URLs = {
  main: {
    ...MainURLs,
  },
  next: {
    ...NextURLs,
  },
} as const;

export type UrlNamespace = keyof typeof URLs;
export type NameSpaceURLs<T extends UrlNamespace> = keyof (typeof URLs)[T];

export function getURL<T extends UrlNamespace>(
  base: string,
  namespace: T,
  name: NameSpaceURLs<T>
): string {
  return base + URLs[namespace][name];
}
