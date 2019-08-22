// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import Vuex from 'vuex'

import App from './App'
import router from './router'

import GLOBAL from './common/global_variable.js'

Vue.use(VueAxios,axios)
Vue.use(Vuex)

Vue.config.productionTip = false

Vue.prototype.baseUrl = 'http://longyuan-query.test'


const store = new Vuex.Store({
  state: {
    todos: [],
    newTodo: {id:null, title: '', completed: false}
  },
  mutations: {
    get_todos (state, todos) {
      state.todos = todos
    },
    complete_todo(state, todo) {
      todo.completed = ! todo.completed
    },
    delete_todo(state, index){
      state.todos.splice(index,1)
    },
    add_todo(state, todo){
      state.todos.push(todo)
    }
  },
  actions: {
    getTodos(store){
      Vue.axios.get('/api/todos').then(response => {
        store.commit('get_todos',response.data)
      })
    },
    completeTodo(store, todo){
      Vue.axios.patch('/api/todo/'+ todo.id +'/completed').then(response=>{
        store.commit('complete_todo', todo)
      })
    },
    removeTodo(store, index, todo){
      Vue.axios.delete('/api/todo/'+ todo.id +'/completed').then(response=>{
        store.commit('delete_todo', index)
      })
    },
    insertTodo(store, todo){
      Vue.axios.post('/api/todo/create',{title:todo.title}).then(response=>{
        store.commit('add_todo', response.data)
      })

      store.state.newTodo = {id:null, title: '', completed: false}
    }
  }
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  store,
  router,
  components: { App },
  template: '<App/>'
})
