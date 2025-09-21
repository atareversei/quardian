import { FloatingPosition } from '@mantine/core';
import { IconKey } from '../icon/icon';

export type FloatingPositionIconType = 'chevron' | 'arrow';

export function getIconForFloatingPosition(
  position: FloatingPosition,
  iconType: FloatingPositionIconType = 'chevron'
): IconKey {
  switch (position) {
    case 'top':
    case 'top-start':
    case 'top-end':
      if (iconType === 'arrow') {
        return 'arrowUp';
      }
      return 'chevronUp';
    case 'right':
    case 'right-end':
    case 'right-start':
      if (iconType === 'arrow') {
        return 'arrowRight';
      }
      return 'chevronRight';
    case 'bottom':
    case 'bottom-start':
    case 'bottom-end':
      if (iconType === 'arrow') {
        return 'arrowDown';
      }
      return 'chevronDown';
    case 'left':
    case 'left-end':
    case 'left-start':
      if (iconType === 'arrow') {
        return 'arrowLeft';
      }
      return 'chevronLeft';
  }
}
