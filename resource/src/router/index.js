import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '../views/layout/Layout'

export const constantRouterMap = [
  // 登录页面
  { path: '/login', component: () => import('@/views/login/index'), hidden: true },
  // 404页面
  { path: '/404', component: () => import('@/views/404'), hidden: true },
  // Layout 页面
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    name: 'Dashboard',
    children: [{
      path: 'dashboard',
      meta: { title: '首页', icon: 'dashboard' },
      component: () => import('@/views/dashboard/index')
    }]
  },
  // 用户管理
  {
    path: '/admin/users',
    redirect: '/index',
    component: Layout,
    children: [{
      path: 'index',
      name: 'users',
      component: () => import('@/views/user/index'),
      meta: { title: '用户管理', icon: 'users' }
    }, {
      path: ':id(\\d+)',
      name: 'usersInfoView',
      component: () => import('@/views/user/info'),
      hidden: true
    },
    {
      path: ':id(\\d+)/edit',
      name: 'usersInfoEdit',
      component: () => import('@/views/card/index'),
      hidden: true
    }]
  },
  // 图片管理
  {
    path: '/images',
    component: Layout,
    redirect: '/images/management/',
    meta: { title: '图片管理', icon: 'manage' },
    children: [
      {
        path: 'management/',
        name: 'Images',
        component: () => import('@/views/card/index'),
        meta: { title: '所有图片', icon: 'images' }
      },
      {
        path: 'management/sync',
        name: 'SyncImages',
        component: () => import('@/views/card/sync/index'),
        meta: { title: '同步图片管理', icon: 'sync' }
      }
    ]
  },
  {
    path: '/logs',
    component: Layout,
    name: 'Logs',
    children: [{
      path: 'index',
      name: 'Logs',
      meta: { title: '日志记录', icon: 'logs' },
      component: () => import('@/views/logs/index')
    }]
  },
  {
    path: '/config',
    component: Layout,
    name: 'config',
    children: [{
      path: 'index',
      name: 'Config',
      meta: { title: '站点设置', icon: 'config' },
      component: () => import('@/views/siteconfig/index')
    }]
  },
  // 重定向
  { path: '*', redirect: '/404', hidden: true }
  // 测试路由 是否会在 Menu 上面自动显示

]

export default new Router({
  // mode: 'history', // 后端支持可开
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRouterMap
})
