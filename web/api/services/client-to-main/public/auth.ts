import { User } from '@/types/entities/user';
import { Client } from '../../../request/client';

type ForgetPasswordRequestOTPRequest = { email: string; password: string };
type ForgetPasswordRequestOTPResponse = {
  user: User;
  token: string;
};
const forgetPasswordRequestOTP = async (data: ForgetPasswordRequestOTPRequest) => {
  const res = await Client.ClientToMain.post<ForgetPasswordRequestOTPResponse>(
    'forget_password_request_otp',
    data
  );
  return res;
};

type ForgetPasswordVerifyOTPRequest = { email: string; password: string };
type ForgetPasswordVerifyOTPResponse = {
  user: User;
  token: string;
};
const forgetPasswordVerifyOTP = async (data: ForgetPasswordVerifyOTPRequest) => {
  const res = await Client.ClientToMain.post<ForgetPasswordVerifyOTPResponse>(
    'forget_password_verify_otp',
    data
  );
  return res;
};

type ForgetPasswordUpdatePasswordRequest = { email: string; password: string };
type ForgetPasswordUpdatePasswordResponse = {
  user: User;
  token: string;
};
const forgetPasswordUpdatePassword = async (data: ForgetPasswordUpdatePasswordRequest) => {
  const res = await Client.ClientToMain.post<ForgetPasswordUpdatePasswordResponse>(
    'forget_password_update_password',
    data
  );
  return res;
};

export const Auth = {
  forgetPasswordRequestOTP,
  forgetPasswordVerifyOTP,
  forgetPasswordUpdatePassword,
};
