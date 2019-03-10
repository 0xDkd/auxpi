import request from '@/utils/request'

export function getList(params) {
  return request({
    url: 'test/table/list',
    method: 'get',
    params
  })
}
