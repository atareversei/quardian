'use client';

import Image from 'next/image';
import Link from 'next/link';
import { IconHome, IconReload } from '@tabler/icons-react';
import { Button, Stack, Text, Title } from '@mantine/core';
import img from '@/assets/error-images/generic-error.webp';
import { l } from '@/languages/language';
import { NextErrorProps } from '@/types/common/next-utils';

export default function RootError({ error, reset }: NextErrorProps) {
  return (
    <div className="take-all-screen">
      <div className="error-box">
        <figure className="error-box_image">
          <Image src={img} width={400} alt="" />
        </figure>
        <div className="error-box_content">
          <Stack gap={5}>
            <Title order={2}>{l.comn.erro.rootErrorTitle}</Title>
            <Text>{l.comn.erro.rootErrorDescription}</Text>
            <Text>
              {l.comn.erro.causeOfError}: {error.message}
            </Text>
          </Stack>
          <div className="error-box_actions">
            <Button leftSection={<IconHome />} component={Link} href="/">
              {l.comn.erro.home}
            </Button>
            <Button leftSection={<IconReload />} onClick={reset}>
              {l.comn.erro.refresh}
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
}
