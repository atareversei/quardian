import { API } from '@/api/api';
import { RegisterRequest } from '@/api/services/client-to-next/public/auth';
import { Envelope } from '@/api/utils/envelope';
import { RequestUtils } from '@/api/utils/request-utils';

export async function POST(request: Request) {
  try {
    const req = (await request.json()) as RegisterRequest;
    const res = await API.ntm.auth.register(req);
    if (!res.ok) {
      return Envelope.envelopeMainByNext(res);
    }
    return Envelope.envelopeMainByNext(res, {
      headers: [
        [
          'Set-Cookie',
          `token=${res.data.token}; Path=/; Secure; HttpOnly; SameSite=Strict; max-age=31536000`,
        ],
      ],
    });
  } catch (err) {
    return RequestUtils.handleNextServerFailure(request, err);
  }
}
