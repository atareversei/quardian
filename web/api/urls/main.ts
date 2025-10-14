export const MainURLs = {
  register: '/public/auth/register',
  login: '/public/auth/login',
  forget_password_request_otp: '/public/auth/forget-password/request-otp',
  forget_password_verify_otp: '/public/auth/forget-password/verify',
  forget_password_update_password: '/public/auth/forget-password/update-password',

  profile_me: '/user/profile/me',

  devices_list: '/user/tickets',
};

export type MainURLs = keyof typeof MainURLs;

