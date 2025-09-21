export const CookieKey = {
  token: 'token',
} as const;

export type CookieKey = keyof typeof CookieKey;
