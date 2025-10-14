import { Client } from '@/api/request/client';
import { User } from '@/types/entities/user';

type RegisterRequest = { name: string; email: string; password: string };
type RegisterResponse = {
  user: User;
  token: string;
};
const register = async (data: RegisterRequest) => {
  const res = await Client.NextToMain.post<RegisterResponse>('register', data);
  return res;
};

export type LoginRequest = { email: string; password: string };
type LoginResponse = {
  user: User;
  token: string;
};
const login = async (data: LoginRequest) => {
  const res = await Client.NextToMain.post<LoginResponse>('login', data);
  return res;
};

export const Auth = {
  register,
  login,
};
