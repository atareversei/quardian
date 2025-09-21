'use client';

import React from 'react';
import { API } from '@/api/api';
import { User } from '@/types/entities/user';
import { SessionMeta } from './meta';

type SessionInfo = {
  token: string | null;
  user: User | null;
  meta: SessionMeta | null;
};

type AuthenticationContext = {
  sessionInfo: SessionInfo;
  isInitializing: boolean;
  isSessionInfoLoading: boolean;
  updateSessionInfo: () => void;
  setToken: (token: string) => void;
  setUser: (user: User) => void;
  logout: (returnTo: string | null) => void;
};

type AuthenticationProvider = {
  children: React.ReactNode;
};

const initialSessionInfo = { user: null, token: null, meta: null };
const initialIsInitializing = true;
const initialIsSessionInfoLoading = false;
const initialUpdateSessionInfo = () => {};
const initialSetToken = (token: string) => {
  token;
};
const initialSetUser = (user: User) => {
  user;
};
const initialLogout = (returnTo: string | null) => {
  return returnTo as unknown as Promise<void>;
};

export const AuthenticationContext = React.createContext<AuthenticationContext>({
  sessionInfo: initialSessionInfo,
  isInitializing: initialIsInitializing,
  isSessionInfoLoading: initialIsSessionInfoLoading,
  updateSessionInfo: initialUpdateSessionInfo,
  setToken: initialSetToken,
  setUser: initialSetUser,
  logout: initialLogout,
});

export const AuthenticationProvider = ({ children }: AuthenticationProvider) => {
  const [sessionInfo, setSessionInfo] = React.useState<SessionInfo>(initialSessionInfo);
  const [isInitializing, setIsInitializing] = React.useState(initialIsInitializing);
  const [isSessionInfoLoading, setIsSessionInfoLoading] = React.useState(
    initialIsSessionInfoLoading
  );
  const [shouldUpdate, setShouldUpdate] = React.useState(false);

  React.useEffect(() => {
    async function initialize() {
      setIsInitializing(true);
      await getTokenAndProfile();
      setIsInitializing(false);
    }

    initialize();
  }, [shouldUpdate]);

  function setToken(token: string): void {
    setSessionInfo((prevSession) => ({
      ...(prevSession && { ...prevSession }),
      token,
    }));
  }

  function setUser(user: User): void {
    setSessionInfo((prevSession) => ({
      ...(prevSession && { ...prevSession }),
      user,
    }));
  }

  function updateSessionInfo() {
    setShouldUpdate((prev) => !prev);
  }

  async function getTokenAndProfile() {
    setIsSessionInfoLoading(true);
    const res = await API.ctn.profile.getProfile();
    if (res.ok) {
      setSessionInfo({
        token: res.data.__ntc_token,
        user: res.data,
        meta: {},
      });
    }
    setIsSessionInfoLoading(false);
  }

  async function logout(returnTo: string | null): Promise<void> {
    // TODO: implement
  }

  return (
    <AuthenticationContext.Provider
      value={{
        sessionInfo: sessionInfo as SessionInfo,
        isInitializing,
        isSessionInfoLoading,
        updateSessionInfo,
        setToken,
        setUser,
        logout,
      }}
    >
      {children}
    </AuthenticationContext.Provider>
  );
};
