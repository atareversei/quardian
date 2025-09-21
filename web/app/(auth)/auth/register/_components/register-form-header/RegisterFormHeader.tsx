import { Stack, Text, Title } from '@mantine/core';
import { l } from '@/languages/language';

export function RegisterFormHeader() {
  return (
    <Stack gap={8}>
      <Title order={1} fw={900}>
        {l.publ.auth.registerPageTitle}
      </Title>
      <Text size="md" c="dimmed">
        {l.publ.auth.registerPageSubTitle}
      </Text>
    </Stack>
  );
}
