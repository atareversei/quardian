'use client';

import { UseFormReturnType } from '@mantine/form';
import { APIResponse, HTTPStatus } from '@/api/types/request';
import { Notifications } from '@/utils/notify/notifications';

type Messages = Partial<Record<HTTPStatus, string>>;

export function handleResponseNotifications(
  res: APIResponse<any>,
  form: UseFormReturnType<any>,
  title: string,
  messages?: Messages
): void {
  const apiMessage = res.message;
  const customMessage = (messages && messages[res.status]) ?? '';
  const message = apiMessage || customMessage;
  const showMessage = message !== '';

  if (!showMessage) return;

  if (res.ok) {
    Notifications.success(title, message);
  } else if (res.status === HTTPStatus.UnprocessableContent && res.formError) {
    form.setErrors(res.formError);
    const numberOfFields = Object.entries(res.formError).length;
    Notifications.formError(title, numberOfFields);
  } else {
    Notifications.error(title, message);
  }
}
