import { cookies } from 'next/headers';
import { HTTPStatus } from '@/api/types/request';
import { Envelope } from '@/api/utils/envelope';
import { CookieKey } from '@/utils/cookie/cookie';

export async function GET() {
  const cookieStore = await cookies();
  const authToken = cookieStore.get(CookieKey.token);
  const isAuthorized = authToken && authToken.value;

  if (!isAuthorized) {
    return Envelope.envelopeMainByNext({ ok: false, status: HTTPStatus.Unauthorized });
  }

  return Envelope.envelopeMainByNext({
    ok: true,
    status: HTTPStatus.Success,
    data: { __ntc_token: authToken?.value },
  });
}
