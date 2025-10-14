'use client';

import React, { Suspense } from 'react';
import { motion } from 'motion/react';
import { AuthContainer } from './_components/auth-container/AuthContainer';
import { AuthContainerVisuals } from './_components/auth-visuals/AuthVisuals';
import { useAuthBackgroundChange } from './_hooks/useAuthBackgroundChange';
import { useCurrentAuthPage } from './_hooks/useCurrentAuthPage';
import c from './layout.module.css';

type Props = {
  children: React.ReactNode;
};

export default function AuthLayout({ children }: Props) {
  const { page } = useCurrentAuthPage();
  const { bgRef } = useAuthBackgroundChange(page);

  return (
    <div className={c.bg} ref={bgRef}>
      <motion.div
        layout
        className={c.page}
        initial={{ y: 10 }}
        animate={{ y: 0 }}
        transition={{ layout: { duration: 0.05 } }}
      >
        <div className={c.container}>
          <AuthContainer>
            {/* TODO: find a solution for Suspense and Skeleton issue */}
            <Suspense fallback="loading...">{children}</Suspense>
          </AuthContainer>

          <AuthContainerVisuals />
        </div>
      </motion.div>
    </div>
  );
}
