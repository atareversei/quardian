import { TypedQuery } from '@/api/types/api-filters';
import { RequestMeta } from '@/api/types/request';
import { printObject } from '@/test-utils/print';
import { TestCase } from '@/test-utils/types';
import {
  extractIndexedQueryParams,
  generateMetaURL,
  generateQueryParams,
  GenerateQueryParamsFromRecordType,
} from '@/utils/query-param/query-param';

describe('generateQueryParams', () => {
  type GenerateQueryParamsDataType = {
    queryParams: GenerateQueryParamsFromRecordType;
  };

  const testCases: TestCase<GenerateQueryParamsDataType>[] = [
    { data: { queryParams: { name: 'ata' } }, expected: '?name=ata', goal: 'handles basic usage' },
    {
      data: { queryParams: {} },
      expected: '',
      goal: 'handles empty objects',
    },
    {
      data: {
        queryParams: {
          age: '14',
          name: 'ata',
          isStudent: true,
          university: null,
          whatever: undefined,
        },
      },
      expected: '?age=14&name=ata&isStudent=true',
      goal: 'handles multiple fields and it also handles "null" and "undefined"',
    },
    {
      data: {
        queryParams: {
          cities: [
            { id: 1, type: 'export' },
            { id: 2, type: 'import' },
            { id: 3, type: 'import' },
            { id: 3, type: 'export' },
          ],
          name: 'ata',
        },
      },
      expected:
        '?cities[0][id]=1&cities[0][type]=export&cities[1][id]=2&cities[1][type]=import&cities[2][id]=3&cities[2][type]=import&cities[3][id]=3&cities[3][type]=export&name=ata',
      goal: 'handles typed queries',
    },
  ];

  testCases.forEach((tc) => {
    it(`returns '${tc.expected}' for input of ${printObject(tc.data.queryParams)} - goal: ${tc.goal}`, () => {
      expect(generateQueryParams(tc.data.queryParams)).toBe(tc.expected);
    });
  });
});

describe('extractIndexedQueryParams', () => {
  type ExtractIndexedQueryParamsDataType = {
    url: string;
    key: string;
    typed?: boolean;
  };
  type ExtractIndexedQueryParamsExpectedType = TypedQuery[] | number[];

  const testCases: TestCase<
    ExtractIndexedQueryParamsDataType,
    ExtractIndexedQueryParamsExpectedType
  >[] = [
    {
      data: {
        url: '?city[0][id]=1&city[0][type]=export&city[1][id]=2&city[1][type]=import',
        key: 'city',
        typed: true,
      },
      expected: [
        { id: 1, type: 'export' },
        { id: 2, type: 'import' },
      ],
      goal: 'handles simple typed queries',
    },
    {
      data: {
        url: '?city[1][id]=1&city[1][type]=export&city[3][id]=2&city[3][type]=import',
        key: 'city',
        typed: true,
      },
      expected: [
        { id: 1, type: 'export' },
        { id: 2, type: 'import' },
      ],
      goal: 'handles skipped indexes for typed queries',
    },
    {
      data: {
        url: '?user[0]=1&user[1]=2',
        key: 'user',
      },
      expected: [1, 2],
      goal: 'handles basic indexed queries',
    },
    {
      data: {
        url: '?user[3]=1&user[18]=2',
        key: 'user',
      },
      expected: [1, 2],
      goal: 'handles skipped indexes for indexed queries',
    },
  ];

  testCases.forEach((tc) => {
    it(`returns ${tc.expected} for URL of ${tc.data.url} - goal: ${tc.goal}`, () => {
      expect(extractIndexedQueryParams(tc.data.url, tc.data.key, tc.data.typed)).toEqual(
        tc.expected
      );
    });
  });
});

describe('generateMetaURL', () => {
  type GenerateMetaURLDataType = {
    url: string;
    meta?: RequestMeta;
  };
  const testCases: TestCase<GenerateMetaURLDataType>[] = [
    {
      data: {
        url: 'example.com/{user_id}/ticket_id',
      },
      expected: 'example.com/{user_id}/ticket_id',
      goal: 'handles empty meta object',
    },
    {
      data: {
        url: 'example.com/{ticket_id}/ticket_id',
        meta: {
          pathParams: { ticket_id: 28 },
          queryParams: {
            page: 12,
          },
        },
      },
      expected: 'example.com/28/ticket_id?page=12',
      goal: 'handles path params and query params',
    },
  ];

  testCases.forEach((tc) => {
    it(`returns ${tc.expected} for URL of ${tc.data.url} and query paths of : ${printObject(tc.data.meta?.pathParams)} and query params of ${printObject(tc.data.meta?.queryParams)} - goal: ${tc.goal}`, () => {
      expect(generateMetaURL(tc.data.url, tc.data.meta)).toBe(tc.expected);
    });
  });
});
