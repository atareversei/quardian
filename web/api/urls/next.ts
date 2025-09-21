export const NextURLs = {
  register: '/public/auth/register',
  login: '/public/auth/login',
  token: '/public/auth/token',
  profile_me: '/user/profile',
};

export type NextURLs = keyof typeof NextURLs;
