import { FloatingPosition } from '@mantine/core';
import { IconKey } from '../icon/icon';
import { FloatingPositionIconType, getIconForFloatingPosition } from './floating-position';

describe('getIconForFloatingPosition', () => {
  const testCases: {
    position: FloatingPosition;
    iconType?: FloatingPositionIconType;
    expected: IconKey;
  }[] = [
    { position: 'top', expected: 'chevronUp' },
    { position: 'top-start', iconType: 'arrow', expected: 'arrowUp' },
    { position: 'top-end', expected: 'chevronUp' },
    { position: 'right', expected: 'chevronRight' },
    { position: 'bottom', iconType: 'arrow', expected: 'arrowDown' },
    { position: 'bottom-start', iconType: 'arrow', expected: 'arrowDown' },
    { position: 'left', expected: 'chevronLeft' },
    { position: 'left', iconType: 'arrow', expected: 'arrowLeft' },
  ];

  testCases.forEach(({ position, iconType, expected }) => {
    it(`returns ${expected} for position ${position} and icon type "${iconType ?? 'chevron'}"`, () => {
      expect(getIconForFloatingPosition(position, iconType)).toBe(expected);
    });
  });
});
