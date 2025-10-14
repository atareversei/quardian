import Link from 'next/link';
import { Anchor, Text } from '@mantine/core';
import { l } from '@/languages/language';

export function PrivacyAndRules() {
  return (
    <Text c="dimmed" size="xs">
      {l.publ.auth.privacyAndRulesAgreementPartOne}{' '}
      <Anchor component={Link} href="#">
        {l.publ.auth.privacyAndRulesAgreementPartTwoLink}
      </Anchor>{' '}
      {l.publ.auth.privacyAndRulesAgreementPartThree}{' '}
      <Anchor component={Link} href="#">
        {l.publ.auth.privacyAndRulesAgreementPartFourLink}
      </Anchor>{' '}
      {l.publ.auth.privacyAndRulesAgreementPartFive}
    </Text>
  );
}
