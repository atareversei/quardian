import { TypedQuery } from '@/api/types/api-filters';
import { RequestMeta } from '@/api/types/request';

export type GenerateQueryParamsFromRecordType = Record<
  string,
  string | number | boolean | number[] | TypedQuery[] | null | undefined
>;

export function generateQueryParams(
  queryParams: Record<
    string,
    string | number | boolean | number[] | TypedQuery[] | null | undefined
  >
): string {
  if (!queryParams) {
    return '';
  }
  let queryString = '';
  const entries = Object.entries(queryParams);
  for (let i = 0; i < entries.length; i += 1) {
    const [key, value] = entries[i];

    if (Array.isArray(value) && typeof value[0] === 'number') {
      value.forEach((item, i) => {
        queryString === '' ? (queryString += '?') : (queryString += '&');
        queryString += `${key}[${i}]=${item}`;
      });
    }

    if (Array.isArray(value) && typeof value[0] === 'object') {
      value.forEach((item, i) => {
        queryString === '' ? (queryString += '?') : (queryString += '&');
        queryString += `${key}[${i}][id]=${(item as TypedQuery).id}&${key}[${i}][type]=${(item as TypedQuery).type}`;
      });
    }

    if (typeof value === 'boolean' || typeof value === 'number' || typeof value === 'string') {
      if (value !== null && typeof value !== 'undefined' && value !== '') {
        queryString === '' ? (queryString += '?') : (queryString += '&');
        queryString += `${key}=${value}`;
      }
    }
  }
  return queryString;
}

export function extractIndexedQueryParams(
  url: string,
  key: string,
  typed: boolean = false
): number[] | TypedQuery[] {
  const [, queryStr] = url.split('?');
  if (!queryStr) {
    return [];
  }
  const queries = queryStr.split('&');

  const typedRegexpType = new RegExp(
    `${key}[\\d{1,2}][type]`.replaceAll('[', '\\[').replaceAll(']', '\\]')
  );
  const typedRegexpId = new RegExp(
    `${key}[\\d{1,2}][id]`.replaceAll('[', '\\[').replaceAll(']', '\\]')
  );
  const typedRegexIndexStart = new RegExp('\\[\\d{1,2}\\]');
  const regexp = new RegExp(`${key}[\\d{1,2}]`.replaceAll('[', '\\[').replaceAll(']', '\\]'));

  if (!typed) {
    const extraction: number[] = [];
    queries.forEach((item) => {
      if (regexp.test(item)) {
        const [, idStr] = item.split('=');
        const id = parseInt(idStr, 10);
        if (!Number.isNaN(id)) {
          extraction.push(id);
        }
      }
    });
    return extraction;
  }

  function getIndex(item: string): number {
    const indexStart = item.search(typedRegexIndexStart);
    let indexStr = '';
    for (let i = indexStart + 1; i < item.length; i++) {
      if (item[i] === ']') break;
      indexStr += item[i];
    }
    return parseInt(indexStr, 10);
  }

  const extraction: Partial<TypedQuery>[] = [];
  queries.forEach((item) => {
    if (typedRegexpId.test(item)) {
      const index = getIndex(item);

      const [, idStr] = item.split('=');
      const id = parseInt(idStr, 10);

      if (extraction[index]?.type) extraction[index].id = id;
      else extraction[index] = { id };
    }

    if (typedRegexpType.test(item)) {
      const index = getIndex(item);
      const [, type] = item.split('=');

      if (extraction[index]?.id) {
        extraction[index].type = type;
      } else {
        extraction[index] = { type };
      }
    }
  });

  return extraction.filter(
    (extracted) => extracted && extracted.id && extracted.type
  ) as TypedQuery[];
}

export function generateMetaURL(url: string, meta?: RequestMeta): string {
  if (!meta) {
    return url;
  }
  let richURL = url;
  if (meta.pathParams) {
    const entries = Object.entries(meta.pathParams);
    for (let i = 0; i < entries.length; i += 1) {
      const [key, value] = entries[i];
      richURL = richURL.replaceAll(`{${key}}`, value.toString());
    }
  }
  if (meta.queryParams) {
    richURL += generateQueryParams(meta.queryParams);
  }
  return richURL;
}
