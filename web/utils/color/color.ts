const colors = {
  dimmed: 'var(--mantine-color-dimmed)',
  greenText: 'var(--mantine-color-green-text)',
  redText: 'var(--mantine-color-red-text)',
};

type ColorKey = keyof typeof colors;

export function getColorCustomVariables(color: ColorKey) {
  return colors[color];
}
