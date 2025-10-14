import { ActionIcon } from '@mantine/core';
import { i } from '@/utils/icon/icon';

export function GoogleLogin() {
  return (
    <ActionIcon size="lg" variant="default">
      {i('brandGoogleFilled', 'lg')}
    </ActionIcon>
  );
}
