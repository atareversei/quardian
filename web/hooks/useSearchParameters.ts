'use client';

import { useSearchParams } from 'next/navigation';
import { SearchParamKey } from '@/utils/search-param/search-param';

export function useSearchParameters(key: SearchParamKey): string | null {
  return useSearchParams().get(key);
}
