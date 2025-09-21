'use client';

import { AuthenticationProvider } from '@/providers/auth/AuthProvider';

export default function Providers({ children }: any) {
  return <AuthenticationProvider>{children}</AuthenticationProvider>;
}
