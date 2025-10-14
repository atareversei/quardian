import { ScrollArea, Stack } from '@mantine/core';
import c from './auth-container.module.css';

type Props = {
  children: React.ReactNode;
};

export function AuthContainer({ children }: Props) {
  return (
    <ScrollArea scrollbars="y" type="scroll">
      <div className={c.container}>
        <Stack gap={36}>{children}</Stack>
      </div>
    </ScrollArea>
  );
}
