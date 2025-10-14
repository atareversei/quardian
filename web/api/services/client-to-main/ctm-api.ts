import { Auth } from '@/api/services/client-to-main/public/auth';
import { Ticket } from '@/api/services/client-to-main/user/ticket';

export const CTM = {
  auth: Auth,
  ticket: Ticket,
};
