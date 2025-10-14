'use client';

import React, { useContext } from 'react';
import { useForm } from '@mantine/form';
import { API } from '@/api/api';
import { HTTPStatus } from '@/api/types/request';
import { handleResponseNotifications } from '@/api/utils/response-notification';
import { useReturnToRoute } from '@/hooks/useReturnToRoute';
import { useSearchParameters } from '@/hooks/useSearchParameters';
import { l } from '@/languages/language';
import { AuthenticationContext } from '@/providers/auth/AuthProvider';
import { validateEmail, validatePassword } from '@/utils/form-validation/form-validation';
import { SearchParamKey } from '@/utils/search-param/search-param';

export type LoginFormFields = {
  email: string;
  password: string;
};

export function useLoginForm() {
  const [isSubmitting, setIsSubmitting] = React.useState(false);
  const emailSearchParamValue = useSearchParameters(SearchParamKey.email);
  const { setToken, setUser } = useContext(AuthenticationContext);
  const returnTo = useReturnToRoute({ back: false });

  const form = useForm<LoginFormFields>({
    mode: 'uncontrolled',
    initialValues: {
      email: emailSearchParamValue ?? '',
      password: '',
    },
    validate: {
      email: (value) => validateEmail(value),
      password: (value) => validatePassword(value),
    },
  });

  const handleSubmit = async (values: typeof form.values) => {
    setIsSubmitting(true);
    const res = await API.ctn.auth.login(values);
    handleResponseNotifications(res, form, l.publ.auth.youHaveLoggedInNotificationTitle, {
      [HTTPStatus.Success]: l.publ.auth.youHaveLoggedInNotification,
    });
    if (res.ok) {
      setToken(res.data.token);
      setUser(res.data.user);
      returnTo();
    }
    setIsSubmitting(false);
  };

  return { form, handleSubmit, isSubmitting };
}
