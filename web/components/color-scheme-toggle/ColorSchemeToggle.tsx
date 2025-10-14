'use client';

import { ButtonVariant, FloatingPosition, MantineColor, MantineSize } from '@mantine/core';
import { ColorSchemeToggleButton } from './components/color-scheme-toggle-button/ColorSchemeToggleButton';
import { ColorSchemeToggleIcon } from './components/color-scheme-toggle-icon/ColorSchemeToggleIcon';
import { ColorSchemeToggleSelect } from './components/color-scheme-toggle-select/ColorSchemeToggleSelect';

type Props = {
  type: 'button' | 'icon' | 'button-with-icon' | 'select';
  variant?: ButtonVariant;
  color?: MantineColor;
  size?: MantineSize;
  position?: FloatingPosition;
};

export function ColorSchemeToggle({
  type,
  variant = 'light',
  color = 'primary',
  size = 'sm',
  position = 'bottom',
}: Props) {
  if (type === 'icon') {
    return <ColorSchemeToggleIcon variant={variant} color={color} size={size} />;
  }
  if (type === 'button' || type === 'button-with-icon') {
    return (
      <ColorSchemeToggleButton
        variant={variant}
        color={color}
        size={size}
        withIcon={type === 'button-with-icon'}
      />
    );
  }

  if (type === 'select') {
    return (
      <ColorSchemeToggleSelect variant={variant} color={color} size={size} position={position} />
    );
  }
}
