import request from '@/utils/request'

// 首页图床各状况统计
export function getStorePercent() {
  return request({
    url: '/admin/get_store_report',
    method: 'get'
  })
}

export function getUserSevenReport() {
  return request({
    url: '/admin/get_user_report',
    method: 'get'
  })
}

export function getApiSevenReport() {
  return request({
    url: '/admin/get_api_report',
    method: 'get'
  })
}

export function getAllImageSevenReport() {
  return request({
    url: '/admin/get_all_images_report',
    method: 'get'
  })
}

export function getLocalImageSevenReport() {
  return request({
    url: '/admin/get_local_images_report',
    method: 'get'
  })
}

export function getAuxpiSystemInfo() {
  return request({
    url: '/admin/get_auxpi_info',
    method: 'get'
  })
}

