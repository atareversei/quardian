import { Client } from '@/api/request/client';
import { NextResponseFieldAuthToken } from '@/api/types/next-response';
import { User } from '@/types/entities/user';

type GetProfileResponse = User & NextResponseFieldAuthToken;
const getProfile = async () => {
  const res = await Client.ClientToNext.get<GetProfileResponse>('profile_me');
  return res;
};

export const Profile = {
  getProfile,
};
