import { Client } from '@/api/request/client';
import { User } from '@/types/entities/user';

type GetProfileRequest = { token: string };
type GetProfileResponse = User;
const getProfile = async (data: GetProfileRequest) => {
  const res = await Client.NextToMain.get<GetProfileResponse>('profile_me', {
    headers: [['Authorization', `bearer ${data.token}`]],
  });
  return res;
};

export const Profile = {
  getProfile,
};
