import { http, HttpResponse } from 'msw';
import { setupServer } from 'msw/node';
import { RegisterResponse } from './api/services/client-to-next/public/auth';
import { getURL } from './api/urls/urls';
import { NEXT_API_DOMAIN } from './env-configs';

export const handlers = [
  http.get(getURL(NEXT_API_DOMAIN, 'next', 'register'), () => {
    return HttpResponse.json({
      user: {
        id: 1,
        name: 'ata',
        email: 'ata@gmail.com',
        email_verified_at: null,
      },
      token: 'fake-token',
    } satisfies RegisterResponse);
  }),
];

export const server = setupServer(...handlers);
