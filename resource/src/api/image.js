import request from '@/utils/request'

export function getImageList(params) {
  return request({
    url: '/admin/get_images_list',
    method: 'get',
    params
  })
}

export function delImage(id) {
  return request({
    url: '/admin/del_images',
    method: 'post',
    data: {
      id: id
    }
  })
}

export function getStoreList() {
  return request({
    url: '/admin/get_store_list',
    method: 'get'
  })
}

export function syncImage(params) {
  return request({
    url: '/admin/sync_images',
    method: 'post',
    data: {
      list: params
    }
  })
}

export function getSyncImages(params) {
  return request({
    url: '/admin/get_sync_images',
    method: 'get',
    params
  })
}

export function deleteSyncImageList(id) {
  return request({
    url: '/admin/del_sync_images',
    method: 'post',
    data: {
      id: id
    }
  })
}
