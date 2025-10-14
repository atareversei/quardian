import {
  ActionIcon,
  ButtonVariant,
  MantineColor,
  MantineSize,
  Tooltip,
  useMantineColorScheme,
} from '@mantine/core';
import { l } from '@/languages/language';
import { i } from '@/utils/icon/icon';
import { ColorSchemeToggleUtility } from '../../utils/color-scheme-toggle-utility';

type Props = {
  variant: ButtonVariant;
  color: MantineColor;
  size: MantineSize;
};

export function ColorSchemeToggleIcon({ variant, color, size }: Props) {
  const { colorScheme, setColorScheme } = useMantineColorScheme();

  function handleColorToggle() {
    setColorScheme(ColorSchemeToggleUtility.getNextColorScheme(colorScheme));
  }

  return (
    <Tooltip label={ColorSchemeToggleUtility.getSchemeChangeText(colorScheme)}>
      <ActionIcon
        variant={variant}
        color={color}
        size={size}
        aria-label={l.comn.cmpn.changeColorTheme}
        onClick={handleColorToggle}
      >
        {i(ColorSchemeToggleUtility.getSchemeIcon(colorScheme), size)}
      </ActionIcon>
    </Tooltip>
  );
}
