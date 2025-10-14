'use client';

import React from 'react';
import { Stack } from '@mantine/core';
import { useSearchParameters } from '@/hooks/useSearchParameters';
import { l } from '@/languages/language';
import { Notifications } from '@/utils/notify/notifications';
import { SearchParamKey } from '@/utils/search-param/search-param';
import { AuthContainerFooter } from '../_components/auth-container-footer/AuthContainerFooter';
import { RegisterFormHeader } from './_components/register-form-header/RegisterFormHeader';
import { RegisterFormNav } from './_components/register-form-nav/RegisterFormNav';
import { RegisterForm } from './_components/register-form/RegisterForm';

export default function RegisterPage() {
  const [emailQueryParameter, setEmailQueryParameter] = React.useState<string>(
    useSearchParameters(SearchParamKey.email) ?? ''
  );

  Notifications.success(
    l.publ.auth.yourAccountWasCreatedNotificationTitle,
    l.publ.auth.yourAccountWasCreatedNotification
  );

  return (
    <Stack gap={36}>
      <RegisterFormNav emailQueryParameter={emailQueryParameter} />
      <Stack gap={24}>
        <RegisterFormHeader />
        <RegisterForm setEmailQueryParameter={setEmailQueryParameter} />
      </Stack>
      <AuthContainerFooter />
    </Stack>
  );
}
