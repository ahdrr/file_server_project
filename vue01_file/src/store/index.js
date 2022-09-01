/* eslint-disable no-tabs */
/* eslint-disable indent */
import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

// 初始化时用sessionStore.getItem('token'),这样子刷新页面就无需重新登录
const state = {
	token: window.localStorage.getItem('token'),
	username: '',
	list: []
}

const getters = {
    count: state => isComplete => {
		return state.list.filter(item => item.isComplete === isComplete).length
	},
	todoList: state => {
		return state.list
	}
}

const mutations = {
	LOGIN: (state, data) => {
		// 更改token的值
		state.token = data
		window.localStorage.setItem('token', data)
	},
	LOGOUT: (state) => {
		// 登出的时候要清除token
		state.token = null
		window.localStorage.removeItem('token')
	},
	USERNAME: (state, data) => {
		// 把用户名存起来
		state.username = data
		window.localStorage.setItem('username', data)
	}

}

const actions = {
	UserLogin ({ commit }, data) {
		commit('LOGIN', data)
	},
	UserLogout ({ commit }) {
		commit('LOGOUT')
	},
	UserName ({ commit }, data) {
		commit('USERNAME', data)
	}
}

export default new Vuex.Store({
	state,
	mutations,
	actions,
	getters
})
