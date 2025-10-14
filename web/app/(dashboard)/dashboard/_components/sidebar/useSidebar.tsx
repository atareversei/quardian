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
      text: l.dash.main.dashboardSidebar,
    },
    {
      icon: 'activityHeartbeat',
      text: l.dash.main.liveTrafficSidebar,
      sectionType: 'notification',
    },
    {
      icon: 'package',
      text: l.dash.main.flowsSidebar,
    },
    {
      icon: 'alertHexagon',
      text: l.dash.main.alertsSidebar,
      sectionType: 'notification',
      leftSection: (
        <Badge color="dark" circle>
          5
        </Badge>
      ),
    },
    {
      icon: 'flag',
      text: l.dash.main.captureFiltersSidebar,
      sectionType: 'expected',
      leftSection: <Badge color="yellow">{l.dash.main.expectedSidebar}</Badge>,
    },
    {
      icon: 'history',
      text: l.dash.main.historicalDataSidebar,
    },
    {
      icon: 'eyeCog',
      text: l.dash.main.machineLearningSidebar,
    },
    {
      icon: 'settings',
      text: l.dash.main.settingsSidebar,
      bottomStart: true,
    },
    {
      icon: 'doorExit',
      text: l.dash.main.logoutSidebar,
    },
  ];

  return { sidebarData };
}

