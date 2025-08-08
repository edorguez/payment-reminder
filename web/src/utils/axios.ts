import axios from 'axios';
import { getIdToken } from 'firebase/auth';
import { auth } from '../firebase';

const api = axios.create({
  headers: { 'Content-Type': 'application/json' },
});

api.interceptors.request.use(async (config) => {
  const user = auth.currentUser;
  if (!user) throw new Error('User not logged in');
  const token = await getIdToken(user, false);
  config.headers.Authorization = `Bearer ${token}`;
  return config;
});

api.interceptors.response.use(
  (res) => res,
  (error) => {
    const status  = error.response?.status ?? 0;
    const message = error.response?.data?.error ??
                    error.response?.data?.message ??
                    error.message;
    throw { status, message };
  }
);

export { api };
