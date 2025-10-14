import { SidebarButtonWithTooltip } from '@/app/(dashboard)/dashboard/_components/sidebar/components/sidebar-button/SidebarButton';
import { SidebarExpandButtonWithTooltip } from '@/app/(dashboard)/dashboard/_components/sidebar/components/sidebar-button/SidebarExpandButton';
import { useSidebar } from '@/app/(dashboard)/dashboard/_components/sidebar/useSidebar';
import { isDev } from '@/env-configs';
import c from './sidebar.module.css';

type Props = {};

export function Sidebar(props: Props) {
  const { sidebarData } = useSidebar();

  function renderSidebar() {
    return sidebarData.map((data) => {
      if (data.onlyDev && !isDev) return null;

      return (
        <SidebarButtonWithTooltip
          isActive={false}
          icon={data.icon}
          text={data.text}
          tooltipText={data.text}
          sectionType={data.sectionType}
          leftSection={data.leftSection}
          style={{ ...(data.bottomStart && { marginBlockStart: 'auto' }) }}
        />
      );
    });
  }

  return (
    <div className={c.container}>
      <aside className={c.aside}>
        <SidebarExpandButtonWithTooltip />

        {renderSidebar()}
      </aside>
    </div>
  );
}
