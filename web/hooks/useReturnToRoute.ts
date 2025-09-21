'use client';

import { useRouter } from 'next/navigation';
import { Routes } from '@/utils/route/route';
import { useSearchParameters } from './useSearchParameters';

type PriorityKey = 'back' | 'searchParam' | 'url';

type Config = {
  /**
   * Whether the router should navigate to the previous route in history.
   */
  back: boolean;

  /**
   * Whether the router should use the `return_url` search parameter to navigate.
   * @default true
   */
  searchParam?: boolean;

  /**
   * A specific URL to navigate to when other options are not applicable.
   */
  url?: string;

  /**
   * Defines the priority order for determining the route to navigate.
   * Must be a tuple containing all three options in a specific order.
   * @default ['searchParam', 'back', 'url']
   */
  priority?: [PriorityKey, PriorityKey, PriorityKey];
};

/**
 * Navigates the user to a route based on the given configuration.
 * The navigation follows a priority order:
 * 1. If `back` is enabled and prioritized, it navigates to the previous route in history.
 * 2. If `searchParam` is enabled and prioritized, it checks for a `return_url` search parameter and navigates accordingly.
 * 3. If `url` is provided and prioritized, it navigates to the specified URL.
 *
 * @param config - Configuration object specifying navigation behavior.
 * @returns A function that, when invoked, performs the navigation.
 */
export function useReturnToRoute({
  back,
  searchParam = true,
  url = Routes.home,
  priority = ['searchParam', 'back', 'url'],
}: Config) {
  const returnUrl = useSearchParameters('return_url');
  const router = useRouter();
  function returnToRoute() {
    for (let i = 0; i < priority?.length; i++) {
      const p = priority[i];
      if (p === 'searchParam' && searchParam && returnUrl) {
        router.push(returnUrl);
        break;
      } else if (p === 'url' && url) {
        router.push(url);
        break;
      } else if (p === 'back' && back) {
        router.back();
        break;
      }
    }
  }
  return returnToRoute;
}
