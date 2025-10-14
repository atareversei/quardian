'use client';

import React from 'react';
import { Stack } from '@mantine/core';
import { useSearchParameters } from '@/hooks/useSearchParameters';
import { SearchParamKey } from '@/utils/search-param/search-param';
import { AuthContainerFooter } from '../_components/auth-container-footer/AuthContainerFooter';
import { ResetPasswordFormHeader } from './_components/reset-password-form-header/ResetPasswordFormHeader';
import { ResetPasswordFormNav } from './_components/reset-password-form-nav/ResetPasswordFormNav';

export default function resetPasswordPage() {
  const [emailQueryParameter] = React.useState<string>(
    useSearchParameters(SearchParamKey.email) ?? ''
  );

  return (
    <Stack gap={36}>
      <ResetPasswordFormNav emailQueryParameter={emailQueryParameter} />
      <Stack gap={24}>
        <ResetPasswordFormHeader />
      </Stack>
      <AuthContainerFooter />
    </Stack>
  );
}
