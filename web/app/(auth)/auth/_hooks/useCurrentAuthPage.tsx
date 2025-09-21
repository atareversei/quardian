'use client';

import { usePathname } from 'next/navigation';
import { Routes } from '@/utils/route/route';

export type AuthPage = 'register' | 'login' | 'reset-password';

export function useCurrentAuthPage() {
  const pathname = usePathname();
  let page: AuthPage | null = null;
  switch (pathname) {
    case Routes.register:
      page = 'register';
      break;
    case Routes.login:
      page = 'login';
      break;
    case Routes.resetPassword:
      page = 'reset-password';
      break;
    default:
      page = 'register';
  }

  return { page };
}
