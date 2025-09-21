import { MantineColorScheme } from '@mantine/core';
import { l } from '@/languages/language';
import { IconKey } from '@/utils/icon/icon';

function getNextColorScheme(currentColorScheme: MantineColorScheme): MantineColorScheme {
  switch (currentColorScheme) {
    case 'light':
      return 'dark';
    case 'dark':
      return 'auto';
    case 'auto':
      return 'light';
  }
}

function getSchemeText(colorScheme: MantineColorScheme): string {
  switch (colorScheme) {
    case 'light':
      return l.comn.cmpn.lightTheme;
    case 'dark':
      return l.comn.cmpn.darkTheme;
    case 'auto':
      return l.comn.cmpn.systemDefaultTheme;
    default:
      return l.comn.cmpn.changeColorTheme;
  }
}

function getSchemeChangeText(colorScheme: MantineColorScheme): string {
  switch (colorScheme) {
    case 'light':
      return l.comn.cmpn.changeToDarkTheme;
    case 'dark':
      return l.comn.cmpn.changeToSystemDefaultTheme;
    case 'auto':
      return l.comn.cmpn.changeToLightTheme;
    default:
      return l.comn.cmpn.changeColorTheme;
  }
}

function getSchemeIcon(colorScheme: MantineColorScheme): IconKey {
  switch (colorScheme) {
    case 'light':
      return 'sun';
    case 'dark':
      return 'moon';
    case 'auto':
      return 'brandWindows';
  }
}

export const ColorSchemeToggleUtility = {
  getSchemeChangeText,
  getSchemeText,
  getNextColorScheme,
  getSchemeIcon,
};
