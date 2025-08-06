import { http } from '../../utils/http';
import type { UserDto } from './account.types';

const accountUrl: string = import.meta.env.VITE_API_ACCOUNT_URL;

export const accountService = {
  getUser: (params: { email?: string; firebaseId?: string }): Promise<UserDto> => {
    const sp = new URLSearchParams();
    if (params.email) sp.set('email', params.email);
    if (params.firebaseId) sp.set('firebaseId', params.firebaseId);
    return http(`${accountUrl}/users?${sp.toString()}`);
  },

  createUser: (): Promise<UserDto> =>
    http(accountUrl, { method: 'POST' }),

  updateUser: (id: string): Promise<UserDto> =>
    http(`${accountUrl}/users/${id}`, { method: 'PUT' }),

  deleteUser: (id: string): Promise<void> =>
    http(`${accountUrl}/users/${id}`, { method: 'DELETE' }),
} as const;
