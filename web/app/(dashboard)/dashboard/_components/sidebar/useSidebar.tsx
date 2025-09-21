import React from 'react';
import { Badge } from '@mantine/core';
import { l } from '@/languages/language';
import { IconKey } from '@/utils/icon/icon';

type SidebarDataItem = {
  icon: IconKey;
  text: string;
  onlyDev?: boolean;
  leftSection?: React.ReactNode;
  sectionType?: 'notification' | 'new' | 'expected' | 'none';
  onClick?: () => void;
  bottomStart?: boolean;
};

export function useSidebar() {
  const sidebarData: SidebarDataItem[] = [
    {
      icon: 'home',
      text: l.dash.main.summarySidebar,
    },
    {
      icon: 'package',
      text: l.dash.main.orderSidebar,
      sectionType: 'notification',
      leftSection: (
        <Badge color="dark" circle>
          12
        </Badge>
      ),
      onlyDev: true,
    },
    {
      icon: 'basketHeart',
      text: l.dash.main.favoriteSidebar,
      onlyDev: true,
    },
    {
      icon: 'messageCircle',
      text: l.dash.main.commentSidebar,
      sectionType: 'expected',
      leftSection: <Badge color="yellow">{l.dash.main.expectedSidebar}</Badge>,
      onlyDev: true,
    },
    {
      icon: 'mapPins',
      text: l.dash.main.addressSidebar,
      onlyDev: true,
    },
    {
      icon: 'history',
      text: l.dash.main.historySidebar,
      onlyDev: true,
    },
    {
      icon: 'bell',
      text: l.dash.main.notificationSidebar,
      sectionType: 'notification',
      leftSection: (
        <Badge color="dark" circle>
          5
        </Badge>
      ),
      onlyDev: true,
    },
    {
      icon: 'ticket',
      text: l.dash.main.ticketSidebar,
      bottomStart: true,
    },
    {
      icon: 'user',
      text: l.dash.main.profileSidebar,
    },
    {
      icon: 'doorExit',
      text: l.dash.main.logoutSidebar,
    },
  ];

  return { sidebarData };
}
