import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '../views/layout/Layout'

export const constantRoutes = [
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
      meta: { title: '首页', icon: 'dashboard', noCache: true, affix: true },
      component: () => import('@/views/dashboard/index')
    }]
  },
  // 分发上传
  {
    path: '/dispatch',
    component: Layout,
    name: 'Dispatch',
    meta: { title: '分发管理', icon: 'distribute', noCache: true, affix: true },
    children: [{
      path: 'upload/',
      meta: { title: '分发上传', icon: 'upload', noCache: true, affix: true },
      component: () => import('@/views/upload/index')
    }, {
      path: 'manager/',
      meta: { title: '管理链接', icon: 'list', noCache: true, affix: true },
      component: () => import('@/views/upload/show')
    }

    ]
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
      meta: { title: '用户管理', icon: 'users', noCache: true, affix: true }
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
    meta: { title: '图片管理', icon: 'manage', noCache: true, affix: true },
    children: [
      {
        path: 'management/',
        name: 'Images',
        component: () => import('@/views/card/index'),
        meta: { title: '所有图片', icon: 'images', noCache: true, affix: true }
      },
      {
        path: 'management/sync',
        name: 'SyncImages',
        component: () => import('@/views/card/sync/index'),
        meta: { title: '同步图片管理', icon: 'sync', noCache: true, affix: true }
      }
    ]
  },
  // 日志记录
  {
    path: '/logs',
    component: Layout,
    name: 'Logs',
    children: [{
      path: 'index',
      name: 'Logs',
      meta: { title: '日志记录', icon: 'logs', noCache: true, affix: true },
      component: () => import('@/views/logs/index')
    }]
  },
  // 各种设置
  {
    path: '/config',
    component: Layout,
    name: 'config',
    meta: { title: '设置', icon: 'config' },
    children: [{
      path: 'index',
      name: 'siteConfig',
      meta: { title: '站点设置', icon: 'config', noCache: true, affix: true },
      component: () => import('@/views/siteconfig/index')
    }, {
      path: 'menu',
      name: 'menuConfig',
      meta: { title: '菜单设置', icon: 'cloud', noCache: true, affix: true },
      component: () => import('@/views/siteconfig/menu')
    }, {
      path: 'dispatch',
      name: 'storeConfig',
      meta: { title: '图床设置', icon: 'img-manager', noCache: true, affix: true },
      component: () => import('@/views/siteconfig/dispatch')
    }]
  },

  // 重定向
  { path: '*', redirect: '/404', hidden: true }

  // 测试路由 是否会在 Menu 上面自动显示
]

export const asyncRoutes = [
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    name: 'Dashboard',
    children: [{
      path: 'dashboard',
      meta: { title: '首页',
        icon: 'dashboard',
        roles: ['admin']
      },
      component: () => import('@/views/dashboard/index')
    }]
  },
  // 分发上传
  {
    path: '/upload/',
    component: Layout,
    // redirect: '/d',
    name: 'Upload',
    children: [{
      path: '/v2',
      meta: { title: '分发上传', icon: 'distribute', noCache: true, affix: true },
      component: () => import('@/views/upload/index')
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
      meta: { title: '用户管理', icon: 'users', noCache: true, affix: true }
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
    meta: { title: '图片管理', icon: 'manage', noCache: true, affix: true },
    children: [
      {
        path: 'management/',
        name: 'Images',
        component: () => import('@/views/card/index'),
        meta: { title: '所有图片', icon: 'images', noCache: true, affix: true }
      },
      {
        path: 'management/sync',
        name: 'SyncImages',
        component: () => import('@/views/card/sync/index'),
        meta: { title: '同步图片管理', icon: 'sync', noCache: true, affix: true }
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
      meta: { title: '日志记录', icon: 'logs', noCache: true, affix: true },
      component: () => import('@/views/logs/index')
    }]
  },
  {
    path: '/config',
    component: Layout,
    name: 'config',
    meta: { title: '设置', icon: 'config' },
    children: [{
      path: 'index',
      name: 'Config1',
      meta: { title: '站点设置', icon: 'config', noCache: true, affix: true },
      component: () => import('@/views/siteconfig/index')
    }, {
      path: 'menu',
      name: 'Config2',
      meta: { title: '菜单设置', icon: 'cloud', noCache: true, affix: true },
      component: () => import('@/views/siteconfig/menu')
    }]
  }
]

// export default new Router({
//   // mode: 'history', // 后端支持可开
//   scrollBehavior: () => ({ y: 0 }),
//   routes: constantRouterMap
// })

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
