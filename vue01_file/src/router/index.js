import Vue from 'vue'
import Router from 'vue-router'
import store from '../store/index'
import frontpage from '@/components/frontpage'
import login from '@/components/login'
import register from '@/components/register'
Vue.use(Router)

const originalPush = Router.prototype.push
Router.prototype.push = function push (location) {
  return originalPush.call(this, location).catch(err => err)
}

const router = new Router({
  mode: 'history',
  routes: [
    {
      path: '/login',
      name: 'login',
      component: login
    },
    {
      path: '/*',
      name: 'home',
      component: frontpage,
      meta: {
        requiresAuth: true
      }
    },
    {
      path: '/frontpage*',
      name: 'frontpage',
      component: frontpage,
      meta: {
        requiresAuth: true
      }
    },

    {
      path: '/register',
      name: 'register',
      component: register
    }
  ]
})

// 注册全局钩子用来拦截导航
router.beforeEach((to, from, next) => {
  // 获取store里面的token
  let token = store.state.token

  // 判断要去的路由有没有requiresAuth
  if (to.meta.requiresAuth) {
    if (token) {
      next()
    } else {
      next({
        path: '/login',
        query: { redirect: to.fullPath } //  将刚刚要去的路由path（却无权限）作为参数，方便登录成功后直接跳转到该路由
      })
    }
  } else {
    next() // 如果无需token,那么随它去吧
  }
})
export default router
