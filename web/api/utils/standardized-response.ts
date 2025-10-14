import { isDev } from '@/env-configs';
import { l } from '@/languages/language';
import { isJSONParsingError } from '@/utils/error/json';
import { isGetRequestWithBodyError } from '@/utils/error/request';
import { ErrorResponse, HTTPStatus, SuccessResponse } from '../types/request';

function badRequest(json: any): ErrorResponse {
  return { ok: false, status: HTTPStatus.BadRequest, message: json.message };
}

function unauthorized(): ErrorResponse {
  return { ok: false, status: HTTPStatus.Unauthorized };
}

function forbidden(): ErrorResponse {
  return { ok: false, status: HTTPStatus.Forbidden };
}

function conflict(json: any): ErrorResponse {
  return { ok: false, status: HTTPStatus.Conflict, meta: json.meta };
}

function unprocessableContent(json: any): ErrorResponse {
  const fe: Record<string, string[]> = {};
  if (json.errors) {
    const entries = Object.entries(json.errors);
    for (let i = 0; i < entries.length; i += 1) {
      const [key, value] = entries[i];
      if (typeof value === 'string') {
        fe[key] = [value];
      }
      if (Array.isArray(value)) {
        fe[key] = value;
      }
    }
  }

  return {
    ok: false,
    status: HTTPStatus.UnprocessableContent,
    message: json.message ?? l.comn.form.formHasIssues,
    formError: fe,
  };
}

function internal(): ErrorResponse {
  return { ok: false, status: HTTPStatus.InternalServerError };
}

function unhandled(): ErrorResponse {
  console.error('UNHANDLED ERROR');
  return {
    ok: false,
    status: HTTPStatus.UnhandledError,
    message: 'This error has not been handled in api layer.',
  };
}

function error(status: number, json: any): ErrorResponse {
  switch (status) {
    case HTTPStatus.BadRequest:
      return badRequest(json);

    case HTTPStatus.Unauthorized:
      return unauthorized();

    case HTTPStatus.Forbidden:
      return forbidden();

    case HTTPStatus.Conflict:
      return conflict(json);

    case HTTPStatus.UnprocessableContent:
      return unprocessableContent(json);

    case HTTPStatus.InternalServerError:
      return internal();

    default:
      return unhandled();
  }
}

function failure(err: unknown): ErrorResponse {
  if (isDev) {
    console.error(err);
  }

  if (isJSONParsingError(err)) {
    return { ok: false, status: HTTPStatus.JSONParsingError, message: err.message };
  }

  if (isGetRequestWithBodyError(err)) {
    return { ok: false, status: HTTPStatus.GetRequestWithBody, message: err.message };
  }

  return { ok: false, status: HTTPStatus.NetworkError };
}

function success<T>(json: any): SuccessResponse<T> {
  return {
    ok: true,
    status: HTTPStatus.Success,
    message: json.message,
    data: json.data as T,
    paginate: json.paginate ?? null,
  };
}

export const StandardizedResponse = {
  success,
  error,
  failure,
};
