import request from '@/utils/request'

export function getUser(params) {
  return request({
    url: '/admin/get_users_list',
    method: 'get',
    params
  })
}

export function getUserImages(uid, params) {
  return request({
    url: '/users/' + uid + '/images',
    method: 'get',
    params
  })
}

export function getUserInfo(uid) {
  return request({
    url: '/users/' + uid + '/info',
    method: 'get'
  })
}

export function DeleteUser(uid) {
  return request({
    url: '/admin/delete_user',
    method: 'post',
    data: { id: uid }
  })
}

