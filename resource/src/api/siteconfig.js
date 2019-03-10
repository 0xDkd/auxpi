import request from '@/utils/request'

// 首页图床各状况统计
export function getConfig() {
  return request({
    url: '/admin/get_site_config',
    method: 'get'
  })
}
// 获取证书
export function getRsaKey() {
  return request({
    url: '/admin/get_rsa_key',
    method: 'get'
  })
}
// 获取证书
export function update(params) {
  return request({
    url: '/admin/update_site_config',
    method: 'post',
    data: params
  })
}
// 重置证书
export function resetKey() {
  return request({
    url: '/admin/update_site_config',
    method: 'post'
  })
}
