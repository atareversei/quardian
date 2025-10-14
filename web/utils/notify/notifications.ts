'use client';

import { NotificationData, notifications } from '@mantine/notifications';
import { l } from '@/languages/language';

type Options = Omit<NotificationData, 'title' | 'message'>;

function success(title: string, message: string, options?: Options) {
  notifications.show({
    title,
    message,
    color: 'green',
    ...options,
  });
}

function information(title: string, message: string, options?: Options) {
  notifications.show({
    title,
    message,
    ...options,
  });
}

function warning(title: string, message: string, options?: Options) {
  notifications.show({
    title,
    message,
    color: 'orange',
    autoClose: 8_000,
    ...options,
  });
}

function error(title: string, message: string, options?: Options) {
  notifications.show({
    title,
    message,
    color: 'red',
    autoClose: 8_000,
    ...options,
  });
}

function formError(title: string, numberOfErroredFields: number, options?: Options) {
  error(
    title,
    l.comn.form.formHasIssues.replace('{number}', numberOfErroredFields.toString()),
    options
  );
}

export const Notifications = {
  success,
  information,
  warning,
  error,
  formError,
};
