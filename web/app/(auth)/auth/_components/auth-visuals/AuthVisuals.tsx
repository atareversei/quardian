'use client';

import React from 'react';
import { motion } from 'motion/react';
import c from './auth-visuals.module.css';

export function AuthContainerVisuals() {
  return (
    <div className={c.container}>
      <div className={c.visuals}>
        <motion.div
          className={c.top}
          animate={{
            scale: [1, 1.05, 1.05, 1],
            rotate: [6, 12, 8, 6],
            origin: 10,
          }}
          transition={{
            duration: 8 * 3,
            ease: 'easeInOut',
            repeat: Infinity,
            repeatDelay: 1,
          }}
        />

        <motion.div
          className={c.middle}
          animate={{
            rotate: [-2, -4, -4, -2],
          }}
          transition={{
            duration: 18 * 3,
            ease: 'easeInOut',
            repeat: Infinity,
            repeatDelay: 1,
          }}
        />

        <motion.div
          className={c.bottom}
          animate={{
            rotate: [4, 12, 8, 4],
            scale: [1, 1.05, 1.05, 1],
            origin: -10,
          }}
          transition={{
            duration: 8 * 3,
            ease: 'easeInOut',
            repeat: Infinity,
            repeatDelay: 1,
          }}
        />
      </div>
    </div>
  );
}
