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
// 更新菜单
export function updateMenu(params) {
  return request({
    url: '/admin/update_menu',
    method: 'post',
    data: params
  })
}
// 获取激活中的图床
export function enableStores() {
  return request({
    url: '/admin/get_enableStores',
    method: 'get'
  })
}
// 获取关未激活的图床
export function disableStores() {
  return request({
    url: '/admin/get_enableStores',
    method: 'get'
  })
}
// 获取各种配置的信息
export function getStoreOption(params) {
  return request({
    url: '/options/stores',
    method: 'get',
    params
  })
}
// 更新图床账号信息
export function updateStoreAccount(params, suffix) {
  console.log(params)

  return request({
    url: '/admin/update_stores_options/' + suffix,
    method: 'post',
    data: params
  })
}
// 更新图床状态
export function updateStoreStatus(params) {
  return request({
    url: '/admin/update_store/',
    method: 'post',
    data: params
  })
}
// 通过 Options 接口获取
export function getOptions(params) {
  return request({
    url: '/options/info',
    method: 'get',
    params
  })
}
// 通过 Options 更新数据
export function updateOptions(params, suffix) {
  return request({
    url: '/options/update?key=' + suffix.key + '&group=' + suffix.group,
    method: 'post',
    data: params
  })
}
