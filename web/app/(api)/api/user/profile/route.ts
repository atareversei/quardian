import { cookies } from 'next/headers';
import { API } from '@/api/api';
import { HTTPStatus } from '@/api/types/request';
import { Envelope } from '@/api/utils/envelope';
import { RequestUtils } from '@/api/utils/request-utils';
import { CookieKey } from '@/utils/cookie/cookie';

export async function GET(request: Request) {
  const cookieStore = await cookies();
  const authToken = cookieStore.get(CookieKey.token);
  const isAuthorized = authToken && authToken.value;

  if (!isAuthorized) {
    return Envelope.envelopeMainByNext({ ok: false, status: HTTPStatus.Unauthorized });
  }

  try {
    const res = await API.ntm.profile.getProfile({ token: authToken?.value });
    if (!res.ok) {
      return Envelope.envelopeMainByNext(res);
    }
    return Envelope.envelopeMainByNext({
      ...res,
      data: { ...res.data, __ntc_token: authToken.value },
    });
  } catch (err) {
    return RequestUtils.handleNextServerFailure(request, err);
  }
}
