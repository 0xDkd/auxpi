import request from '@/utils/request'

export function getDispatchList(params) {
  return request({
    url: '/admin/get_dispatch_list',
    method: 'get',
    params
  })
}
