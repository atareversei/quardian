import React from 'react';
import { Text, Tooltip, UnstyledButton } from '@mantine/core';
import { useSidebarState } from '@/app/(dashboard)/dashboard/_stores/sidebar-state/useSidebarState';
import { l } from '@/languages/language';
import { i } from '@/utils/icon/icon';
import c from './sidebar-button.module.css';

export const SidebarExpandButton = React.forwardRef<HTMLButtonElement, any>((props, ref) => {
  const sidebarState = useSidebarState();

  return (
    <UnstyledButton
      ref={ref}
      className={c.button}
      data-expanded={sidebarState.isExpanded}
      onClick={() => {
        sidebarState.toggleExpansion();
      }}
    >
      <div data-expanded={sidebarState.isExpanded} className={c.iconContainer}>
        {i('chevronLeft', 'xl', 'white')}
      </div>
      {sidebarState.isExpanded && (
        <Text c="white" className={c.text} lh={1} fw={700} size="sm">
          {l.dash.main.shrinkSidebar}
        </Text>
      )}
    </UnstyledButton>
  );
});

export function SidebarExpandButtonWithTooltip() {
  const sidebarState = useSidebarState();

  return (
    <Tooltip
      label={l.dash.main.expandSidebar}
      position="left"
      disabled={sidebarState.isExpanded}
      transitionProps={{ duration: 100 }}
    >
      <SidebarExpandButton />
    </Tooltip>
  );
}
