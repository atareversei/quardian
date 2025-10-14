import Image from 'next/image';
import Link from 'next/link';
import { IconHome } from '@tabler/icons-react';
import { Button, Stack, Text, Title } from '@mantine/core';
import img from '@/assets/error-images/not-found.webp';
import { l } from '@/languages/language';

export default function NotFound() {
  return (
    <div className="take-all-screen">
      <div className="error-box">
        <figure className="error-box_image">
          <Image src={img} width={400} alt="" />
        </figure>
        <div className="error-box_content">
          <Stack gap={5}>
            <Title order={2}>{l.comn.erro.rootNotFoundTitle}</Title>
            <Text>{l.comn.erro.rootNotFoundDescription}</Text>
          </Stack>
          <div className="error-box_actions">
            <Button leftSection={<IconHome />} component={Link} href="/">
              {l.comn.erro.home}
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
}
