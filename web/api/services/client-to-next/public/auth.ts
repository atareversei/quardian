import { Client } from '@/api/request/client';
import { NextResponseFieldAuthToken } from '@/api/types/next-response';
import { User } from '@/types/entities/user';

export type RegisterRequest = { name: string; email: string; password: string };
export type RegisterResponse = {
  user: User;
  token: string;
};
const register = async (data: RegisterRequest) => {
  const res = await Client.ClientToNext.post<RegisterResponse>('register', data);
  return res;
};

export type LoginRequest = { email: string; password: string };
type LoginResponse = {
  user: User;
  token: string;
};
const login = async (data: LoginRequest) => {
  const res = await Client.ClientToNext.post<LoginResponse>('login', data);
  return res;
};

type GetTokenResponse = NextResponseFieldAuthToken;
const getToken = async () => {
  const res = await Client.ClientToNext.get<GetTokenResponse>('token');
  return res;
};

export const Auth = {
  register,
  login,
  getToken,
};
