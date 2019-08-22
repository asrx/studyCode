import Vue from 'vue'
import Vuex from 'vuex'
import GLOBAL from '@/common/global_variable.js'

Vue.use(Vuex)

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
      Vue.axios.get(GLOBAL.baseURL + '/api/todos').then(response => {
        store.commit('get_todos',response.data)
      })
    },
    completeTodo(store, todo){
      Vue.axios.patch(GLOBAL.baseURL + '/api/todo/'+ todo.id +'/completed').then(response=>{
        store.commit('complete_todo', todo)
      })
    },
    removeTodo(store, index, todo){
      Vue.axios.delete(GLOBAL.baseURL + '/api/todo/'+ todo.id +'/completed').then(response=>{
        store.commit('delete_todo', index)
      })
    },
    insertTodo(store, todo){
      Vue.axios.post(GLOBAL.baseURL + '/api/todo/create',{title:todo.title}).then(response=>{
        store.commit('add_todo', response.data)
      })

      store.state.newTodo = {id:null, title: '', completed: false}
    }
  }
})

export default store
