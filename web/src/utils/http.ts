import { getIdToken } from 'firebase/auth';
import { auth } from '../firebase';

export async function http<T>(
  path: string,
  init?: RequestInit
): Promise<T> {
  const user = auth.currentUser;
  if (!user) throw new Error('User not logged in');

  const token = await getIdToken(user, /* forceRefresh */ false);

  const res = await fetch(path, {
    ...init,
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
      ...(init?.headers ?? {}),
    },
  });

  if (!res.ok) {
    const text = await res.text();
    throw new Error(text || res.statusText);
  }

  const text = await res.text();
  return text ? (JSON.parse(text) as T) : (undefined as T);
}
