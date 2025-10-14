'use client';

export const SearchParamKey = {
  return_url: 'return_url',
  email: 'email',
} as const;

export type SearchParamKey = keyof typeof SearchParamKey;
