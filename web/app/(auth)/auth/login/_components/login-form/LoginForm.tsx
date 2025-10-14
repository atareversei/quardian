'use client';

import Link from 'next/link';
import { Anchor, Divider, Group, PasswordInput, Stack, Text, TextInput } from '@mantine/core';
import { FormButton } from '@/components/form-button/FormButton';
import { l } from '@/languages/language';
import { i } from '@/utils/icon/icon';
import { Routes } from '@/utils/route/route';
import { PrivacyAndRules } from '../../../_components/privacy-and-rules/PrivacyAndRules';
import { GoogleLogin } from './components/google-signup/GoogleSignup';
import { useLoginForm } from './hooks/useLoginForm';

const icons = {
  email: i('at'),
  password: i('asterisk'),
};

type Props = {
  setEmailQueryParameter: React.Dispatch<React.SetStateAction<string>>;
};

export function LoginForm({ setEmailQueryParameter }: Props) {
  const { form, handleSubmit, isSubmitting } = useLoginForm();

  return (
    <form onSubmit={form.onSubmit(handleSubmit)}>
      <Stack>
        <TextInput
          withAsterisk
          type="email"
          label={l.publ.auth.email}
          size="sm"
          className="ltr-input english-font"
          leftSection={icons.email}
          autoFocus
          autoComplete="email"
          {...form.getInputProps('email')}
          onChange={(e) => {
            setEmailQueryParameter(e.target.value);
            form.getInputProps('email').onChange(e);
          }}
        />
        <PasswordInput
          withAsterisk
          label={l.publ.auth.password}
          size="sm"
          leftSection={icons.password}
          autoComplete="off"
          {...form.getInputProps('password')}
        />
        <Text size="xs" c="dimmed">
          {l.publ.auth.haveForgottenYourPassword}{' '}
          <Anchor component={Link} href={Routes.resetPassword}>
            {l.publ.auth.resetPassword}
          </Anchor>
        </Text>
        <Divider />
        <Stack gap={16}>
          <PrivacyAndRules />
          <Group wrap="nowrap">
            <FormButton isSubmitting={isSubmitting}>{l.publ.auth.loginSubmitButton}</FormButton>
            <Divider orientation="vertical" />
            <GoogleLogin />
          </Group>
        </Stack>
      </Stack>
    </form>
  );
}
