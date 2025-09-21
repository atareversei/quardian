import { Client } from '@/api/request/client';
import { GenerateRequestQueryMeta } from '@/api/types/request';

type ListDevicesRequestMeta = GenerateRequestQueryMeta<'page' | 'per_page'>;
type ListDevicesResponse = {};
const listDevices = async (meta: ListDevicesRequestMeta) => {
  const res = await Client.ClientToMain.get<ListDevicesResponse>('devices_list', { meta });
  return res;
};

export const Ticket = { listTicket: listDevices };

