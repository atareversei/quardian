import { Envelope, EnvelopeOptions } from '@/api/utils/envelope';
import { GetRequestWithBodyError } from '@/utils/error/request';

function attachBody(options: RequestInit, body: any): string {
  if (body && options.method === 'GET') throw new GetRequestWithBodyError();

  if (body instanceof FormData) {
    options.body = body;
    return 'multipart/form-data';
  }

  options.body = JSON.stringify(body);
  return 'application/json';
}

function handleNextServerFailure(
  request: Request,
  err: unknown,
  options?: EnvelopeOptions
): Response {
  // TODO: log errors
  console.error(
    `NextJS had an issue while making a request.\n URL: ${request.url} \n Error: ${err}`
  );
  return Envelope.envelopeFailureOfNext(options);
}

export const RequestUtils = {
  attachBody,
  handleNextServerFailure,
};
