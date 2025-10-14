'use client';

import React from 'react';
import { Divider, Group, PasswordInput, Stack, TextInput } from '@mantine/core';
import { FormButton } from '@/components/form-button/FormButton';
import { Statements } from '@/components/statements/Statements';
import { l } from '@/languages/language';
import { i } from '@/utils/icon/icon';
import { PrivacyAndRules } from '../../../_components/privacy-and-rules/PrivacyAndRules';
import { GoogleSignup } from './components/google-signup/GoogleSignup';
import { useRegisterForm } from './hooks/useRegisterForm';
import { useRegisterFormPasswordStatements } from './hooks/useRegisterFormPasswordStatements';

const icons = {
  name: i('user'),
  email: i('at'),
  password: i('asterisk'),
};

type Props = {
  setEmailQueryParameter: React.Dispatch<React.SetStateAction<string>>;
};

export function RegisterForm({ setEmailQueryParameter }: Props) {
  const { form, handleSubmit, isSubmitting } = useRegisterForm();
  const { statements } = useRegisterFormPasswordStatements(form);

  return (
    <form onSubmit={form.onSubmit(handleSubmit)}>
      <Stack>
        <TextInput
          withAsterisk
          label={l.publ.auth.name}
          description={l.publ.auth.nameDescription}
          size="sm"
          leftSection={icons.name}
          autoFocus
          autoComplete="name"
          {...form.getInputProps('name')}
        />
        <TextInput
          withAsterisk
          type="email"
          label={l.publ.auth.email}
          size="sm"
          className="ltr-input english-font"
          leftSection={icons.email}
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
        <Statements statements={statements} />
        <Divider />
        <Stack gap={16}>
          <PrivacyAndRules />
          <Group wrap="nowrap">
            <FormButton isSubmitting={isSubmitting}>{l.publ.auth.registerSubmitButton}</FormButton>
            <Divider orientation="vertical" />
            <GoogleSignup />
          </Group>
        </Stack>
      </Stack>
    </form>
  );
}
