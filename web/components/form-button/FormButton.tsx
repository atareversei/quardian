import { ButtonVariant, MantineColor, MantineSize } from '@mantine/core';
import { AsyncButton } from '../async-button/AsyncButton';

type Props = {
  isSubmitting: boolean;
  fullWidth?: boolean;
  variant?: ButtonVariant;
  color?: MantineColor;
  size?: MantineSize;
  disabled?: boolean;
  children: React.ReactNode;
};

export function FormButton({
  isSubmitting,
  fullWidth = true,
  variant = 'filled',
  color = 'primary',
  size = 'sm',
  disabled = false,
  children,
}: Props) {
  return (
    <AsyncButton
      isSubmitting={isSubmitting}
      variant={variant}
      color={color}
      size={size}
      type="submit"
      fullWidth={fullWidth}
      disabled={isSubmitting || disabled}
    >
      {children}
    </AsyncButton>
  );
}
