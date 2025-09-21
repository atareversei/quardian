'use client';

import { createTheme, MantineColorsTuple, Notification, virtualColor } from '@mantine/core';
import notificationClasses from './styles/lib/mantine-notification.module.css';

const vitaliBlue: MantineColorsTuple = [
  '#e5f3ff',
  '#cde2ff',
  '#9ac2ff',
  '#64a0ff',
  '#3884fe',
  '#1d72fe',
  '#0969ff',
  '#0058e4',
  '#004ecd',
  '#0043b5',
];

export const theme = createTheme({
  fontFamily: '"IBM Plex Sans Arabic", sans-serif',
  headings: { fontFamily: '"IBM Plex Sans Arabic", sans-serif' },
  colors: {
    vitaliBlue,
    primary: virtualColor({
      name: 'primary',
      light: 'vitaliBlue',
      dark: 'blue',
    }),
  },
  components: {
    Notification: Notification.extend({
      defaultProps: {
        radius: 'lg',
        withBorder: true,
        classNames: notificationClasses,
      },
    }),
  },
  primaryColor: 'primary',
  radius: { xs: '0.175rem', sm: '0.3rem', md: '0.6rem', lg: '1.2rem', xl: '2.4rem' },
  defaultRadius: 'md',
});

