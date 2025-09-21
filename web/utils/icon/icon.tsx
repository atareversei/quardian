import {
  IconActivityHeartbeat,
  IconAlertHexagon,
  IconArrowDown,
  IconArrowLeft,
  IconArrowRight,
  IconArrowUp,
  IconAsterisk,
  IconAt,
  IconBasketHeart,
  IconBell,
  IconBrandGoogleFilled,
  IconBrandWindows,
  IconCheck,
  IconChevronDown,
  IconChevronLeft,
  IconChevronRight,
  IconChevronUp,
  IconDoorExit,
  IconEyeCog,
  IconFlag,
  IconHistory,
  IconHome,
  IconMapPins,
  IconMessageCircle,
  IconMoon,
  IconPackage,
  IconPointFilled,
  IconSettings,
  IconSun,
  IconTicket,
  IconUser,
  IconX,
} from '@tabler/icons-react';
import { MantineColor, MantineSize } from '@mantine/core';

const icons = {
  activityHeartbeat: IconActivityHeartbeat,
  alertHexagon: IconAlertHexagon,
  arrowDown: IconArrowDown,
  arrowLeft: IconArrowLeft,
  arrowRight: IconArrowRight,
  arrowUp: IconArrowUp,
  asterisk: IconAsterisk,
  at: IconAt,
  basketHeart: IconBasketHeart,
  bell: IconBell,
  brandWindows: IconBrandWindows,
  brandGoogleFilled: IconBrandGoogleFilled,
  check: IconCheck,
  chevronDown: IconChevronDown,
  chevronLeft: IconChevronLeft,
  chevronRight: IconChevronRight,
  chevronUp: IconChevronUp,
  doorExit: IconDoorExit,
  eyeCog: IconEyeCog,
  flag: IconFlag,
  history: IconHistory,
  home: IconHome,
  mapPins: IconMapPins,
  messageCircle: IconMessageCircle,
  moon: IconMoon,
  package: IconPackage,
  pointFilled: IconPointFilled,
  settings: IconSettings,
  sun: IconSun,
  ticket: IconTicket,
  user: IconUser,
  x: IconX,
};

const IconSize = {
  xs: 12,
  sm: 14,
  md: 16,
  lg: 18,
  xl: 22,
};

export type IconKey = keyof typeof icons;

export function i(
  icon: IconKey,
  size: MantineSize = 'sm',
  color: MantineColor = 'currentcolor'
): React.ReactNode {
  const Icon = icons[icon];
  return <Icon size={IconSize[size]} color={color} />;
}

