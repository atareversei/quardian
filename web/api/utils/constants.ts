import { HeaderItem } from '@/api/types/request';

export const RequestHeaderConstants = {
  ContentTypeJson: ['Content-Type', 'application/json'],
  AcceptJson: ['Accept', 'application/json'],
  AcceptLanguageFa: ['Accept-Language', 'fa'],
} satisfies { [key: string]: HeaderItem };
