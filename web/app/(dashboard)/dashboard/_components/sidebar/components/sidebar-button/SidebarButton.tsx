import React from 'react';
import { MantineStyleProp, Text, Tooltip, UnstyledButton } from '@mantine/core';
import { useSidebarState } from '@/app/(dashboard)/dashboard/_stores/sidebar-state/useSidebarState';
import { i, IconKey } from '@/utils/icon/icon';
import c from './sidebar-button.module.css';

type Props = {
  isActive: boolean;
  icon: IconKey;
  text: string;
  leftSection?: React.ReactNode;
  sectionType?: 'notification' | 'new' | 'expected' | 'none';
  onClick?: () => void;
  style?: MantineStyleProp;
};

export const SidebarButton = React.forwardRef<HTMLButtonElement, Props>(
  ({ isActive, icon, text, leftSection, sectionType, onClick, style }, ref) => {
    const sidebarState = useSidebarState();

    return (
      <UnstyledButton
        ref={ref}
        className={`${c.button} ${isActive ? c.active : ''}`}
        data-expanded={sidebarState.isExpanded}
        onClick={onClick}
        style={style}
      >
        {sectionType === 'notification' && <div className={c.notification}>{leftSection}</div>}
        <div className={c.iconContainer}>{i(icon, 'xl', 'white')}</div>
        {sidebarState.isExpanded && (
          <div className={c.content}>
            <Text c="white" className={c.text} lh={1} fw={700} size="sm">
              {text}
            </Text>

            {leftSection && <div className={c.leftSection}>{leftSection}</div>}
          </div>
        )}
      </UnstyledButton>
    );
  }
);

export function SidebarButtonWithTooltip(props: Props & { tooltipText: string }) {
  const sidebarState = useSidebarState();

  return (
    <Tooltip label={props.tooltipText} position="left" disabled={sidebarState.isExpanded}>
      <SidebarButton {...props} />
    </Tooltip>
  );
}
