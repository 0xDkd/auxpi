import request from '@/utils/request'
import { stringify } from 'qs'
export function login(username, password) {
  return request({
    url: '/auth/login',
    method: 'post',
    data: stringify({
      username,
      password
    })
  })
}

export function getInfo(token) {
  return request({
    url: '/auth/info',
    method: 'get',
    params: { token }
  })
}

export function logout() {
  return request({
    url: '/auth/logout',
    method: 'post'
  })
}
