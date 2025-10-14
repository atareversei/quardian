'use client';

import React from 'react';
import { Stack } from '@mantine/core';
import { useSearchParameters } from '@/hooks/useSearchParameters';
import { SearchParamKey } from '@/utils/search-param/search-param';
import { AuthContainerFooter } from '../_components/auth-container-footer/AuthContainerFooter';
import { LoginFormHeader } from './_components/login-form-header/LoginFormHeader';
import { LoginFormNav } from './_components/login-form-nav/LoginFormNav';
import { LoginForm } from './_components/login-form/LoginForm';

export default function LoginPage() {
  const [emailQueryParameter, setEmailQueryParameter] = React.useState<string>(
    useSearchParameters(SearchParamKey.email) ?? ''
  );

  return (
    <Stack gap={36}>
      <LoginFormNav emailQueryParameter={emailQueryParameter} />
      <Stack gap={24}>
        <LoginFormHeader />
        <LoginForm setEmailQueryParameter={setEmailQueryParameter} />
      </Stack>
      <AuthContainerFooter />
    </Stack>
  );
}
