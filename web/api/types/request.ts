import { PathParams, QueryParams } from './api-filters';

export enum HTTPStatus {
  NetworkError = 1,
  CORS = 2,
  JSONParsingError = 3,
  GetRequestWithBody = 4,
  UnhandledError = 99,
  Success = 200,
  BadRequest = 400,
  Unauthorized = 401,
  Forbidden = 403,
  NotFound = 404,
  Conflict = 409,
  UnprocessableContent = 422,
  InternalServerError = 500,
}

export type Method = 'GET' | 'POST' | 'PUT' | 'PATCH' | 'DELETE';
export type RequestMeta = {
  queryParams?: QueryParams;
  pathParams?: PathParams;
};
export type GenerateRequestMeta<Q extends keyof QueryParams, P extends keyof PathParams> = {
  queryParams: Required<Pick<QueryParams, Q>>;
  pathParams: Required<Pick<PathParams, P>>;
};
export type GenerateRequestQueryMeta<Q extends keyof QueryParams> = {
  queryParams: Required<Pick<QueryParams, Q>>;
};
export type GenerateRequestPathMeta<P extends keyof PathParams> = {
  pathParams: Required<Pick<PathParams, P>>;
};
export type Body = FormData | Record<string, any>;
export type FormError = { [key: string]: string[] };

type Pagination = { current_page: number; last_page: number; total: number };

export type HeaderItem = [string, string];

export type RequestConfig = {
  method: Method;
  apiPrefix?: string;
  meta?: RequestMeta;
  headers?: HeaderItem[];
  duplex?: 'half' | 'full';
};
export type GeneralRequestConfig = Omit<RequestConfig, 'method' | 'meta' | 'apiPrefix'> & {
  apiPrefix: string;
};

export type SuccessResponse<T> = {
  ok: true;
  status: HTTPStatus.Success;
  data: T;
  message?: string;
  paginate?: Pagination;
};
export type ErrorResponse = {
  ok: false;
  status: HTTPStatus;
  message?: string;
  formError?: FormError;
  meta?: any;
};
export type APIResponse<T> = SuccessResponse<T> | ErrorResponse;
