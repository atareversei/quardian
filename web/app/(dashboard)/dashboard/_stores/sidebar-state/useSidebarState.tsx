import { create } from 'zustand';

type SidebarStateStoreModel = {
  isExpanded: boolean;
  toggleExpansion: () => void;
};

export const useSidebarState = create<SidebarStateStoreModel>()((set) => ({
  isExpanded: true,
  toggleExpansion: () => set((state) => ({ isExpanded: !state.isExpanded })),
}));
