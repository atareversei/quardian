import { RequestHeaderConstants } from '@/api/utils/constants';
import { isBrowser } from '@/utils/environement/environment';
import { parseJSONResponse } from '@/utils/json/json';
import { generateMetaURL } from '@/utils/query-param/query-param';
import { APIResponse, Body, GeneralRequestConfig, RequestConfig } from '../types/request';
import { getURL, NameSpaceURLs, UrlNamespace } from '../urls/urls';
import { intercept } from '../utils/interceptors';
import { RequestUtils } from '../utils/request-utils';
import { StandardizedResponse } from '../utils/standardized-response';

export class Request<NS extends UrlNamespace> {
  private _generalConfig: GeneralRequestConfig;
  private _namespace: NS;

  constructor(namespace: NS, config: GeneralRequestConfig) {
    this._generalConfig = config;
    this._namespace = namespace;
  }

  public async request<T>(
    url: NameSpaceURLs<typeof this._namespace>,
    config: RequestConfig,
    body?: Body
  ): Promise<APIResponse<T>> {
    try {
      const options: RequestInit = {
        method: config.method,
      };

      const meta = config.meta;
      const contentType = RequestUtils.attachBody(options, body);
      options.headers = [
        ...(this._generalConfig?.headers ?? []),
        ...(config.headers ?? []),
        ['Content-Type', contentType],
        RequestHeaderConstants.AcceptJson,
        RequestHeaderConstants.AcceptLanguageFa,
      ];

      const res = await fetch(
        generateMetaURL(
          getURL(config.apiPrefix ?? this._generalConfig.apiPrefix, this._namespace, url),
          meta
        ),
        options
      );
      const json = await parseJSONResponse(res);
      console.log(json);
      if (!res.ok) {
        const stdRes = StandardizedResponse.error(res.status, json);
        if (isBrowser) intercept(stdRes, res.status, []);
        return stdRes;
      }
      return StandardizedResponse.success<T>(json);
    } catch (err) {
      const stdFail = StandardizedResponse.failure(err);
      if (isBrowser) intercept(stdFail, stdFail.status, []);
      return stdFail;
    }
  }

  public async get<T>(
    url: NameSpaceURLs<typeof this._namespace>,
    config?: Partial<RequestConfig>
  ): Promise<APIResponse<T>> {
    return this.request<T>(url, { ...config, method: 'GET' });
  }

  public async post<T>(
    url: NameSpaceURLs<typeof this._namespace>,
    body: Body,
    config?: Partial<RequestConfig>
  ): Promise<APIResponse<T>> {
    return this.request<T>(url, { ...config, method: 'POST' }, body);
  }

  public async put<T>(
    url: NameSpaceURLs<typeof this._namespace>,
    body: Body,
    config?: Partial<RequestConfig>
  ): Promise<APIResponse<T>> {
    return this.request<T>(url, { ...config, method: 'PUT' }, body);
  }

  public async patch<T>(
    url: NameSpaceURLs<typeof this._namespace>,
    body: Body,
    config?: Partial<RequestConfig>
  ): Promise<APIResponse<T>> {
    return this.request<T>(url, { ...config, method: 'PATCH' }, body);
  }

  public async delete<T>(
    url: NameSpaceURLs<typeof this._namespace>,
    body: Body,
    config?: Partial<RequestConfig>
  ): Promise<APIResponse<T>> {
    return this.request<T>(url, { ...config, method: 'DELETE' }, body);
  }
}
