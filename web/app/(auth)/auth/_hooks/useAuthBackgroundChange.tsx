import React from 'react';
import { AuthPage } from './useCurrentAuthPage';

export function useAuthBackgroundChange(currentPage: AuthPage) {
  const bgRef = React.useRef<HTMLDivElement>(null);

  React.useEffect(() => {
    if (bgRef.current) {
      if (currentPage === 'register') {
        bgRef.current.style.setProperty('--auth-gradient-color-from', 'var(--gradient-lime-from)');
        bgRef.current.style.setProperty('--auth-gradient-color-to', 'var(--gradient-lime-to)');
      }
      if (currentPage === 'login') {
        bgRef.current.style.setProperty(
          '--auth-gradient-color-from',
          'var(--gradient-primary-from)'
        );
        bgRef.current.style.setProperty('--auth-gradient-color-to', 'var(--gradient-primary-to)');
      }
      if (currentPage === 'reset-password') {
        bgRef.current.style.setProperty(
          '--auth-gradient-color-from',
          'var(--gradient-purple-from)'
        );
        bgRef.current.style.setProperty('--auth-gradient-color-to', 'var(--gradient-purple-to)');
      }
    }
  }, [currentPage]);

  return { bgRef };
}
