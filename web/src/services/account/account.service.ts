// services/account/account.service.ts
import { api } from '../../utils/axios';
import { toResult } from '../../utils/api-result';
import { URLS } from '../../utils/service-url';
import type { UserDto } from './account.types';

const accountApi = {
  get: <T>(path: string, params?: URLSearchParams) =>
    api.get<T>(`${URLS.account}${path}${params ? `?${params}` : ''}`),

  post: <T>(path: string, data?: unknown) =>
    api.post<T>(`${URLS.account}${path}`, data),

  put: <T>(path: string, data?: unknown) =>
    api.put<T>(`${URLS.account}${path}`, data),

  delete: <T>(path: string) =>
    api.delete<T>(`${URLS.account}${path}`),
};

export const accountService = {
  getUser: (params: { email?: string; firebaseId?: string }) => {
    const sp = new URLSearchParams();
    if (params.email) sp.set('email', params.email);
    if (params.firebaseId) sp.set('firebaseId', params.firebaseId);
    return toResult(accountApi.get<UserDto | null>('/users', sp));
  },

  createUser: () => toResult(accountApi.post<UserDto>('/users')),

  updateUser: (id: string) => toResult(accountApi.put<UserDto>(`/users/${id}`)),

  deleteUser: (id: string) => toResult(accountApi.delete<void>(`/users/${id}`)),
} as const;
