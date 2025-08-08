import type { AxiosResponse } from 'axios';
import type { ApiResult } from '../types/http';

export async function toResult<T>(
  promise: Promise<AxiosResponse<T>>
): Promise<ApiResult<T>> {
  try {
    const { data } = await promise;
    return { ok: true, data };
  } catch (err) {
    const { status, message } = err as { status: number; message: string };
    return { ok: false, status, message };
  }
}
