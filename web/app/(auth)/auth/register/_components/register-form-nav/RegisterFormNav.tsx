'use client';

import Link from 'next/link';
import { Button, Group, Text } from '@mantine/core';
import { useReturnToRoute } from '@/hooks/useReturnToRoute';
import { l } from '@/languages/language';
import { i } from '@/utils/icon/icon';
import { generateQueryParams } from '@/utils/query-param/query-param';
import { Routes } from '@/utils/route/route';
import { SearchParamKey } from '@/utils/search-param/search-param';

type Props = {
  emailQueryParameter: string;
};

export function RegisterFormNav({ emailQueryParameter }: Props) {
  const returnToRoute = useReturnToRoute({ back: true });
  const queryParameters = generateQueryParams({ [SearchParamKey.email]: emailQueryParameter });

  return (
    <header>
      <nav>
        <Group justify="space-between">
          <Button
            leftSection={i('arrowRight')}
            variant="subtle"
            size="sm"
            type="button"
            onClick={returnToRoute}
          >
            {l.publ.auth.return}
          </Button>

          <Group gap={0}>
            <Text size="sm">{l.publ.auth.haveAccount}</Text>
            <Button
              component={Link}
              href={Routes.login + queryParameters}
              variant="subtle"
              size="compact-sm"
              type="button"
            >
              {l.publ.auth.loginToAccount}
            </Button>
          </Group>
        </Group>
      </nav>
    </header>
  );
}
