import { ColorSchemeToggle } from '@/components/color-scheme-toggle/ColorSchemeToggle';

export function AuthContainerFooter() {
  return (
    <footer>
      <ColorSchemeToggle type="select" size="md" variant="subtle" position="top" />
    </footer>
  );
}
