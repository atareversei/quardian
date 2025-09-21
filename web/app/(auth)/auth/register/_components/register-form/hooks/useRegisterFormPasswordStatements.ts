'use client';

import React from 'react';
import { UseFormReturnType } from '@mantine/form';
import { ValidationStatement, ValidationStatementStatus } from '@/components/statements/Statements';
import { l } from '@/languages/language';
import { RegisterFormFields } from './useRegisterForm';

const passwordLengthStatementText = l.comn.form.passwordMustBeAtLeast8Characters;

const initialStatements: ValidationStatement[] = [
  { text: passwordLengthStatementText, passed: 'neutral' },
];

export function useRegisterFormPasswordStatements(form: UseFormReturnType<RegisterFormFields>) {
  const [statements, setStatements] = React.useState<ValidationStatement[]>(
    () => initialStatements
  );

  form.watch('password', ({ value }) => {
    const lengthStatement = passwordLengthStatement(
      value,
      true,
      statements.find((item) => item.text === passwordLengthStatementText)?.passed ?? 'neutral'
    );
    setStatements([lengthStatement]);
  });

  return { statements };
}

function passwordLengthStatement(
  password: string,
  // TODO: find a better solution -> aggressive and prevState are connected to each other.
  aggressive: boolean = false,
  prevState?: ValidationStatementStatus
): ValidationStatement {
  let isMoreThan8Chars = false;
  const statement: ValidationStatement = {
    text: passwordLengthStatementText,
    passed: 'neutral',
  };

  if (password.length >= 8) {
    isMoreThan8Chars = true;
  }

  if (!isMoreThan8Chars && aggressive && prevState !== 'neutral') {
    statement.passed = 'failed';
  } else if (!isMoreThan8Chars && !aggressive) {
    statement.passed = 'neutral';
  } else if (isMoreThan8Chars) {
    statement.passed = 'passed';
  }

  if (password.length === 0) {
    statement.passed = 'neutral';
  }

  return statement;
}
