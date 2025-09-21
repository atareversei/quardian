import { l } from '@/languages/language';

/**
 * Represents the return type of a validation function.
 * - `string` indicates a validation error message.
 * - `null` means the value is valid.
 */
type ValidationReturnType = string | null;

const nameMinLength = 2;
const emailRegex = /^\S+@\S+\.\S+$/;
const passwordMinLength = 8;

export function isNotEmpty(value: string): ValidationReturnType {
  if (value.length === 0) {
    return l.comn.form.thisFieldIsRequired;
  }
  return null;
}

export function validateName(value: string, required: boolean = true): ValidationReturnType {
  if (required && isNotEmpty(value)) {
    return isNotEmpty(value);
  }
  if (value.length < nameMinLength) {
    return l.comn.form.nameMustBeAtLeast2Characters;
  }
  return null;
}

export function validateEmail(value: string, required: boolean = true): ValidationReturnType {
  if (required && isNotEmpty(value)) {
    return isNotEmpty(value);
  }
  if (!emailRegex.test(value)) {
    return l.comn.form.emailHasInvalidFormat;
  }
  return null;
}

export function validatePassword(value: string, required: boolean = true): ValidationReturnType {
  if (required && isNotEmpty(value)) {
    return isNotEmpty(value);
  }
  if (value.length < passwordMinLength) {
    return l.comn.form.passwordMustBeAtLeast8Characters;
  }
  return null;
}
