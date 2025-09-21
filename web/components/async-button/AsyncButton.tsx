import { Button, ButtonVariant, Loader, MantineColor, MantineSize } from '@mantine/core';

type Props = {
  isSubmitting: boolean;
  fullWidth?: boolean;
  variant?: ButtonVariant;
  color?: MantineColor;
  size?: MantineSize;
  type?: 'button' | 'submit';
  disabled?: boolean;
  clickHandler?: () => void;
  children: React.ReactNode;
};

export function AsyncButton({
  isSubmitting,
  fullWidth = true,
  variant = 'filled',
  color = 'primary',
  size = 'sm',
  type = 'button',
  disabled = false,
  clickHandler,
  children,
}: Props) {
  return (
    <Button
      onClick={clickHandler}
      variant={variant}
      color={color}
      size={size}
      type={type}
      fullWidth={fullWidth}
      disabled={isSubmitting || disabled}
    >
      {!isSubmitting && <span>{children}</span>}
      {isSubmitting && <Loader size={size} />}
    </Button>
  );
}
