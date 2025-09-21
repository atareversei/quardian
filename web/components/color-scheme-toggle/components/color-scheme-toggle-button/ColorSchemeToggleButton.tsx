import {
  Button,
  ButtonVariant,
  MantineColor,
  MantineSize,
  useMantineColorScheme,
} from '@mantine/core';
import { i } from '@/utils/icon/icon';
import { ColorSchemeToggleUtility } from '../../utils/color-scheme-toggle-utility';

type Props = {
  variant: ButtonVariant;
  color: MantineColor;
  size: MantineSize;
  withIcon: boolean;
};

export function ColorSchemeToggleButton({ variant, color, size, withIcon }: Props) {
  const { colorScheme, setColorScheme } = useMantineColorScheme();

  function handleColorToggle() {
    setColorScheme(ColorSchemeToggleUtility.getNextColorScheme(colorScheme));
  }

  return (
    <Button
      variant={variant}
      color={color}
      size={size}
      onClick={handleColorToggle}
      leftSection={withIcon ? i(ColorSchemeToggleUtility.getSchemeIcon(colorScheme), size) : null}
    >
      {ColorSchemeToggleUtility.getSchemeChangeText(colorScheme)}
    </Button>
  );
}
