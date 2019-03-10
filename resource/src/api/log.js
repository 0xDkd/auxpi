import request from '@/utils/request'

export function getLogs(params) {
  return request({
    url: '/admin/get_logs_list',
    method: 'get',
    params
  })
}
