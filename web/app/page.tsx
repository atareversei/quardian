'use client';

import React from 'react';
import { AuthenticationContext } from '@/providers/auth/AuthProvider';
import { extractIndexedQueryParams } from '@/utils/query-param/query-param';

export default function HomePage() {
  const { sessionInfo } = React.useContext(AuthenticationContext);

  return (
    <div>
      {sessionInfo.token && <p>token: {sessionInfo.token}</p>}
      {sessionInfo.user && (
        <p>
          user name: {sessionInfo.user.name} email: {sessionInfo.user.email}
        </p>
      )}
    </div>
  );
}
