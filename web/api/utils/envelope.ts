import { APIResponse, HeaderItem, HTTPStatus } from '@/api/types/request';
import { RequestHeaderConstants } from '@/api/utils/constants';

export type EnvelopeOptions = {
  headers?: HeaderItem[];
};

function envelopeMainByNext(res: APIResponse<any>, options?: EnvelopeOptions): Response {
  console.log(res.status);
  if (!res.ok) {
    return new Response(
      JSON.stringify({
        ...(res.message && { message: res.message }),
        ...(res.formError && { errors: res.formError }),
        ...(res.meta && { meta: res.meta }),
      }),
      {
        status: res.status,
        headers: [RequestHeaderConstants.ContentTypeJson, ...(options?.headers ?? [])],
      }
    );
  }

  return new Response(
    JSON.stringify({
      ...(res.message && { message: res.message }),
      ...(res.data && { data: res.data }),
      ...(res.paginate && { paginate: res.paginate }),
    }),
    {
      status: HTTPStatus.Success,
      headers: [RequestHeaderConstants.ContentTypeJson, ...(options?.headers ?? [])],
    }
  );
}

function envelopeFailureOfNext(options?: EnvelopeOptions) {
  return new Response(JSON.stringify({ message: 'NextJS server failed' }), {
    status: HTTPStatus.InternalServerError,
    headers: [RequestHeaderConstants.ContentTypeJson, ...(options?.headers ?? [])],
  });
}

export const Envelope = {
  envelopeMainByNext,
  envelopeFailureOfNext,
};
