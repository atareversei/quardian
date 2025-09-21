export const LOCALHOST = process.env.NEXT_PUBLIC_LOCAL ?? 'http://localhost:3000';
export const NEXT_API_DOMAIN = process.env.NEXT_PUBLIC_NEXT_API_DOMAIN ?? 'http://localhost:3000';
export const NEXT_API_VERSION = process.env.NEXT_PUBLIC_NEXT_API_VERSION ?? '/api';
export const MAIN_API_DOMAIN = process.env.NEXT_PUBLIC_MAIN_API_DOMAIN ?? 'http://localhost:8000';
export const MAIN_API_VERSION = process.env.NEXT_PUBLIC_MAIN_API_VERSION ?? '/api';
export const isDev = process.env.NODE_ENV === 'development';
export const isProd = process.env.NODE_ENV === 'production';
export const isTest = process.env.NODE_ENV === 'test';
