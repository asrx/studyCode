import Vue from 'vue'
import Router from 'vue-router'
import HelloWorld from '@/components/HelloWorld'
import Todos from '@/components/Todos'
import Todo from '@/components/Todo'
import TodoForm from '@/components/TodoForm'

Vue.use(Router)

const routes = [
  { path: '/', component: Todos},
  { path: '/todo/:id', component: Todo, name: 'todo' },
]

export default new Router({
  routes
})
