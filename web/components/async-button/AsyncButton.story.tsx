import { Meta, StoryObj } from '@storybook/react';
import { fn } from '@storybook/test';
import { l } from '@/languages/language';
import { AsyncButton } from './AsyncButton';

const meta: Meta<typeof AsyncButton> = {
  title: 'Async Button',
  component: AsyncButton,
  tags: ['autodocs'],
  args: {
    clickHandler: fn(),
  },
};

export default meta;

type Story = StoryObj<typeof AsyncButton>;

export const States: Story = {
  args: {
    fullWidth: false,
    children: l.publ.auth.loginSubmitButton,
  },
};
