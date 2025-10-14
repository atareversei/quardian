import { Client } from '@/api/request/client';
import { GenerateRequestQueryMeta } from '@/api/types/request';

type TicketPriority = 'low' | 'medium' | 'high';

type ListTicketsRequestMeta = GenerateRequestQueryMeta<'page' | 'per_page'>;
type ListTicketsResponse = {};
const listTicket = async (meta: ListTicketsRequestMeta) => {
  const res = await Client.ClientToMain.get<ListTicketsResponse>('tickets_list', { meta });
  return res;
};

type StoreTicketsRequest = { title: string; priority: TicketPriority };
type StoreTicketsResponse = {};
const storeTicket = async (data: StoreTicketsRequest) => {
  const res = await Client.ClientToMain.post<StoreTicketsResponse>('tickets_list', data);
  return res;
};

export const Ticket = { listTicket, storeTicket };
