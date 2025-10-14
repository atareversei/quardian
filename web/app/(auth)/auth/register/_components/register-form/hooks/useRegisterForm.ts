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
import {
  validateEmail,
  validateName,
  validatePassword,
} from '@/utils/form-validation/form-validation';
import { SearchParamKey } from '@/utils/search-param/search-param';

export type RegisterFormFields = {
  name: string;
  email: string;
  password: string;
};

export function useRegisterForm() {
  const [isSubmitting, setIsSubmitting] = React.useState(false);
  const returnTo = useReturnToRoute({ back: false });
  const { setToken, setUser } = useContext(AuthenticationContext);

  const emailSearchParamValue = useSearchParameters(SearchParamKey.email);

  const form = useForm<RegisterFormFields>({
    mode: 'uncontrolled',
    initialValues: {
      name: '',
      email: emailSearchParamValue ?? '',
      password: '',
    },
    validate: {
      name: (value) => validateName(value),
      email: (value) => validateEmail(value),
      password: (value) => validatePassword(value),
    },
  });

  const handleSubmit = async (values: typeof form.values) => {
    setIsSubmitting(true);
    const res = await API.ctn.auth.register(values);
    handleResponseNotifications(res, form, l.publ.auth.yourAccountWasCreatedNotificationTitle, {
      [HTTPStatus.Success]: l.publ.auth.yourAccountWasCreatedNotification,
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
