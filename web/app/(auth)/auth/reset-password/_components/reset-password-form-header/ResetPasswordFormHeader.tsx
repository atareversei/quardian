import { Stack, Title } from '@mantine/core';
import { l } from '@/languages/language';

export function ResetPasswordFormHeader() {
  return (
    <Stack gap={8}>
      <Title order={1} fw={900}>
        {l.publ.auth.resetPasswordPageTitle}
      </Title>
    </Stack>
  );
}
