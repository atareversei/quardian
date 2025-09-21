export const GetRequestWithBodyErrorName = 'GetRequestWithBodyError';

export class GetRequestWithBodyError extends Error {
  constructor() {
    super("can't set a body on a 'GET' request.");
    this.name = GetRequestWithBodyErrorName;
  }
}

export function isGetRequestWithBodyError(error: unknown): error is GetRequestWithBodyError {
  return error instanceof GetRequestWithBodyError && error.name === GetRequestWithBodyErrorName;
}
