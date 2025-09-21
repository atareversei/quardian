'use client';

import { isDev } from '@/env-configs';
import { l } from '@/languages/language';
import { Notifications } from '@/utils/notify/notifications';
import { Interceptor } from '../types/call-api';
import { ErrorResponse, HTTPStatus } from '../types/request';

export function intercept(res: ErrorResponse, r: HTTPStatus, off: Interceptor[]) {
  if (r === HTTPStatus.NetworkError && !off?.includes(HTTPStatus.NetworkError)) {
    return handleNetworkError(res);
  }
  if (r === HTTPStatus.Unauthorized && !off?.includes(HTTPStatus.Unauthorized)) {
    return handleUnauthorized(res);
  }
  if (r === HTTPStatus.Forbidden && !off?.includes(HTTPStatus.Forbidden)) {
    return handleForbidden(res);
  }
  if (r === HTTPStatus.UnprocessableContent && !off?.includes(HTTPStatus.UnprocessableContent)) {
    return handleUnprocessableContent(res);
  }
  if (r === HTTPStatus.Conflict && !off?.includes(HTTPStatus.Conflict)) {
    return handleConflict(res);
  }
  if (r === HTTPStatus.InternalServerError && !off?.includes(HTTPStatus.InternalServerError)) {
    return handleInternalServerError(res);
  }

  if (isDev && res.status < HTTPStatus.Success) {
    handleFailure(res);
  }

  return res;
}

function handleNetworkError(res: ErrorResponse) {
  Notifications.error(l.comn.http.networkErrorTitle, l.comn.http.networkError);
  return res;
}

async function handleUnauthorized(res: ErrorResponse) {
  const timeToSeeTheNotification = 6_000;

  Notifications.error(l.comn.http.unauthorizedTitle, l.comn.http.unauthorized, {
    autoClose: timeToSeeTheNotification,
  });

  // await callAPI('logout');

  // setTimeout(() => {
  //   localStorage.removeItem(AuthenticationTokenKey);
  //   localStorage.removeItem(CustomerKey);
  //   window.location.reload();
  // }, timeToSeeTheNotification);

  return res;
}

function handleForbidden(res: ErrorResponse) {
  const timeToSeeTheNotification = 4_000;

  Notifications.warning(l.comn.http.forbiddenTitle, l.comn.http.forbidden, {
    autoClose: timeToSeeTheNotification,
  });
  setTimeout(() => {
    // TODO - change to use routes
    if (window.location.pathname === '/dashboard') {
      window.location.href = '/';
    } else {
      window.location.href = '/dashboard';
    }
  }, timeToSeeTheNotification);
  return res;
}

function handleConflict(res: ErrorResponse) {
  return res;
}

function handleUnprocessableContent(res: ErrorResponse) {
  return res;
}

function handleInternalServerError(res: ErrorResponse) {
  Notifications.error(l.comn.http.internalServerErrorTitle, l.comn.http.internalServerError);
  return res;
}

function handleFailure(res: ErrorResponse) {
  Notifications.error('Development Error', res.message ?? 'Something went wrong while developing', {
    style: { direction: 'ltr' },
    autoClose: false,
  });
}
