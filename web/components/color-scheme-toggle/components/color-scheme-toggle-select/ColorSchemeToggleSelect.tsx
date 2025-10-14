import {
  Button,
  ButtonVariant,
  FloatingPosition,
  Group,
  MantineColor,
  MantineSize,
  Popover,
  UnstyledButton,
  useMantineColorScheme,
} from '@mantine/core';
import { getIconForFloatingPosition } from '@/utils/floating-position/floating-position';
import { i } from '@/utils/icon/icon';
import { ColorSchemeToggleUtility } from '../../utils/color-scheme-toggle-utility';
import c from './color-scheme-toggle-select.module.css';

type Props = {
  variant: ButtonVariant;
  color: MantineColor;
  size: MantineSize;
  position: FloatingPosition;
};

export function ColorSchemeToggleSelect({ variant, size, color, position }: Props) {
  const { colorScheme, setColorScheme } = useMantineColorScheme();

  return (
    <Popover position={position} withArrow shadow="md">
      <Popover.Target>
        <Button
          variant={variant}
          color={color}
          size={size}
          leftSection={i(ColorSchemeToggleUtility.getSchemeIcon(colorScheme), size)}
          rightSection={i(getIconForFloatingPosition(position), size)}
        >
          {ColorSchemeToggleUtility.getSchemeText(colorScheme)}
        </Button>
      </Popover.Target>
      <Popover.Dropdown>
        <Group>
          <UnstyledButton
            className={c.light}
            onClick={() => {
              setColorScheme('light');
            }}
          >
            <Group gap={4} align="center">
              {i('sun', 'lg')} {ColorSchemeToggleUtility.getSchemeText('light')}
            </Group>
          </UnstyledButton>
          <UnstyledButton
            className={c.dark}
            onClick={() => {
              setColorScheme('dark');
            }}
          >
            <Group gap={4} align="center">
              {i('moon', 'lg')} {ColorSchemeToggleUtility.getSchemeText('dark')}
            </Group>
          </UnstyledButton>
          <UnstyledButton
            className={c.auto}
            onClick={() => {
              setColorScheme('auto');
            }}
          >
            <Group gap={4} align="center">
              {i('brandWindows', 'lg')} {ColorSchemeToggleUtility.getSchemeText('auto')}
            </Group>
          </UnstyledButton>
        </Group>
      </Popover.Dropdown>
    </Popover>
  );
}
