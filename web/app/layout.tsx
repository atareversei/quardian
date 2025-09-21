import '../global.css';
import '@mantine/core/styles.css';

import React from 'react';
import { ColorSchemeScript, mantineHtmlProps, MantineProvider } from '@mantine/core';
import { l } from '@/languages/language';
import { theme } from '../theme';

import '@mantine/notifications/styles.css';

import { Notifications } from '@mantine/notifications';
import Providers from '@/providers/providers';

export const metadata = {
  title: l.publ.root.pageTitle,
  description: 'I am using Mantine with Next.js!',
};

export default function RootLayout({ children }: { children: any }) {
  return (
    <html lang="fa" dir="rtl" {...mantineHtmlProps}>
      <head>
        <ColorSchemeScript />
        <link rel="shortcut icon" href="/favicon.svg" />
        <meta
          name="viewport"
          content="minimum-scale=1, initial-scale=1, width=device-width, user-scalable=no"
        />
        <link rel="preconnect" href="https://fonts.googleapis.com" />
        <link rel="preconnect" href="https://fonts.gstatic.com" crossOrigin="anonymous" />
        <link
          href="https://fonts.googleapis.com/css2?family=Vazirmatn:wght@100..900&display=swap"
          rel="stylesheet"
        />
      </head>
      <body>
        <MantineProvider theme={theme}>
          <Providers>
            <Notifications zIndex={999} />
            {children}
          </Providers>
        </MantineProvider>
      </body>
    </html>
  );
}
