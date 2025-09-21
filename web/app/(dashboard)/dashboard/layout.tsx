'use client';

import { Sidebar } from '@/app/(dashboard)/dashboard/_components/sidebar/Sidebar';
import { useSidebarState } from '@/app/(dashboard)/dashboard/_stores/sidebar-state/useSidebarState';
import c from './layout.module.css';

export default function ProfileLayout({ children }: { children: any }) {
  const sidebarState = useSidebarState();

  return (
    <div className={c.container}>
      <div className={c.padding}>
        <div className={c.center} data-expanded={sidebarState.isExpanded}>
          <Sidebar />
          <div className={c.content}>{children}</div>
        </div>
      </div>
    </div>
  );
}
