import {
  MAIN_API_DOMAIN,
  MAIN_API_VERSION,
  NEXT_API_DOMAIN,
  NEXT_API_VERSION,
} from '@/env-configs';
import { Request } from './request';

const ClientToMain = new Request<'main'>('main', {
  apiPrefix: MAIN_API_DOMAIN + MAIN_API_VERSION,
});
const NextToMain = new Request<'main'>('main', {
  apiPrefix: MAIN_API_DOMAIN + MAIN_API_VERSION,
  duplex: 'half',
});
const ClientToNext = new Request<'next'>('next', {
  apiPrefix: NEXT_API_DOMAIN + NEXT_API_VERSION,
});

export const Client = {
  ClientToMain,
  ClientToNext,
  NextToMain,
};
