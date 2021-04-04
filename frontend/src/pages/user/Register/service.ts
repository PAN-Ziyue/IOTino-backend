import request from 'umi-request';
import type { UserRegisterParams } from './index';


export interface response {
  
}


export async function Register(params: UserRegisterParams) {
  return request('/api/register', {
    method: 'POST',
    data: params,
  });
}